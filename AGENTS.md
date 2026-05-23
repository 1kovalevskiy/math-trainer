# Руководство по репозиторию math-trainer

## 1. Назначение проекта
`math-trainer` — одноразовое TUI-приложение для тренировки устного счета у детей.

Технологический стек:
- Go 1.25;
- Bubble Tea;
- Lip Gloss;
- bubblezone для mouse hit-zones;
- cleanenv + godotenv для конфигурации;
- slog для структурированного логирования.

Ключевое продуктовое решение: приложение рассчитано на одну TUI-сессию. После закрытия приложения состояние тренировки очищается, потому что storage in-memory.

## 2. Текущая структура
- `cmd/main.go` — минимальная точка входа: создает `context.Background()`, вызывает `app.InitApp`, затем `RunApp`; `panic` допустим только здесь.
- `cmd/app` — composition root и жизненный цикл:
  - `app.go` — `App`, `InitApp`, `RunApp`, centralized closers;
  - `config.go` — загрузка конфига;
  - `logger.go` — настройка `slog`;
  - `providers.go` — явный, но пока пустой шаг providers;
  - `storages.go` — создание `mathMemory.Storage`;
  - `controllers.go` — создание `mathController.Controller`;
  - `programs.go` — создание Bubble Tea program и bubblezone.
- `internal/configs` — пакет `config`; сейчас есть только `App.LogLevel`, `APP_CONFIG_PATH`, `APP_LOG_LEVEL`, автозагрузка `.env`.
- `internal/models/math` — доменные модели, enum/value-типы и доменные ошибки. Файлы разделены по сущностям: `exercise.go`, `training_state.go`, `training_snapshot.go`, `training_settings.go`, `example_result.go`, enum-файлы, `errors.go`.
- `internal/storages/memory/math` — in-memory storage состояния одной тренировки:
  - публичные операции: `GetState`, `SaveState`, `ClearState`, `CloseStorage`;
  - storage потокобезопасен через `sync.RWMutex`;
  - при чтении/записи клонирует state/results, чтобы не отдавать наружу mutable aliases;
  - `CloseStorage` очищает state.
- `internal/controllers/math` — orchestration и вся математика:
  - владеет контрактом `trainingStorage`;
  - синхронный API для TUI: `GetDefaultSettings`, `NormalizeSettings`, `StartTraining`, `SubmitAnswer`, `SkipCurrent`, `CancelTraining`, difficulty navigation;
  - генерация примеров, проверка ответа, сбор результатов и snapshot/summary находятся здесь;
  - тестовая инъекция генератора идет через `WithExerciseGenerator`.
- `internal/app/tui` — Bubble Tea UI:
  - root `Model` хранит текущий экран, размеры окна, настройки и submodels;
  - subpackages `start`, `settings`, `task`, `result` отвечают за локальное состояние экрана, `Update`, `View`, typed messages;
  - `commands.go` вызывает контроллер из `tea.Cmd` и возвращает typed messages;
  - `layout.go` центрирует контент и держит подсказки на фиксированной высоте от низа;
  - `ui/theme.go` содержит общие стили;
  - `shared` содержит только UI-format helpers и mouse helpers.

## 3. Направление зависимостей
Разрешенное направление зависимостей:

`cmd/app -> internal/app/tui -> internal/controllers/math -> (internal/models/math, internal/storages contracts)`

Фактически:
- `cmd/app` знает конкретные implementations (`mathmemory.Storage`, `mathcontroller.Controller`) и собирает приложение.
- `internal/app/tui` знает только интерфейс `mathController`, объявленный в `internal/app/tui/controller.go`.
- `internal/controllers/math` зависит от доменных моделей и от локального интерфейса `trainingStorage`, но не от конкретного memory storage.
- `internal/storages/memory/math` зависит только от `internal/models/math`.
- `internal/models/math` не зависит от UI, storage, controller, Bubble Tea или SDK.

