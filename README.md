# gwttr

[![Build](https://github.com/alrayyes/gwttr/actions/workflows/build.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/build.yml)
[![Tests](https://github.com/alrayyes/gwttr/actions/workflows/test.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/test.yml)
[![golangci-lint](https://github.com/alrayyes/gwttr/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/golangci-lint.yml)
[![Linting](https://github.com/alrayyes/gwttr/actions/workflows/linting.yml/badge.svg)](https://github.com/alrayyes/gwttr/actions/workflows/linting.yml)
[![codecov](https://codecov.io/gh/alrayyes/gwttr/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/gwttr)
[![Release](https://img.shields.io/github/v/release/alrayyes/gwttr)](https://github.com/alrayyes/gwttr/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/gwttr.svg)](https://pkg.go.dev/github.com/alrayyes/gwttr)
[![License](https://img.shields.io/github/license/alrayyes/gwttr)](https://choosealicense.com/licenses/gpl-3.0/)

A proof-of-concept Go client for [wttr.in](https://wttr.in/), the weather
service you can curl. Run it and it prints the current conditions for Honolulu
as coloured ASCII art.

The location is fixed. There's no argument parsing yet, so `gwttr` reports on
Honolulu and nowhere else. See
[issue #210](https://github.com/alrayyes/gwttr/issues/210).

## Requirements

To run a released binary you need nothing at all. The builds are static, with
`CGO_ENABLED=0`. Linux and Windows get x86-64, i386 and arm64 archives, and
macOS gets x86-64 and arm64.

To build it yourself you need **Go 1.24 or newer**. That's the version in
`go.mod`, and it's what CI builds with.

To work on it you also need:

- **[bun](https://bun.sh/)** for the git hooks and the Markdown, YAML and JSON
  tooling. Not npm, yarn or pnpm. The lockfile is `bun.lock`.
- **[golangci-lint](https://golangci-lint.run/) 2.12.2**, which the hooks and CI
  both run. The version is pinned in
  `.github/workflows/golangci-lint.yml`.

You only need [goreleaser](https://goreleaser.com/) if you want to reproduce a
release build locally.

## Installation

Download an archive for your platform from
[the latest release](https://github.com/alrayyes/gwttr/releases/latest), unpack
it, and put the `gwttr` binary somewhere on your `PATH`.

Or build from source:

```shell
go build
```

Or install it straight into your Go bin directory:

```shell
go install github.com/alrayyes/gwttr@latest
```

## Usage

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

You can use the `wttrclient` package on its own if you just want the forecast as
a string. The
[reference](https://pkg.go.dev/github.com/alrayyes/gwttr/wttrclient) has a
runnable example.

## Contributing

Clone the repository and install the tooling:

```shell
bun install
```

That installs the git hooks too. `bun install` runs `lefthook install` through
the `prepare` script, so there's no separate setup step to forget.

After that the hooks do the work:

- **pre-commit** formats what it can and stages the result. Biome writes JSON,
  Prettier writes Markdown and YAML, and `golangci-lint fmt` writes Go.
- **commit-msg** runs [commitlint](https://commitlint.js.org/). Commits follow
  [Conventional Commits](https://www.conventionalcommits.org/), which is also
  how the next version number gets decided. `feat:` and `fix:` are load bearing,
  not decoration.
- **pre-push** checks everything again in report-only mode and runs the tests.
  Nothing at this stage writes to your files, so what you push is what you
  reviewed.

CI runs the same commands, so a green hook should mean a green pipeline.

You can run any of them by hand:

```shell
bun run lint          # Biome
bun run format:check  # Prettier, report only
bun run md:lint       # markdownlint
go test ./...
```

The tests never touch the network. Everything that speaks HTTP points at an
`httptest` server instead, so the suite passes on a train.

Work lands through a pull request. Please keep it to one change per pull
request. That's the difference between a review and a rubber stamp.

## Releases

Nobody picks a version number here.
[release-please](https://github.com/googleapis/release-please) reads the
Conventional Commits that land on `main` and keeps a release pull request open
carrying the next version and the changelog entry. Merging it tags the release.
[goreleaser](https://goreleaser.com/) then builds that tag, cross-compiles, and
attaches the archives.

Both steps live in `.github/workflows/release-please.yml`, and they're in the
same workflow on purpose. A tag pushed from inside a workflow can't be relied on
to trigger another one, so a separate on-tag workflow fails silently. The first
symptom is a release with no binaries on it.

`.release-please-manifest.json` holds the current version. If a release ever
goes wrong, correct that file rather than the tag.

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
