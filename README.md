# Coding Playground

[![Tests](https://github.com/iamhectorsosa/interview-prep/workflows/test.yml/badge.svg)](https://github.com/iamhectorsosa/interview-prep/actions)

A collection of problems and solutions implemented in Go and TypeScript with unit tests.

## Structure

```txt
├── go/problems/     # Go solutions with tests
├── ts/problems/     # TypeScript solutions with tests
└── Makefile         # Execute commands at root
```

## Usage

```bash
# Run tests
make test-go        # Watch Go tests
make test-ts        # Watch TypeScript tests
make test-go-run    # Run Go tests once
make test-ts-run    # Run TypeScript tests once
```

## Requirements

- Go 1.25+
- Node.js with pnpm
- [rg](https://github.com/BurntSushi/ripgrep) and [entr](https://eradman.com/entrproject/) for Go file watching
