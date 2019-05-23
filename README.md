# Clarity backend code challenge

This project is a Go implementation of [Clarity](https://clarity.ai/) blackend code challlenge. [The document](./docs/clarity_code_challenge.pdf) in the doc/ directory.

The project structure is based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Solution description

Based on requirements I've created two binaries:

- `listhosts`, which reads a log file and list of hostnames connected to the given host during the given period.
- `parselog`, 

## Requirements

- [Go 1.11+](https://golang.org/), as programming language.
- [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

### How to build

```terminal
make
```

This build the binaries `listhosts` and `parselog` in the project root directory.

To check/verify project dependencies dependencies:

```terminal
make dependencies
```

## Parse the data with time_init and time_end (`listhosts`)

```terminal
./listhosts --start=<time_init> --end=<time_end> --host=<hostname> [-file=<log_filename>] [--verbose]
```

## Unlimited input parser (`parselog`)

> TBD

```terminal
./parselog [--host=<hostname>] [FILE]
```
