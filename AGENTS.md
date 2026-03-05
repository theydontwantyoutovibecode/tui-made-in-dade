# AGENTS.md — TUI Template

## Stack

- **Go 1.24+** with modules
- **Bubbletea v2** — Elm-architecture TUI framework (`charm.land/x/bubbletea/v2`)
- **Bubbles v2** — pre-built TUI components (`charm.land/x/bubbles/v2`)
- **Lipgloss v2** — terminal styling (`charm.land/x/lipgloss/v2`)

All libraries use v2 imports from `charm.land/` (not `github.com/charmbracelet/`).

## Project Structure

```
.
├── main.go             # Entry point (tea.NewProgram)
├── model.go            # Model struct and Init()
├── update.go           # Update() message handler
├── view.go             # View() renderer
├── keys.go             # Key bindings
├── styles.go           # Lipgloss styles
├── internal/
│   └── config/config.go
├── go.mod / go.sum
├── .read-only/         # Reference libraries (auto-synced, do not edit)
│   └── manifest.txt
└── .tickets/           # Ticket tracking (tk CLI)
```

## Development Workflow

All work is **ticket-driven**. Never start implementation without a ticket.

### Planning Phase

1. Receive a request → analyze fully before acting
2. Create granular tickets with `tk`:
   ```
   tk create "Add file picker with fuzzy search" -t task -p 1 --tags "feature" -d "Definition of done: textinput for filter, list shows matching files, enter opens file, esc returns to main view, test for filter logic"
   ```
3. Each ticket must have: definition of done, caveats, test plan
4. Surface open questions before implementation

### Execution Phase

- **One ticket = one commit.** Never work on multiple tickets simultaneously.
- `tk start <id>` → implement → verify → `tk close <id>` → commit
- Run: `go run .`
- Test: `go test ./...`
- Build: `go build -o bin/myapp .`

### Reference Libraries

The `.read-only/` directory contains reference implementations. Use for context — never modify.

## Stack Notes — Bubbletea v2 Breaking Changes

These are critical differences from Bubbletea v1:

- `View()` returns `tea.View`, not `string` — use `v := tea.NewView(s)` and set `v.AltScreen = true`
- No `tea.WithAltScreen()` — set `v.AltScreen = true` in the `View` return value
- `KeyMsg` is an interface with `.Key()` method, not a struct with `.Type`
- Check modifiers: `msg.Key().Mod.Contains(tea.ModCtrl)` instead of `msg.Type == tea.KeyCtrlC`
- `textinput.Focus()` returns `tea.Cmd` — must be returned from `Update()`
- `viewport.New()` takes variadic options, not `(width, height)` args
- `list.New()` and `table.New()` also use variadic option patterns
- Imports: `charm.land/x/bubbletea/v2`, `charm.land/x/bubbles/v2`, `charm.land/x/lipgloss/v2`
- The Elm architecture: `Init() -> (Model, tea.Cmd)`, `Update(tea.Msg) -> (tea.Model, tea.Cmd)`, `View() -> tea.View`
