# CLAUDE.md

## Project Overview

`pangu.go` is a text spacing library that automatically inserts whitespace between CJK (Chinese, Japanese, Korean) characters and half-width characters (alphabetical letters, numerical digits, and symbols). It follows the pangu.js v10 engine, shipped as a zero-dependency Go module with a package API and a CLI.

## Common Development Commands

```bash
go test -race -shuffle=on ./...                     # Run all tests
go test -run 'TestSymbolSlash' -v                   # Run tests matching a name
gofmt -l . && go vet ./... && go fix -diff ./...    # Lint, as CI does
go fix ./...                                        # Apply modernize fixes
```

## Where Things Live

- Spacing engine and package API: `*.go` in the repo root (`package pangu`)
- CLI: `cmd/pangu/`
- Parity spec, ported 1:1 from pangu.js `tests/shared/`: `*_test.go` in the repo root (text fixtures in `testdata/`)
- Domain language and algorithm semantics: `../pangu.js/CONTEXT.md`; decision records: `../pangu.js/docs/adr/`

## Gotchas

- Spacing rule changes flow downstream from pangu.js: fix or tweak rules upstream first, then port. Read `CONTEXT.md` before touching the engine.
- Engine internals are idiomatic Go, not js-shaped; output parity is locked by the 1:1 ported test suite. Port each pangu.js `expect(...)` as one table case, run as a `t.Run` subtest named by its input.
- Go's `regexp` is RE2: no lookahead, lookbehind, or backreferences. Port a lookaround by capturing the context and writing it back, or by a check in code. A capture consumes what a lookaround only peeked at, so matches sharing a character stop overlapping: `([CJK])([A])([CJK])` turns `中a中b中` into `中 a 中b中`. Test the shared-character case for every translated rule.
- Go's `\s` is ASCII `[\t\n\f\r ]`; js's also matches `\v`, NBSP, `　`, and ` `. Spell out the js set when porting `\s` or `[^\s...]`.
- js string lengths and offsets count UTF-16 code units; Go's count UTF-8 bytes. Port `text.length`, `slice()`, and callback offsets rune-aware, not with `len()`.
- Commented-out cases, FIXME comments, and `t.Skip` in tests are intentional 1:1 ports of upstream FIXME cases and `it.todo`/`describe.todo`: leave them.
- Standard library only. A module, `tool` directive included, lands in every importer's module graph, so adding one needs the user's approval. Keep the `go` directive at 1.26, and check the go.mod diff after `go get`.
- Write code comments in English with ASCII characters only. Never paste CJK sample text into a comment; describe the shape generically (`CJK | CJK`, `A+CJK`) and use `\uXXXX` escape notation when a specific character matters.

## External Tool Documentation

Pre-resolved Context7 IDs for the `find-docs` skill. Pass them to `ctx7 docs` and skip `ctx7 library`:

| Tool           | `libraryId`                   |
| -------------- | ----------------------------- |
| Go             | `/golang/go`                  |
| Go docs        | `/websites/go_dev_doc`        |
| GitHub Actions | `/websites/github_en_actions` |
