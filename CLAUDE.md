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
- CLI: `internal/cli/`, run by `cmd/pangu/`, with `pangu.js` `tests/node/` ported to `cli_test.go` (text fixtures in `internal/cli/testdata/`). Deviations from the js CLI: parsed by `flag`, so usage errors exit 2 and text starting with `-` needs `--` first; stdin from `/dev/null` counts as a terminal, because the standard library has no isatty
- Parity spec, ported 1:1 from pangu.js `tests/shared/`: `symbol_test.go` and `text_test.go`, one `TestXxx` per pangu.js test file (`TestSymbolPeriod` is `symbol-period.test.ts`); helpers and API tests in `pangu_test.go`
- Domain language and algorithm semantics: `../pangu.js/CONTEXT.md`; decision records: `../pangu.js/docs/adr/`

## Gotchas

- Spacing rule changes flow downstream from pangu.js: fix or tweak rules upstream first, then port. Read `CONTEXT.md` before touching the engine.
- Engine internals are idiomatic Go, not js-shaped; output parity is locked by the 1:1 ported test suite. Port each pangu.js `expect(...)` as one table case, run as a `t.Run` subtest named by its input.
- Go's `regexp` is RE2: no lookahead, lookbehind, or backreferences. Port a lookaround as the `accept` check of `replaceLookaround`, which retries a rejected match one character later the way js scans. A capture would consume what the lookaround only peeked at, so matches sharing a character stop overlapping: `([CJK])([A])([CJK])` turns `中a中b中` into `中 a 中b中`. Capture and write back only a Private Use Area placeholder marker, which never starts another match. Add each new lookaround rule's shared-character case to `TestSpaceTextSharedCharacter`.
- Go's `\s` is ASCII `[\t\n\f\r ]`; js's also matches `\v`, NBSP, `\u3000`, and `\u2028` (`jsSpace` in `rules.go`). Use `jsSpace` when porting `\s` or `[^\s...]`.
- `SpaceText` skips a rule group when the text has none of the group's trigger characters (the `strings.Contains`/`ContainsAny` checks); pangu.js runs every rule. When porting a new rule or a change to a rule's character class, keep each check listing every character its rules need. A missing character makes the rule silently skip that input.
- js string lengths and offsets count UTF-16 code units; Go's count UTF-8 bytes. Port `text.length`, `slice()`, and callback offsets rune-aware, not with `len()`.
- Commented-out cases, FIXME comments, and `testSpaceTextFails` groups in tests are intentional 1:1 ports of upstream FIXME cases and `it.fails`: leave them. `testSpaceTextFails` fails once every case in its group passes; then move the group to `testSpaceText`.
- Standard library only. A module, `tool` directive included, lands in every importer's module graph, so adding one needs the user's approval. Keep the `go` directive at 1.26, and check the go.mod diff after `go get`.
- Write code comments in English with ASCII characters only. Never paste CJK sample text into a comment; describe the shape generically (`CJK | CJK`, `A+CJK`) and use `\uXXXX` escape notation when a specific character matters.

## External Tool Documentation

Pre-resolved Context7 IDs for the `find-docs` skill. Pass them to `ctx7 docs` and skip `ctx7 library`:

| Tool           | `libraryId`                   |
| -------------- | ----------------------------- |
| Go             | `/golang/go`                  |
| Go docs        | `/websites/go_dev_doc`        |
| GitHub Actions | `/websites/github_en_actions` |
