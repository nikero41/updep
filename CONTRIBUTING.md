# Contributing to updep

Thank you for your interest in contributing to updep! We welcome contributions from the community.

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Git
- A JavaScript package manager (npm, yarn, pnpm, or bun) for testing

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/updep.git
   cd updep
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Development Workflow

### Making Changes

1. Create a new branch for your feature or bug fix:

   ```bash
   git switch -c feature/your-feature-name
   ```

2. Make your changes and test them:

   ```bash
   make run
   # or
   make dev-install
   ```

3. Format your code:

   ```bash
   make format
   ```

4. Build the project:

   ```bash
   make build
   ```

### Code Quality

Before submitting your changes, ensure:

- **Code is formatted**: Run `go fmt ./...`
- **No linting errors**: Our CI runs `golangci-lint` with 19+ linters
- **Code compiles**: Run `make build` successfully
- **Tests pass**: Run `go test ./...` (when tests are added)

### Commit Messages

Write clear, concise commit messages:

- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Move cursor to..." not "Moves cursor to...")
- Reference issues and pull requests when relevant

Good examples:

```
Add yarn support for package manager detection
Fix cursor navigation bug in package list
Update documentation for keybindings
```

### Submitting a Pull Request

1. Push your changes to your fork:

   ```bash
   git push origin feature/your-feature-name
   ```

2. Open a pull request on GitHub

3. Fill out the pull request template with:
   - Description of changes
   - Related issues
   - Testing performed
   - Screenshots (if applicable)

4. Wait for CI checks to pass

5. Address any review feedback

## What to Contribute

### Good First Issues

Look for issues labeled `good first issue` - these are great starting points for new contributors.

### Bug Reports

Found a bug? Please open an issue with:

- Steps to reproduce
- Expected behavior
- Actual behavior
- Your environment (OS, Go version, package manager)

## Code Style

- Follow standard Go conventions
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and small
- Use the Bubble Tea framework patterns for UI components

## Project Structure

```
updep/
├── cmd/updep/          # Main application entry point
├── pkg/
│   ├── components/   # Reusable UI components
│   ├── config/       # Configuration and themes
├── .github/          # GitHub Actions workflows
└── Makefile          # Build commands
```

## Testing

Currently, the project doesn't have extensive tests. Adding tests is a great way to contribute! We welcome:

- Unit tests for models and utilities
- Integration tests for package manager interactions
- UI component tests

## Getting Help

- Open a [Discussion](https://github.com/snikoletopoulos/updep/discussions) for questions
- Check existing issues and pull requests
- Reach out to maintainers in your pull request

## Code of Conduct

Be respectful and inclusive. We're all here to build something great together.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
