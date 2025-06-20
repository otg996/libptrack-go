
# Contributing to the Programming Project Tracker Library

Thank you for your interest in contributing! This document provides the specific technical setup for this repository. For the overall project philosophy, commit standards, and process, please refer to the [**main contributing guide**](https://github.com/otg996/programming-project-tracker-spec/blob/main/CONTRIBUTING.md) in the specification repository.

## Development Environment Setup

Our development environment relies on Go for the core code and Node.js for our development workflow tooling.

### Prerequisites

1. **Go** (version 1.23 or newer)
2. **Node.js** (LTS version)
3. **pnpm** (our preferred package manager)

### Installation

1. **Fork** the repository and **clone** your fork locally.
2. Navigate into the project directory.
3. **Install all dependencies and Git hooks with two commands:**

    ```bash
    # Install Node.js dev tools (commitizen, commitlint, etc.)
    # This also runs 'lefthook install' automatically.
    pnpm install

    # Download Go dependencies (testify, golangci-lint)
    go mod tidy
    ```

Your environment is now fully configured to match our CI system.

## Running Tests and Linters

Before committing, you can manually run the same checks our pre-commit hooks do:

* **Run Linter:** `go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...`
* **Run Tests:** `go test -v ./...`
