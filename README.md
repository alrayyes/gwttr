# gwttr

[![Build](https://github.com/alrayyes/gwttr/actions/workflows/build.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/build.yml)
[![Tests](https://github.com/alrayyes/gwttr/actions/workflows/test.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/test.yml)
[![golangci-lint](https://github.com/alrayyes/gwttr/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/golangci-lint.yml)
[![Linting](https://github.com/alrayyes/gwttr/actions/workflows/linting.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/linting.yml)
[![Codecov](https://codecov.io/gh/alrayyes/gwttr/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/gwttr)
[![Release](https://img.shields.io/github/v/release/alrayyes/gwttr)](https://github.com/alrayyes/gwttr/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/gwttr.svg)](https://pkg.go.dev/github.com/alrayyes/gwttr)
[![Licence](https://img.shields.io/github/license/alrayyes/gwttr)](https://choosealicense.com/licenses/gpl-3.0/)

A proof-of-concept Go client for [wttr.in](https://wttr.in/), the weather
service you can curl. Give it a location, and it prints the current conditions
there as coloured ASCII art.

## Requirements

To run a released binary you need nothing at all. The builds are static, with
`CGO_ENABLED=0`. Linux and Windows get x86-64, i386 and arm64 archives, and
macOS gets x86-64 and arm64.

To build it yourself you need **Go 1.24 or newer**. That's the version in
`go.mod`, and it's what CI builds with.

Working on it needs more than that — the tooling list is in
[CONTRIBUTING.md](CONTRIBUTING.md).

## Installation

Download an archive for your platform from
[the latest release](https://github.com/alrayyes/gwttr/releases/latest), unpack
it, and put the `gwttr` binary somewhere on your `PATH`.

Or build from source:

```shell
go build ./cmd/gwttr
```

Or install it straight into your Go bin directory:

```shell
go install github.com/alrayyes/gwttr/cmd/gwttr@latest
```

## Usage

With no arguments it reports on Honolulu:

```shell
./gwttr
```

```text
Weather report: honolulu

   \  /       Partly cloudy
 _ /"".-.     +29(30) °C
   \_(   ).   ↙ 30 km/h
   /(___(__)  16 km
               0.0 mm
```

Give it a location for anywhere else. Quote anything with a space in it:

```shell
./gwttr berlin
./gwttr "new york"
```

`./gwttr --help` lists the flags, and `./gwttr --version` reports the release
you're running.

This is a command, not a library. The packages behind it live under `internal/`,
so they're documented for whoever works on the repo rather than published for
anyone to import.

## Contributing

Everything about working on this — the tooling, the git hooks, how the checks
run and how a release is cut — is in [CONTRIBUTING.md](CONTRIBUTING.md). Short
version: `bun install`, branch, one change per pull request, and commit under
[Conventional Commits](https://www.conventionalcommits.org/), because those
subjects pick the next version number.

## Licence

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
