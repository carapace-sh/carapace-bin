# Project Layout

```sh
.
├── cmd
│  ├── carapace # main application
│  ├── carapace-fmt # simple formatter
│  ├── carapace-generate # executed by `go generate`
│  ├── carapace-lint # simple linter
│  ├── carapace-parse # simple help output parser
│  └── carapace-shim # binary for runnable specs in windows
├── completers # completers
│  └── example_completer # completer for `example`
│    └── cmd
│      └── action # local (coupled) actions
├── completers_release # optimized completers
├── dist # goreleaser dist folder
├── docs # documentation
├── internal # internal packages
└── pkg # public packages
   ├── actions # shared actions that are also exposed as macros
   │  └── tools # shared actions specific to tools
   ├── conditions # conditions for environment variable completion
   ├── env # environment variables
   ├── styles # style configurations
   └── util # util functions
```
