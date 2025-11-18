# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-11-18

### Added
- Initial release of Flespi Go Client
- Full CRUD operations for Flespi resources:
  - Platform: Webhooks, Subaccounts, Limits
  - Gateway: Channels, Streams, Devices, Calculators, Geofences, Tokens
  - Storage: Containers, CDN
- Context support for request cancellation and timeouts
- Structured error handling with `APIError` type
- Error helper functions: `IsNotFoundError`, `IsUnauthorizedError`, `IsRateLimitError`
- Configurable HTTP client with options:
  - `WithHTTPClient`: Use custom HTTP client
  - `WithTimeout`: Set custom timeout
  - `WithRetryConfig`: Configure retry behavior
  - `WithLogger`: Set custom logger
- Exponential backoff retry mechanism
  - Configurable max retries (default: 3)
  - Configurable backoff duration (default: 1-30s)
  - Automatic retry on transient errors (429, 500, 502, 503, 504)
- Comprehensive logging support
  - `Logger` interface for custom implementations
  - `StdLogger` for stdout logging with log levels
  - `NoOpLogger` for silent operation
  - Log levels: Debug, Info, Warn, Error
- Input validation helpers
  - `ValidateID`, `ValidateRequired`, `ValidateURL`
  - `ValidationError` type
- Comprehensive test suite with 60-75% coverage
- Complete documentation:
  - README with usage examples
  - GoDoc comments for all exported symbols
  - Example programs in `examples/` directory
  - Contributing guidelines
- CI/CD pipeline with GitHub Actions
  - Automated testing on Go 1.22 and 1.23
  - Linting with golangci-lint
  - Code coverage reporting

### Fixed
- Resource leak: HTTP response body is now properly closed
- Error handling: Fixed nil return instead of error in `RequestAPI`
- Typos in public API: `NewSignleWebhook` → `NewSingleWebhook`
- Typos in public API: `CreateChaniedWebhookOption` → `CreateChainedWebhookOption`

### Changed
- Improved HTTP status code handling to accept all 2xx codes (not just 200)
- Enhanced error messages with structured error types
- Better request/response logging throughout the client

## [0.1.0] - YYYY-MM-DD

Initial development release.

[Unreleased]: https://github.com/mixser/flespi-client/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/mixser/flespi-client/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/mixser/flespi-client/releases/tag/v0.1.0
