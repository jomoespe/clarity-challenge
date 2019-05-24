# Clarity backend code challenge

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

The source file for three programs are located in [cmd/](cmd/) folder.


### How to build

```terminal
make
```

This build the binaries `listhosts`, `parselog` and `log-generator` in the project root directory.

## Running the programs

### Parse the data with time_init and time_end (`listhosts`)

```terminal
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

```terminal
# all host related to *Aadvik* between two dates from a file.
./listhosts -start=1565647205599 -end=1565687511867 -host=Aadvik -v test/input-file-10000.txt

# all host related to *Aadvik* from standard input (piped)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik

# count all host related to *Aadvik* from standard input (piping stdin and stdout)
cat test/input-file-10000.txt | ./listhosts -host=Aadvik | wc -l
```

### Unlimited input parser (`parselog`)

```terminal
./parselog [-host=hostname] [-lapse=seconds] [FILE]
```

| Parameter        | Description                                                 |
|------------------|-------------------------------------------------------------|
| `-host=hostname` | The host to find. Default is no host                        |
| `-lapse=seconds` | Number of seconds to gererate report. Default 3600 (1 hour) |
| `FILE`           | The file to process. If no file it process standard input   |

Examples:

```terminal
# Parse a file looking for connections with 'Aadvik' host
./parselog -host=Aadvik test/input-file-10000.txt

# 
./log-generator | ./parselog -lapse=5 -host=dijkstra
```

### Ramdom log lines provider (`log-generator`)

To help running the samples, I've created an small tool
