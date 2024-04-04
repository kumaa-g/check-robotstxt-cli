# Robots.txt Checker

This project provides a command-line tool for checking if Googlebot or Google-Extended are disallowed in a given domain's robots.txt file. The tool takes a CSV file containing a list of domains as input and outputs a CSV file with the results.

## Prerequisites

- Go version 1.16 or higher

or

- Docker environment

## Installation

Clone the repository:

```bash
git clone https://github.com/kumaa-g/check-robotstxt-cli.git
```

Navigate to the project directory:

```bash
cd check-robotstxt-cli
```

## Usage

Prepare a CSV file with a list of domains you would like to check. By refering [example.csv.template](https://github.com/kumaa-g/check-robotstxt-cli/blob/main/example.csv.template)

Run the program with:

```bash
go run . path/to/yourfile.csv > results.csv
```

Replace `path/to/yourfile.csv` with the path to your CSV file. The program will create a `results.csv` file in the same directory.

## Docker

A Dockerfile is included for building a Docker image of the application.

Build the Docker image with:

```bash
docker build -t checker .
```

Run the Docker container with:

```bash
docker run --rm -v $(pwd):/root/mnt:ro checker mnt/path/to/your.csv > result.csv
```

## Testing

Run the tests with:

```bash
go test
```

## License

[LICENSE](https://github.com/kumaa-g/check-robotstxt-cli/blob/main/LICENSE)
