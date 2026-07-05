---
name: go
description: >
  Use when working with the Go CLI (go command) — building, testing, modules,
  workspaces, environment, toolchain, generate, vet, coverage, profiling, fuzzing,
  cross-compilation, and internal tool commands. Triggers on: "go", "go build",
  "go test", "go mod", "go work", "go generate", "go vet", "go tool", "go get",
  "go install", "go run", "go clean", "go list", "go env", "go doc", "go fmt",
  "go fix", "go version", "go telemetry", "GOMODCACHE", "GOPROXY", "GOSUMDB",
  "GOPRIVATE", "GOTOOLCHAIN", "GOOS", "GOARCH", "CGO_ENABLED", "build constraint",
  "build tag", "go.mod", "go.sum", "go.work", "vendor", "pprof", "cover",
  "test2json", "buildjson", "GODEBUG", "GOEXPERIMENT", "minimal version selection",
  "MVS", "module proxy", "toolchain", "fuzzing", "race detector".
user-invocable: true
---

# Go CLI In-Depth Reference

Comprehensive reference for the `go` command-line tool — the driver for building, testing, module management, code generation, and toolchain operations in Go. Covers the module system, build flags, test framework, environment variables, workspaces, and the bundled `go tool` subcommands.

## Data Flow

```
Source code (packages)
  → go build / go install (compile + link)
    → build cache (GOCACHE)
      → executable / archive

go.mod (module definition)
  → go get / go mod tidy (resolve dependencies)
    → module proxy (GOPROXY) / VCS (GOVCS)
      → module cache (GOMODCACHE)
        → go.sum (checksum verification, GOSUMDB)

go test (compile test binary)
  → go vet (automatic, subset)
    → test binary (pkg.test)
      → test cache (package list mode)
        → coverage profile / pprof profile / trace
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| concept, model, module, package, import path, GOPATH, GOROOT, GO111MODULE, module-aware mode, GOPATH mode, package pattern, `all`, `std`, `cmd`, `tool`, wildcard, `...`, reserved package names, main package, test package, `_test` suffix, testdata, vendor, file types, `.go`, `.c`, `.s`, `.syso`, source file naming, `*_GOOS`, `*_GOARCH`, implicit build constraint | [references/concepts.md](references/concepts.md) |
| go build, build flags, `-o`, `-a`, `-n`, `-x`, `-p`, `-race`, `-msan`, `-asan`, `-cover`, `-covermode`, `-coverpkg`, `-v`, `-work`, `-asmflags`, `-buildmode`, `-buildvcs`, `-compiler`, `-gccgoflags`, `-gcflags`, `-installsuffix`, `-json`, `-ldflags`, `-linkshared`, `-mod`, `-modcacherw`, `-modfile`, `-overlay`, `-pgo`, `-pkgdir`, `-tags`, `-trimpath`, `-toolexec`, build mode, archive, c-archive, c-shared, shared, exe, pie, plugin, PGO, profile-guided optimization, trimpath, cross-compilation, GOOS, GOARCH, build constraint, build tag, `//go:build`, `// +build`, buildjson, BuildEvent, shared build flags, `-C` | [references/build.md](references/build.md) |
| go test, test flag, `-run`, `-skip`, `-bench`, `-benchtime`, `-count`, `-failfast`, `-fullpath`, `-fuzz`, `-fuzztime`, `-fuzzminimizetime`, `-json`, `-list`, `-outputdir`, `-parallel`, `-short`, `-shuffle`, `-timeout`, `-v`, `-vet`, `-args`, `-c`, `-exec`, `-o`, `-cover`, `-covermode`, `-coverpkg`, `-coverprofile`, `-benchmem`, `-blockprofile`, `-blockprofilerate`, `-cpuprofile`, `-memprofile`, `-memprofilerate`, `-mutexprofile`, `-mutexprofilefraction`, `-trace`, `-artifacts`, test2json, TestEvent, cacheable flags, local directory mode, package list mode, test caching, `-count=1`, fuzzing, seed corpus, testdata, fuzz cache, test binary, `-test.` prefix, testing functions, TestXxx, BenchmarkXxx, FuzzXxx, ExampleXxx, `t.Parallel`, `-race`, race detector, sanitizer | [references/test.md](references/test.md) |
| go mod, mod init, mod tidy, mod download, mod edit, mod graph, mod vendor, mod verify, mod why, go.mod, go.sum, module path, module version, semver, semantic versioning, require, replace, exclude, retract, tool, ignore, godebug, indirect dependency, lazy module loading, minimal version selection, MVS, version query, `@latest`, `@none`, `@patch`, `@v1.2.3`, module cache, GOMODCACHE, modcacherw, vendor directory, `vendor/modules.txt`, `-mod=vendor`, `-mod=readonly`, `-mod=mod`, module-auth, go.sum, checksum, GOSUMDB, GOPRIVATE, GONOSUMDB, GONOPROXY, GOINSECURE, module proxy, GOPROXY, goproxy protocol, VCS, GOVCS, direct download, zip, `.info`, `.mod`, `.zip` | [references/mod.md](references/mod.md) |
| go tool, tool command, asm, cgo, compile, cover, fix, link, preprofile, vet, `go tool vet`, `go tool cover`, `go tool pprof`, `go tool compile`, `go tool link`, `go tool asm`, `go tool cgo`, `-n`, toolexec, tool definition, `tool` directive in go.mod, `go tool` resolution, builtin tools, additional tools, cmd/vet, vet analyzer, `-vettool`, vet checks, printf, atomic, bool, buildtags, directive, errorsas, ifaceassert, nilfunc, stringintconv, tests, shadow analyzer, `-fix`, `-diff`, `go tool fix`, fixer, hostport | [references/tool.md](references/tool.md) |
| go env, environment variable, GOARCH, GOOS, GOPATH, GOROOT, GOBIN, GOCACHE, GOCACHEPROG, GODEBUG, GOENV, GOFLAGS, GOINSECURE, GOMODCACHE, GOPRIVATE, GONOPROXY, GONOSUMDB, GOPROXY, GOSUMDB, GOTMPDIR, GOTOOLCHAIN, GOVCS, GOWORK, GO111MODULE, GOAUTH, GOEXPERIMENT, GOFIPS140, GOAMD64, GOARM, GOARM64, GO386, GOMIPS, GOMIPS64, GOPPC64, GORISCV64, GOWASM, CGO_ENABLED, CGO_CFLAGS, CGO_CPPFLAGS, CGO_CXXFLAGS, CGO_FFLAGS, CGO_LDFLAGS, CC, CXX, FC, AR, PKG_CONFIG, GOEXE, GOGCCFLAGS, GOHOSTARCH, GOHOSTOS, GOMOD, GOTELEMETRY, GOTELEMETRYDIR, GOTOOLDIR, GOVERSION, `go env -w`, `go env -u`, `go env -json`, `go env -changed`, env config file, GOENV | [references/environment.md](references/environment.md) |
| go work, work init, work use, work sync, work edit, work vendor, go.work, workspace, use directive, replace directive, go directive, toolchain directive, godebug directive, GOWORK, workspace mode, multi-module, build list, sync, `-r` flag, GoWork JSON, Use struct, Replace struct, workspace vendor | [references/work.md](references/work.md) |
| go generate, generate directive, `//go:generate`, generator, code generation, `$GOFILE`, `$GOLINE`, `$GOPACKAGE`, `$GOARCH`, `$GOOS`, `$GOROOT`, `$DOLLAR`, `$PATH`, `-command`, alias, `-run`, `-skip`, generated file convention, `DO NOT EDIT`, `generate` build tag, go get, go install, version suffix, `@latest`, `@v1.0.0`, module-aware install, GOBIN, go clean, clean flags, `-cache`, `-testcache`, `-modcache`, `-fuzzcache`, `-i`, `-r`, go run, `-exec`, go doc, doc flags, `-all`, `-src`, `-short`, `-u`, `-cmd`, `-http`, go fmt, gofmt, go fix, fix tool, go vet, vet flags, `-vettool`, `-c`, `-json`, go version, `-m`, `-v`, go telemetry, telemetry mode, go bug | [references/cli.md](references/cli.md) |

