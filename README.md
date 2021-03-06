# Clarity backend code challenge

[![Build Status](https://travis-ci.com/jomoespe/clarity-challenge.svg?token=pvtAthG3oqWcKLGsBRBA&branch=master)](https://travis-ci.com/jomoespe/clarity-challenge)

This project is my implementation of [Clarity](https://clarity.ai/) [blackend code challlenge](./docs/clarity_code_challenge.pdf).

## Solution description

Based on requirements I've created two programs:

- `listhosts` to implement first requirement: *Parse the data with time_init and time_end*.
- `parselog` to implement second requirement: *Unlimited input parser*.

Also, I've creatd an utility, `log-generator` for testing purposes. It generates a log line (`timestamp source-source target-host`) every 100 milliseconds in standard output.

The project structure is based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

### Why Go languaje?

The requirements are about processing files, so, generating a native application looks like the more natural way to implement it. This ease the integration with other tools, allowing orchestration, piping with other utilities, etc; let's say *as much POSIX as possible*.

Also, taking into account that are like a *system applications*, using a language optimized for this things will help to improve performance and resource consumption.

## Build the project

### Requirements

- [Go 1.11+](https://golang.org/), as programming language.
- [GNU Make](https://www.gnu.org/software/make/) as build automation tool.

### Project struxcrture

The source file for three programs are located in [cmd/](cmd/) folder. There are an main package for each program: `listhost`, `parselog`and `logsupplier`.

There are two code packages:

- `logparser` with common behaviour, line *opening the file reader* or *parsing a log line*.
- `types` with the data types used by program (`set` and `hostconnections`)

### How to build

```bash
make
```

This build the binaries `listhosts`, `parselog` and `log-generator` in the project root directory.

## Running the programs

### Parse the data with time_init and time_end (`listhosts`)

Print a list of hostnames connected to the given host during the given period

```bash
./listhosts [-start=time_init] [-end=time_end] [-host=hostname] [-v] [-h] [FILE]
```

| Parameter          | Description                                               |
|--------------------|-----------------------------------------------------------|
| `-start=time_init` | The start date. Defaul is from begining                   |
| `-end=time_end`    | The end date. Default is to the end                       |
| `-host=hostname`   | The host to find. Default is all host                     |
| `-v`               | Print errors and warnings in stardard error               |
| `-h`               | Show command line arguments and exit                      |
| `FILE`             | The file to process. If no file it process standard input |

Examples:

```bash
# all host related to *Aadvik* between two dates from a file.
./listhosts -start=1565647205599 -end=1565687511867 -host=Aadvik -v test/input-file-10000.txt

# all host related to *Aadvik* from standard input (piped)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik

# count all host related to *Aadvik* from standard input (piping stdin and stdout)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik | wc -l
```

### Unlimited input parser (`parselog`)

Process a log file and, for a period of time, reports:

- a list of hostnames connected to the given host.
- a list of hostnames received connections from given host.
- the hostname that generated most connections.

```bash
./parselog [-host=hostname] [-lapse=seconds] [FILE]
```

| Parameter        | Description                                                 |
|------------------|-------------------------------------------------------------|
| `-host=hostname` | The host to find. Default is no host                        |
| `-lapse=seconds` | Number of seconds to gererate report. Default 3600 (1 hour) |
| `FILE`           | The file to process. If no file it process standard input   |

Examples:

```bash
# Parse a file looking for connections with 'Aadvik' host
./parselog -host=Aadvik test/input-file-10000.txt

# Parse logs from input stream, generating report each 10 seconds for host dijkstra
./log-generator | ./parselog -lapse=10 -host=dijkstra
```

### Ramdom log lines provider (`log-generator`)

To help running the samples, I've created an small tool that generates log lines in the stardanr output each 100 Milliseconds.

| Parameter             | Description                                                 |
|-----------------------|-------------------------------------------------------------|
| `-delay=milliseconds` | Number of milliseconds to wait between log line generation  |

```bash
./log-generator [-delay=milliseconds]
```
