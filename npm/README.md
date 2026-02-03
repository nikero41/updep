# updep

[![CI](https://github.com/nikero41/updep/actions/workflows/ci.yml/badge.svg)](https://github.com/nikero41/updep/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/snikoletopoulos/updep)](https://github.com/snikoletopoulos/updep)
[![License](https://img.shields.io/github/license/snikoletopoulos/updep)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/snikoletopoulos/updep)](https://goreportcard.com/report/github.com/snikoletopoulos/updep)

> **Note**: This project is currently under active development. Features and APIs may change.

Interactive TUI for updating JavaScript dependencies. Works with npm, yarn, pnpm, and bun.

## Demo

> TODO: Add demo video

## Features

- Interactive terminal UI for dependency updates
- Review and select which packages to update
- Compare current, wanted, and latest versions
- Support for multiple package managers (npm support available, yarn/pnpm/bun coming soon)
- Built with Go for fast performance

## Installation

### From Source

```bash
git clone https://github.com/nikero41/updep.git
cd updep
make install
```

## Usage

Navigate to your JavaScript project directory and run:

```bash
updep
```

### Keybindings

- `↑/↓` or `j/k` - Navigate between packages
- `Space` - Toggle package selection
- `w` - Select wanted version
- `l` - Select latest version
- `q` or `Ctrl+C` - Quit

## Requirements

- Go 1.25+ (for building from source)
- npm, yarn, pnpm, or bun installed
- A JavaScript project with a `package.json` file

## Roadmap

- [ ] npm support
- [ ] yarn support
- [ ] pnpm support
- [ ] bun support
- [ ] Configuration file support
- [ ] Custom color themes

## Contributing

Contributions are welcome! Please check out our [Contributing Guide](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE) for details

## Acknowledgments

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss)
