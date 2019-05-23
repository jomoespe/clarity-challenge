# Clarity backend code challenge

This project is a Go implementation of [Clarity](https://clarity.ai/) blackend code challlenge. [The document](./docs/clarity_code_challenge.pdf) in the doc/ directory.

The project structure is based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Solution description

Based on requirements I've created two programs:

- `listhosts`, to implement first requirement: *Parse the data with time_init and time_end*.
- `parselog`, to implement second requirement: *Unlimited input parser*.

## Build the project

### Requirements

- [Go 1.11+](https://golang.org/), as programming language.
- [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

### How to build

```terminal
make
```

This build the binaries `listhosts` and `parselog` in the project root directory.

## Running the programs

### Parse the data with time_init and time_end (`listhosts`)

```terminal
./listhosts [-start=<time_init>] [-end=<time_end>] [-host=<hostname>] [-v] [-h] [FILE]
```

All parameters are optional.

| `-start=<time_init>` | The start date. Defaul is from begining                   |
| `-end=<time_end>]`   | The end date. Default is to the end                       |
| `-host=<hostname>`   | The host to find. Default is all host                     |
| `-v`           | Print errors and warnings in stardard error               |
| `-h`                 | Show command line arguments and exit                      |
| `FILE`               | The file to process. If no file it process standard input |

Examples:

```terminal
# all host related to *Aadvik* between two dates from a file.
./listhosts -start=1565647205599 -end=1565687511867 -host=Aadvik -v test/input-file-10000.txt

# all host related to *Aadvik* from standard input (piped)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik

# count all host related to *Aadvik* from standard input (piped)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik | wc -l
```

### Unlimited input parser (`parselog`)

> TBD

```terminal
./parselog [--host=<hostname>] [FILE]
```