## Quick Guide

- **What are Go's core concepts (modules, packages, GOPATH)?** → [references/concepts.md](references/concepts.md)
- **How do I build/compile Go code?** → [references/build.md](references/build.md)
- **What build flags are available?** → [references/build.md](references/build.md)
- **How do build constraints/tags work?** → [references/build.md](references/build.md)
- **How do I run tests?** → [references/test.md](references/test.md)
- **What test flags are available?** → [references/test.md](references/test.md)
- **How does test caching work?** → [references/test.md](references/test.md)
- **How do I do fuzzing?** → [references/test.md](references/test.md)
- **How do coverage and profiling work?** → [references/test.md](references/test.md)
- **How do modules and go.mod work?** → [references/mod.md](references/mod.md)
- **How do I add/remove dependencies?** → [references/mod.md](references/mod.md)
- **How does the module proxy and checksum work?** → [references/mod.md](references/mod.md)
- **What tools are available via `go tool`?** → [references/tool.md](references/tool.md)
- **How does `go vet` work?** → [references/tool.md](references/tool.md)
- **What environment variables does the go command use?** → [references/environment.md](references/environment.md)
- **How do I set Go env vars persistently?** → [references/environment.md](references/environment.md)
- **How do workspaces work?** → [references/work.md](references/work.md)
- **How does `go generate` work?** → [references/cli.md](references/cli.md)
- **How do I install a tool or binary?** → [references/cli.md](references/cli.md)
- **How do I clean caches?** → [references/cli.md](references/cli.md)
- **How do I inspect a Go binary's version?** → [references/cli.md](references/cli.md)
- **How do I cross-compile?** → [references/build.md](references/build.md) and [references/environment.md](references/environment.md)
- **How do I use a custom Go toolchain?** → [references/environment.md](references/environment.md)
- **How do I vendor dependencies?** → [references/mod.md](references/mod.md)
- **How do I debug the build cache?** → [references/environment.md](references/environment.md) and [references/build.md](references/build.md)

## Cross-Project References

- For cobra CLI development (commands, flags, completions), use the **cobra** skill (in this repo).
- For carapace library development (Action API, traverse engine, shell formatters), use the **carapace-dev** skill (in this repo).
- For carapace user-facing integration (specs, macros, setup), use the **carapace** skill (in this repo).
- For Go testing framework internals (`testing.T`, `testing.B`, `testing.F`), see the [Go testing package docs](https://pkg.go.dev/testing) — this skill covers the `go test` command, not the testing API.
- For pprof usage details beyond `go tool pprof`, see the [pprof docs](https://pkg.go.dev/runtime/pprof).
