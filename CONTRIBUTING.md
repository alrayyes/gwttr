# Contributing

The [README](README.md) is for whoever runs `gwttr`. This is for whoever
changes it.

## What you need

- **Go 1.24 or newer**, the version in `go.mod` and the one CI builds with.
- **[bun](https://bun.sh/)** for the git hooks and the Markdown, YAML and JSON
  tooling. Not npm, yarn or pnpm. The lockfile is `bun.lock`.
- **[golangci-lint](https://golangci-lint.run/) 2.12.2**, which the hooks and CI
  both run. The version is pinned in `.github/workflows/golangci-lint.yml`.
  Install that version rather than whichever is current: when the two disagree,
  the hook passes and the pipeline fails, and the failure doesn't say why.
- **[Vale](https://vale.sh/)** for prose style. The pre-commit hook runs it, so
  a commit touching Markdown needs it on your `PATH`. Run `vale sync` once after
  cloning to fetch the style packages, which aren't committed.

[GoReleaser](https://goreleaser.com/) is only needed if you want to reproduce a
release build locally.

## Getting set up

```shell
bun install
```

That installs the git hooks too. `bun install` runs `lefthook install` through
the `prepare` script, so there's no separate setup step to forget.

## The hooks

- **pre-commit** formats what it can and stages the result. Biome writes JSON,
  Prettier writes Markdown and YAML, and `golangci-lint fmt` writes Go.
- **commit-msg** runs [commitlint](https://commitlint.js.org/). Commits follow
  [Conventional Commits](https://www.conventionalcommits.org/), which is also
  how the next version number gets decided. `feat:` and `fix:` are load bearing,
  not decoration.
- **pre-push** checks everything again in report-only mode and runs the tests.
  Nothing at this stage writes to your files, so what you push is what you
  reviewed.

CI runs the same commands, so a green hook should mean a green pipeline. The one
check that isn't in a hook is LTeX, the grammar tier, which runs in
[`prose.yml`](.github/workflows/prose.yml).

You can run any of them by hand:

```shell
bun run lint          # Biome
bun run format:check  # Prettier, report only
bun run md:lint       # markdownlint
go test ./...
```

The tests never reach the network. Everything that speaks HTTP points at a local
`httptest` server instead, so the suite passes on a train.

## Pull requests

Work lands through a pull request. Please keep it to one change per pull
request. That's the difference between a review and a rubber stamp.

## Releases

Nobody picks a version number here.
[release-please](https://github.com/googleapis/release-please) reads the
Conventional Commits that land on `main` and keeps a release pull request open
carrying the next version and the changelog entry. Merging it tags the release.
[GoReleaser](https://goreleaser.com/) then builds that tag, cross-compiles, and
attaches the archives.

Both steps live in `.github/workflows/release-please.yml`, and they're in the
same workflow on purpose. A tag pushed from inside a workflow can't be relied on
to trigger another one, so a separate on-tag workflow fails silently. The first
symptom is a release with no binaries on it.

`.release-please-manifest.json` holds the current version. If a release ever
goes wrong, correct that file rather than the tag.
