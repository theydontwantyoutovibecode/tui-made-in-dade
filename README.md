# tui

A Go terminal UI application template using Charm Bubbletea v2. Created with [dade](https://github.com/theydontwantyoutovibecode/dade).

## Prerequisites

- Go 1.22+

Installed automatically by `setup.sh` if missing.

## Project Structure

```
.
├── main.go                 # Entry point, tea.NewProgram
├── model.go                # Model struct and Init()
├── update.go               # Update() message handler
├── view.go                 # View() renderer
├── keys.go                 # Key bindings
├── styles.go               # Lipgloss styles
├── internal/
│   └── config/
│       └── config.go       # Configuration
├── go.mod
├── go.sum
├── setup.sh                # First-run setup
├── dade.toml          # Template manifest
├── AGENTS.md               # AI agent instructions
└── .read-only/
    └── manifest.txt        # Reference repos for AI context
```

## Usage

### Development

```bash
dade dev
```

This runs `go mod download` to fetch dependencies, then `go run .` with `DEBUG=1` set in the environment.

When `DEBUG=1` is set, the program logs to `debug.log` instead of stdout (since stdout is used for the TUI).

### Build

```bash
dade build
```

Runs `go build` and outputs the binary to `./bin/`. The build command runs `go mod tidy` before building.

```bash
dade build --release
```

Builds with `-ldflags '-s -w'` to strip debug symbols and reduce binary size.

```bash
dade build --os linux --arch amd64
```

Cross-compiles for a different platform.

```bash
dade build --all
```

Builds for all supported OS/architecture combinations.

### Not Applicable

- `dade project start` — No server. TUI programs run interactively in the terminal.
- `dade share` — No server to tunnel.
- `dade share --attach` — No server to tunnel.
- `dade dev --open` — No web URL.

## Elm Architecture

Bubbletea uses the Elm Architecture. The program loop works as follows:

1. **Init** — Returns the initial model and any startup commands
2. **Update** — Receives a message, returns an updated model and commands
3. **View** — Renders the model as a string for the terminal

All state lives in the `Model` struct in `model.go`. The `Update` function in `update.go` handles keyboard input and other messages. The `View` function in `view.go` renders the current state using Lipgloss styles from `styles.go`.

### Key Bindings

Key bindings are defined in `keys.go` using Bubbletea's `key.NewBinding`. The `Update` function matches incoming `tea.KeyMsg` values against these bindings.

### Styles

Terminal styles (colors, borders, padding, alignment) are defined in `styles.go` using Lipgloss. These are applied in the `View` function to format output.

## Bubbletea v2

This template uses Bubbletea v2, which has breaking changes from v1:

- `tea.Model` is now `tea.Model` with generics
- `Init()` returns `(Model, tea.Cmd)` instead of `tea.Cmd`
- `Update()` and `View()` signatures changed
- Bubbles components updated to v2 with matching API changes

The `.read-only/` reference libraries contain the v2 source. Refer to them for the current API, not v1 documentation found online.

## Reference Libraries

The `.read-only/manifest.txt` file lists repositories that are shallow-cloned during `dade dev`:
- bubbletea — TUI framework
- bubbles — Common TUI components (text input, list, table, etc.)
- lipgloss — Terminal styling

These provide AI coding agents with reference source code. This is important because Bubbletea v2 documentation is limited and the v1 API shown in most online resources is outdated.

## Ticket Tracking

The `.tickets/` directory is used with the `tk` CLI for ticket-driven development. See `AGENTS.md` for the workflow.
