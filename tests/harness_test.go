package e2e

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui"
	mathcontroller "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	mathmemory "github.com/1kovalevskiy/math-trainer/internal/storages/memory/math"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	defaultExamplesCount = mathmodels.DefaultExamplesCount
	defaultTimeout       = 2 * time.Second
	sessionTimeout       = 10 * time.Second
	pollInterval         = 10 * time.Millisecond
)

var ansiPattern = regexp.MustCompile(`\x1b\[[0-?]*[ -/]*[@-~]`)

type session struct {
	ctx     context.Context
	cancel  context.CancelFunc
	store   *mathmemory.Storage
	gen     *exerciseGenerator
	program *tea.Program
	probe   *probeModel
	runDone chan runResult
}

type runResult struct {
	model tea.Model
	err   error
}

func newSession(t *testing.T, exercises ...mathmodels.Exercise) *session {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), timeoutFor(t, sessionTimeout))
	store := mathmemory.New()
	gen := newExerciseGenerator(exercises...)
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(gen.Generate))
	probe := newProbeModel(tui.NewModel(ctx, controller, nil))
	program := tea.NewProgram(
		probe,
		tea.WithInput(nil),
		tea.WithOutput(io.Discard),
		tea.WithContext(ctx),
		tea.WithoutSignals(),
	)

	s := &session{ctx: ctx, cancel: cancel, store: store, gen: gen, program: program, probe: probe, runDone: make(chan runResult, 1)}

	go func() {
		model, err := program.Run()
		s.runDone <- runResult{model: model, err: err}
	}()

	t.Cleanup(func() { s.stop(t) })

	select {
	case <-probe.ready:
	case <-time.After(timeoutFor(t, defaultTimeout)):
		cancel()
		t.Fatal("program did not render initial view")
	}

	s.send(t, tea.WindowSizeMsg{Width: 80, Height: 24})
	s.eventuallyViewContains(t, "Математический тренажер")

	return s
}

func (s *session) stop(t *testing.T) {
	t.Helper()

	s.program.Quit()
	select {
	case result := <-s.runDone:
		if result.err != nil && !errors.Is(result.err, tea.ErrProgramKilled) {
			t.Fatalf("program run error: %v", result.err)
		}
	case <-time.After(timeoutFor(t, defaultTimeout)):
		s.cancel()
		t.Fatal("program did not stop")
	}

	s.cancel()
}

func (s *session) send(t *testing.T, msg tea.Msg) {
	t.Helper()

	done := make(chan struct{})
	go func() {
		s.program.Send(msg)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(timeoutFor(t, defaultTimeout)):
		t.Fatalf("send timed out for %T", msg)
	}

	s.sync(t)
}

func (s *session) sync(t *testing.T) {
	t.Helper()

	ack := make(syncMsg)
	s.program.Send(ack)

	select {
	case <-ack:
	case <-time.After(timeoutFor(t, defaultTimeout)):
		t.Fatal("program sync timed out")
	}
}

func (s *session) key(t *testing.T, value string) {
	t.Helper()

	msg, err := keyMsg(value)
	if err != nil {
		t.Fatal(err)
	}
	s.send(t, msg)
}

