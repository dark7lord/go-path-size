# go-path-size

[![Hexlet tests and linter status](https://github.com/dark7lord/go-project-242/actions/workflows/test-ci.yml/badge.svg)](https://github.com/dark7lord/go-project-242/actions)

A CLI utility that calculates the size of files and directories, with support for recursive traversal and human-readable output formatting.

[![asciicast](https://asciinema.org/a/826175.svg)](https://asciinema.org/a/826175)

## Install / Uninstall

### Install

```bash
make install
# or
go install ./cmd/hexlet-path-size/
```

After installation, make sure `$(go env GOPATH)/bin` is in your `$PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Build only (binary stays in ./bin)

```bash
make build
```

### Uninstall

```bash
make uninstall
# or
go clean -i ./cmd/hexlet-path-size/
```

## Usage

```bash
hexlet-path-size [global options] <path>
```

The utility supports the following flags:

| Flag | Description |
|------|-------------|
| `-H`, `--human` | Output size in a human-readable format (e.g., `1.5MB` instead of `1536000B`). |
| `-a`, `--all` | Include hidden files and directories (starting with `.`) in the calculation. |
| `-r`, `--recursive` | Recursively calculate the size of all nested directories. |

Quick examples:

```bash
go run ./cmd/hexlet-path-size ./testdata
go run ./cmd/hexlet-path-size -H -r ./testdata
go run ./cmd/hexlet-path-size -a -H -r ./testdata
```

## Examples

Examples below assume `hexlet-path-size` is already installed (`make install` or `go install ./cmd/hexlet-path-size/`).

### Basic file or directory size

```bash
hexlet-path-size ./testdata/sample.txt
```

```text
1024B   ./testdata/sample.txt
```

### Human-readable recursive size

```bash
hexlet-path-size -H -r ./testdata
```

```text
1.5MB   ./testdata
```

### Include hidden files

```bash
hexlet-path-size -a -H -r ./testdata
```

```text
2.1MB   ./testdata
```

## Development

```bash
make test       # run all tests
make lint       # run linters
make build      # build binary to bin/hexlet-path-size
make clean      # remove build artifacts
```