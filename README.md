# Clarity backend code challenge

The project uses [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

## Requirements

- [Go SDK](https://golang.org/), as programming language.
- [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

### How to build

```terminal
make
```

this will generate the binaries `parselog` and `parselogd` in the project root directory.

## Parse the data with time_init and time_end

### How to run

```terminal
./parselog <time_init> <time_end> <hostname> [-file=<log_filename>] [--verbose]
```

## Unlimited input parser

> TBD

```terminal
./parselogd
```
