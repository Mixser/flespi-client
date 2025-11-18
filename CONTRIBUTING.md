# Contributing to Flespi Go Client

Thank you for considering contributing to the Flespi Go Client! We welcome contributions from the community.

## Development Setup

1. **Fork and clone the repository**

```bash
git clone https://github.com/mixser/flespi-client.git
cd flespi-client
```

2. **Install dependencies**

```bash
go mod download
```

3. **Install development tools**

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Development Workflow

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Linting

```bash
# Run golangci-lint
golangci-lint run

# Auto-fix issues where possible
golangci-lint run --fix
```

### Building

```bash
# Build all packages
go build ./...
```

## Code Style

- Follow standard Go conventions and idioms
- Run `gofmt` on all code (this is enforced by CI)
- Write clear, descriptive commit messages
- Add tests for new functionality
- Maintain or improve code coverage (target: >60%)
- Document all exported functions, types, and packages

## Pull Request Process

1. **Create a feature branch**

```bash
git checkout -b feature/your-feature-name
```

2. **Make your changes**
   - Write code following the style guidelines
   - Add or update tests
   - Update documentation if needed
   - Ensure all tests pass
   - Run linting

3. **Commit your changes**

```bash
git add .
git commit -m "Clear description of your changes"
```

4. **Push to your fork**

```bash
git push origin feature/your-feature-name
```

5. **Open a Pull Request**
   - Provide a clear description of the changes
   - Reference any related issues
   - Ensure CI checks pass

## Testing Guidelines

- Write unit tests for all new functionality
- Use table-driven tests where appropriate
- Mock HTTP requests using `httptest`
- Test error conditions and edge cases
- Aim for high test coverage

Example test structure:

```go
func TestNewResource(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    *Resource
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   "test",
            want:    &Resource{Name: "test"},
            wantErr: false,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewResource(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewResource() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // Add assertions
        })
    }
}
```

## Documentation

- Add GoDoc comments for all exported symbols
- Update README.md for user-facing changes
- Add examples in the `examples/` directory for new features
- Keep CHANGELOG.md updated

## Code Review

All submissions require review. We use GitHub pull requests for this purpose. Reviewers will check for:

- Code quality and style
- Test coverage
- Documentation
- Backward compatibility
- Performance implications

## Questions?

Feel free to open an issue for:
- Bug reports
- Feature requests
- Questions about contributing
- General discussions

## License

By contributing to this project, you agree that your contributions will be licensed under the same license as the project.
