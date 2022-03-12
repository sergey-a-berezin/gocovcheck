# gocovcheck: Go Coverage Checker

Helper tools to summarize Go test coverage, for use in scripts.

See the [`runtests`](runtests) script in this repo for a example use of
`gocovcheck`.

## Installation

```
go install github.com/sergey-a-berezin/gocovcheck/gocovcheck@latest
go install github.com/sergey-a-berezin/gocovcheck/jsonread@latest
```

## Usage

```
go test -coverprofile <package1>.cov <package1>
go test -coverprofile <package2>.cov <package2>
...
cat *.cov > all.cov
gocovcheck all.cov 75
```

The above combines the coverage reports from all the packages and checks that
the overall coverage is at least 75%.

`gocovcheck` will return a non-zero error code if the coverage is below the
specified percentage, allowing it to be used in shell conditionals.

You may choose to keep some test configuration parameters as a JSON file,
e.g. as `test_config.json` files within each package, for instance:

```
{ "min_coverage": 75, "max_runtime": 10 }
```

You can then read the values for each key using `jsonread` tool, e.g.:

```
$ jsonread test_config.json
75
```

## Development

If you do _not_ intend to modify the code in this package, you can skip this
section.

The following describes the intended development environment and workflow for
`gocovcheck`.

### Setting up developer environment

```
git clone git@github.com:sergey-a-berezin/gocovcheck.git
cd gocovcheck
make init
```

This will:

- Setup `gopath/bin/activate` script for setting up your Go environment with
  `GOPATH=<abs.path>/gopath;`
- Download all the dependencies listed in `go.mod`;
- Install `golint`, `goconvey`, `gocovcheck` and `jsonread` into the
  environment.

Now you can run:

    ./gopath/binactivate

This starts a new shell with the Go environment. To deactivate, exit the shell:
`exit` or `<Ctrl-D>`.

### Running tests

A quick command line to run the tests:

    make test

This command is good for automated tests, e.g. if you have a continuous
integration setup, and for a summary of the total test coverage.

However, for iterative development, I highly recommend running an interactive
`goconvey` session:

```
make goconvey
```

This command will block (so it's best to run it in a separate shell) and open a
browser window visualizing all of your tests. Modifying a source file will rerun
all the tests and update the web page automatically. This way, you can just
write your code as usual, and immediately know if the code compiles and/or
passes tests.

### Contributing to `gocovcheck`

Pull requests are welcome!

Having said that, given the simple and highly focused nature of the tool, I
consider it completed, and will only support necessary maintenance, e.g. making
sure it works with the latest Go compiler version.