func (s *session) typeText(t *testing.T, value string) {
	t.Helper()

	for _, r := range value {
		s.send(t, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
	}
}

func (s *session) applyExamplesCount(t *testing.T, target int) {
	t.Helper()

	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Настройки тренировки")

	for i := 0; i < 4; i++ {
		s.key(t, "down")
	}
	for current := defaultExamplesCount; current < target; current++ {
		s.key(t, "right")
	}
	for current := defaultExamplesCount; current > target; current-- {
		s.key(t, "left")
	}

	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математический тренажер")
	s.key(t, "up")
}

func (s *session) eventuallyViewContains(t *testing.T, want string) {
	t.Helper()

	deadline := time.Now().Add(timeoutFor(t, defaultTimeout))
	for time.Now().Before(deadline) {
		if strings.Contains(s.probe.lastView(), want) {
			return
		}
		time.Sleep(pollInterval)
	}

	t.Fatalf("view does not contain %q\nlatest view:\n%s", want, s.probe.lastView())
}

func (s *session) requireState(t *testing.T) mathmodels.TrainingState {
	t.Helper()

	ctx, cancel := assertionContext(t)
	defer cancel()

	state, err := s.store.GetState(ctx)
	if err != nil {
		t.Fatalf("GetState() unexpected error: %v", err)
	}
	return state
}

func (s *session) requireNoActiveTraining(t *testing.T) {
	t.Helper()

	deadline := time.Now().Add(timeoutFor(t, defaultTimeout))
	for time.Now().Before(deadline) {
		ctx, cancel := assertionContext(t)
		_, err := s.store.GetState(ctx)
		cancel()
		if errors.Is(err, mathmodels.ErrNoActiveTraining) {
			return
		}
		time.Sleep(pollInterval)
	}

	ctx, cancel := assertionContext(t)
	defer cancel()

	_, err := s.store.GetState(ctx)
	t.Fatalf("GetState() error mismatch: got %v, want %v", err, mathmodels.ErrNoActiveTraining)
}

func (s *session) requireGeneratedSettings(t *testing.T, want ...mathmodels.TrainingSettings) {
	t.Helper()

	got := s.gen.Settings()
	if len(got) != len(want) {
		t.Fatalf("generated settings count mismatch: got %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("generated settings[%d] mismatch: got %+v, want %+v", i, got[i], want[i])
		}
	}
}

type probeModel struct {
	mu      sync.RWMutex
	inner   tea.Model
	view    string
	updates int
	ready   chan struct{}
	readyMu sync.Once
}

type syncMsg chan struct{}

func newProbeModel(inner tea.Model) *probeModel {
	return &probeModel{inner: inner, ready: make(chan struct{})}
}

func (p *probeModel) Init() tea.Cmd {
	p.recordView(p.inner.View())
	p.readyMu.Do(func() { close(p.ready) })
	return p.inner.Init()
}

func (p *probeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if ack, ok := msg.(syncMsg); ok {
		p.recordView(p.inner.View())
		close(ack)
		return p, nil
	}

	next, cmd := p.inner.Update(msg)
	p.mu.Lock()
	p.inner = next
	p.view = stripANSI(next.View())
	p.updates++
	p.mu.Unlock()
	return p, cmd
}

func (p *probeModel) View() string {
	view := p.inner.View()
	p.recordView(view)
	return view
}

func (p *probeModel) lastView() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.view
}

func (p *probeModel) updateCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.updates
}

func (p *probeModel) recordView(view string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.view = stripANSI(view)
}

func exercise(left int, right int, operator mathmodels.Operator) mathmodels.Exercise {
	return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
}

type exerciseGenerator struct {
	mu        sync.Mutex
	exercises []mathmodels.Exercise
	next      int
	settings  []mathmodels.TrainingSettings
}

func newExerciseGenerator(exercises ...mathmodels.Exercise) *exerciseGenerator {
	return &exerciseGenerator{exercises: exercises}
}

func (g *exerciseGenerator) Generate(settings mathmodels.TrainingSettings) (mathmodels.Exercise, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.settings = append(g.settings, settings)
	if g.next >= len(g.exercises) {
		return mathmodels.Exercise{}, errors.New("exercise generator exhausted")
	}

	exercise := g.exercises[g.next]
	g.next++
	return exercise, nil
}

func (g *exerciseGenerator) Settings() []mathmodels.TrainingSettings {
	g.mu.Lock()
	defer g.mu.Unlock()

	return append([]mathmodels.TrainingSettings(nil), g.settings...)
}

func keyMsg(value string) (tea.KeyMsg, error) {
	switch value {
	case "enter":
		return tea.KeyMsg{Type: tea.KeyEnter}, nil
	case "esc":
		return tea.KeyMsg{Type: tea.KeyEsc}, nil
	case "up":
		return tea.KeyMsg{Type: tea.KeyUp}, nil
	case "down":
		return tea.KeyMsg{Type: tea.KeyDown}, nil
	case "left":
		return tea.KeyMsg{Type: tea.KeyLeft}, nil
	case "right":
		return tea.KeyMsg{Type: tea.KeyRight}, nil
	case "backspace":
		return tea.KeyMsg{Type: tea.KeyBackspace}, nil
	case "ctrl+c":
		return tea.KeyMsg{Type: tea.KeyCtrlC}, nil
	default:
		if len([]rune(value)) == 1 {
			return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(value)}, nil
		}
		return tea.KeyMsg{}, fmt.Errorf("unsupported key %q", value)
	}
}

func stripANSI(value string) string {
	return ansiPattern.ReplaceAllString(value, "")
}

func timeoutFor(t *testing.T, preferred time.Duration) time.Duration {
	t.Helper()

	deadline, ok := t.Deadline()
	if !ok {
		return preferred
	}

	remaining := time.Until(deadline) - 100*time.Millisecond
	if remaining <= 0 {
		return preferred
	}
	if remaining < preferred {
		return remaining
	}

	return preferred
}

func assertionContext(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()

	return context.WithTimeout(context.Background(), timeoutFor(t, defaultTimeout))
}
