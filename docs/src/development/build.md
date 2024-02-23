# Build

## Requirements

- [Go](https://go.dev/)
- Ensure `PATH` contains your `~/go/bin` folder (see [doc](https://go.dev/doc/code#Command))


## Development

```sh
cd cmd/carapace
go generate ./... # only needed once and when there are new completers or actions
go install
```

![](./build-development.cast)

## Release

A build with the `release` [tag] uses `completers_release` which optimizes startup times.

```sh
cd cmd/carapace
go generate ./... # always needed for `completers_release` to be updated
go install -ldflags="-s -w" -tags release
```

![](./build-release.cast)

## Separate

Completers can also be built (and thus maintained) separately.

```sh
cd completers/ln_completer
go install -ldflags="-s -w"
ln_completer _carapace
```

![](./build-separate.cast)