Запрещено:
- ходить из TUI напрямую в storage/provider;
- импортировать Bubble Tea в controller/storage/models;
- хранить бизнес-правила в `View` или screen `Update`;
- делать wiring вне `cmd/app`;
- создавать циклические импорты между слоями.

## 4. Принятые архитектурные решения
- Состояние тренировки хранится не в TUI model, а в in-memory storage.
- Доступ к storage есть только у controller через интерфейс.
- TUI держит только UI-состояние: текущий экран, ввод, курсоры кнопок, размеры окна, последнее отображаемое состояние.
- Controller остается синхронным. Асинхронность Bubble Tea организуется только в `internal/app/tui/commands.go`.
- Все решенные/пропущенные примеры текущей сессии лежат в `TrainingState.Results`.
- Следующий пример генерируется через `generateUnused`: controller сравнивает новый `Exercise` с уже использованными `Exercise` из `Results`.
- Повтор определяется полной структурой `Exercise`: `Left`, `Right`, `Operator`.
- Если уникальный пример не найден за `maxUniqueExerciseAttempts`, возвращается `ErrUniqueExerciseExhausted`.
- При закрытии приложения storage очищается; персистентность специально не нужна.

## 5. Инициализация и shutdown
`InitApp` идет строго по шагам:
1. config;
2. logger;
3. providers;
4. storages;
5. controllers;
6. Bubble Tea program.

Правила:
- новый инфраструктурный шаг добавлять явно в цепочку `InitApp`;
- каждый `init*` возвращает `error`;
- closers регистрировать через `addCloser(name, fn)`;
- закрытие идет LIFO через `closeAll`;
- ошибки закрытия агрегируются через `errors.Join`;
- `RunApp` использует `signal.NotifyContext`;
- shutdown context создается непосредственно перед `closeAllWithTimeout`, а не переиспользует уже отмененный run context.

Сейчас closers регистрируются для:
- `math_memory_storage`;
- `bubblezone`;
- `bubbletea_program`.

## 6. Правила для моделей
- Пакет `internal/models/math` содержит только доменные типы и ошибки.
- Новую доменную сущность добавлять отдельным файлом по имени сущности.
- Enum/value-типы держать в отдельных файлах с безопасным `String()` fallback на `"unknown"`.
- Ошибки держать в `errors.go`.
- Не добавлять методы, завязанные на отображение, текст UI или Bubble Tea.

## 7. Правила для controller
- Controller — место для всей математики и orchestration.
- Публичные операции с логикой держать в отдельных файлах в стиле `entity_action.go`, метод называть `ActionEntity`:
  - `training_start.go` -> `StartTraining`;
  - `answer_submit.go` -> `SubmitAnswer`;
  - `current_skip.go` -> `SkipCurrent`;
  - `settings_normalize.go` -> `NormalizeSettings`;
  - `difficulty_get_next.go` -> `GetNextDifficulty`.
- Интерфейсы зависимостей объявлять рядом с использованием, сейчас это `trainingStorage` в `controller.go`.
- `context.Context` передавать сверху вниз, не заменять на `context.Background()` внутри бизнес-операций.
- Ошибки оборачивать через `%w` с контекстом операции.
- Для тестов использовать stub генератор через `WithExerciseGenerator`, а не управлять random напрямую.
- При изменении генерации примеров обязательно проверять сценарии повторов и исчерпания уникальных вариантов.

## 8. Правила для storage
- Memory storage является адаптером инфраструктуры, а не местом бизнес-логики.
- Публичные методы держать по отдельным файлам в стиле `state_get.go` -> `GetState`.
- Методы принимают `context.Context` и сначала проверяют `ctx.Err()`.
- `nil` receiver должен обрабатываться безопасно там, где это уже принято.
- `GetState` возвращает `ErrNoActiveTraining`, если state еще не создан или очищен.
- `SaveState` и `GetState` должны клонировать mutable данные.
- `CloseStorage` должен быть idempotent по смыслу и очищать in-memory state.

