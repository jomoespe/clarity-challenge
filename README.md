# Clarity backend code challenge

This project is a Go implementation of [Clarity](https://clarity.ai/) blackend code challlenge. [The document](./docs/clarity_code_challenge.pdf) in the doc/ directory.

## Requirements

- [Go SDK](https://golang.org/), as programming language.
- [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

### How to build

```terminal
make
```

This build the binaries `parselog` and `parselogd` in the project root directory.

## Parse the data with time_init and time_end (`parselog`)

```terminal
./parselog <time_init> <time_end> <hostname> [-file=<log_filename>] [--verbose]
```

## Unlimited input parser (`parselogd`)

> TBD

```terminal
./parselogd
```