## 9. Правила Bubble Tea / TUI
- `Update` должен быть быстрым и неблокирующим.
- Вызовы controller, I/O и долгие операции выполнять только через `tea.Cmd`.
- `tea.Cmd` возвращает только typed `tea.Msg`.
- `View` только рендерит состояние; сайд-эффекты запрещены.
- Root `internal/app/tui/Model` маршрутизирует сообщения между экранами и запускает commands.
- Screen submodels (`start`, `settings`, `task`, `result`) не должны знать storage.
- Screen submodels могут эмитить только свои typed messages, а root решает, что делать дальше.
- Форматирование доменных значений для UI держать в `internal/app/tui/shared/format.go`, не в models/controller.
- Для mouse support использовать `zone.Mark` во `View` и `shared.InZone` в `Update`.
- Любой новый экран должен иметь:
  - `model.go`;
  - `update.go`;
  - `view.go`;
  - `messages.go`;
  - `commands.go`, если экран эмитит команды.

Текущие UX-решения:
- весь экранный контент центрируется через `renderScreenContent`;
- подсказки находятся на стабильной высоте от низа;
- слишком высокий контент сжимается через `fitBlock` с маркером пропущенных строк;
- стартовый экран: кнопки вертикально, с разрывом между ними;
- settings/task/result: action buttons расположены горизонтально;
- горизонтальные кнопки выбираются `←/→`; для совместимости местами поддержаны `↑/↓`;
- `Ctrl+C` завершает приложение глобально.

## 10. Конфигурация и логирование
- Основной путь к YAML: `APP_CONFIG_PATH`.
- ENV переопределяет YAML через cleanenv.
- `.env` загружается автоматически в `internal/configs`.
- Сейчас поддержан только `APP_LOG_LEVEL`, значения: `DEBUG`, `INFO`, `WARN`/`WARNING`, `ERROR`.
- В коде вне `internal/configs` и `cmd/app` не читать ENV напрямую без причины.
- Логирование только структурированное через `slog`.

## 11. Тестирование
Базовая проверка:

```bash
go test ./...
```

Если sandbox блокирует стандартный Go cache, использовать:

```bash
GOCACHE=/private/tmp/math-trainer-go-cache go test ./...
```

Что тестировать при изменениях:
- controller: orchestration, ошибки, state transitions, генерация без повторов;
- storage: clone semantics, пустое состояние, clear/close, context cancellation;
- TUI submodels: переходы по `tea.Msg`, keyboard/mouse navigation, emitted commands;
- layout: высота, центрирование, стабильная позиция подсказок, обрезка tall content.

Unit-тесты не должны требовать реальных сетей, БД или внешних сервисов.

## 12. Стиль кода
- Всегда запускать `gofmt` на измененных Go-файлах.
- Имена пакетов короткие, строчные, без underscore.
- В этом репозитории есть важные package aliases:
  - `internal/configs` использует `package config`;
  - `internal/controllers/math` использует `package mathcontroller`;
  - `internal/storages/memory/math` использует `package mathmemory`;
  - `internal/models/math` использует `package math`.
- Не складывать несколько публичных методов с логикой в один файл.
- Комментарии добавлять только там, где они реально объясняют неочевидное решение.
- Не делать unrelated refactor вместе с фичей.
- Не откатывать чужие незакоммиченные изменения.

## 13. Workflow новой фичи
1. Понять, на каком слое должна жить логика.
2. Если меняется домен, добавить/расширить модель в `internal/models/math`.
3. Сначала добавить тест на нужное поведение, обычно в controller/storage/TUI subpackage.
4. Реализовать orchestration в `internal/controllers/math`.
5. Если нужно состояние, добавить контракт controller и реализацию adapter в storage.
6. Подключить TUI через `tea.Cmd` + typed `tea.Msg`.
7. Обновить layout/navigation tests, если изменился UI.
8. Запустить `gofmt` и `go test ./...`.

Главное правило: математика и бизнес-решения остаются в controller/model/storage слоях, а `internal/app/tui` отвечает только за отображение, ввод и навигацию.
