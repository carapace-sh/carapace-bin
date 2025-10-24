package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["golang"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("go"),
			Variables: map[string]string{
				"AR":                    "The command to use to manipulate library archives when building with the gccgo compiler",
				"CC":                    "The command to use to compile C code",
				"CGO_CFFLAGS_ALLOW":     "Like CGO_CFLAGS_ALLOW but for the Fortran compiler",
				"CGO_CFFLAGS_DISALLOW":  "Like CGO_CFLAGS_DISALLOW but for the Fortran compiler",
				"CGO_CFFLAGS":           "Like CGO_CFLAGS but for the Fortran compiler",
				"CGO_CFLAGS_ALLOW":      "A regular expression specifying additional flags to allow to appear in #cgo CFLAGS source code directives",
				"CGO_CFLAGS_DISALLOW":   "A regular expression specifying flags that must be disallowed from appearing in #cgo CFLAGS source code directives",
				"CGO_CFLAGS":            "Flags that cgo will pass to the compiler when compiling C code",
				"CGO_CPPFLAGS_ALLOW":    "Like CGO_CFLAGS_ALLOW but for the C preprocessor",
				"CGO_CPPFLAGS_DISALLOW": "Like CGO_CFLAGS_DISALLOW but for the C preprocessor",
				"CGO_CPPFLAGS":          "Like CGO_CFLAGS but for the C preprocessor",
				"CGO_CXXFLAGS_ALLOW":    "Like CGO_CFLAGS_ALLOW but for the C++ compiler",
				"CGO_CXXFLAGS_DISALLOW": "Like CGO_CFLAGS_DISALLOW but for the C++ compiler",
				"CGO_CXXFLAGS":          "Like CGO_CFLAGS but for the C++ compiler",
				"CGO_ENABLED":           "Whether the cgo command is supported",
				"CGO_LDFLAGS_ALLOW":     "Like CGO_CFLAGS_ALLOW but for the linker",
				"CGO_LDFLAGS_DISALLOW":  "Like CGO_CFLAGS_DISALLOW but for the linker",
				"CGO_LDFLAGS":           "Like CGO_CFLAGS but for the linker",
				"CXX":                   "The command to use to compile C++ code",
				"FC":                    "The command to use to compile Fortran code",
				"GCCGO":                 "The gccgo command to run for 'go build -compiler=gccgo'",
				"GCCGOTOOLDIR":          "If set, where to find gccgo tools, such as cgo",
				"GO111MODULE":           "Controls whether the go command runs in module-aware mode or GOPATH mode",
				"GO386":                 "For GOARCH=386, how to implement floating point instructions",
				"GOAMD64":               "For GOARCH=amd64, the microarchitecture level for which to compile",
				"GOARCH":                "The architecture, or processor, for which to compile code",
				"GOARM":                 "For GOARCH=arm, the ARM architecture for which to compile",
				"GOAUTH":                "semicolon-separated list of authentication commands for go-import and HTTPS module mirror interactions",
				"GOBIN":                 "The directory where 'go install' will install a command",
				"GOCACHE":               "The directory where the go command will store cached information for reuse in future builds",
				"GOCACHEPROG":           "A command that implements the go command build cache externally",
				"GOCOVERDIR":            "Directory into which to write code coverage data",
				"GODEBUG":               "Enable various debugging facilities",
				"GOENV":                 "The location of the Go environment configuration file",
				"GOEXE":                 "The executable file name suffix (\".exe\" on Windows, \"\" on other systems)",
				"GOEXPERIMENT":          "Comma-separated list of toolchain experiments to enable or disable",
				"GO_EXTLINK_ENABLED":    "Whether the linker should use external linking mode",
				"GOFLAGS":               "A space-separated list of -flag=value settings to apply to go commands by default",
				"GOGCCFLAGS":            "A space-separated list of arguments supplied to the CC command",
				"GOHOSTARCH":            "The architecture (GOARCH) of the Go toolchain binaries",
				"GOHOSTOS":              "The operating system (GOOS) of the Go toolchain binaries",
				"GOINSECURE":            "Comma-separated list of glob patterns",
				"GOMIPS64":              "For GOARCH=mips64{,le}, whether to use floating point instructions",
				"GOMIPS":                "For GOARCH=mips{,le}, whether to use floating point instructions",
				"GOMODCACHE":            "The directory where the go command will store downloaded modules",
				"GOMOD":                 "The absolute path to the go.mod of the main module",
				"GONOPROXY":             "Comma-separated list of glob patterns",
				"GONOSUMDB":             "Comma-separated list of glob patterns",
				"GOOS":                  "The operating system for which to compile code",
				"GOPATH":                "Controls where various files are stored",
				"GOPPC64":               "For GOARCH=ppc64{,le}, the target ISA (Instruction Set Architecture)",
				"GOPRIVATE":             "Comma-separated list of glob patterns",
				"GOPROXY":               "URL of Go module proxy",
				"GOROOT_FINAL":          "The root of the installed Go tree, when it is installed in a location other than where it is built",
				"GOROOT":                "The root of the go tree",
				"GOSUMDB":               "The name of checksum database to use and optionally its public key and URL",
				"GOTMPDIR":              "The directory where the go command will write temporary source files, packages, and binaries",
				"GOTOOLCHAIN":           "Controls which Go toolchain is used",
				"GOTOOLDIR":             "The directory where the go tools (compile, cover, doc, etc...) are installed",
				"GOVCS":                 "Lists version control commands that may be used with matching servers",
				"GOVERSION":             "The version of the installed Go tree, as reported by runtime.Version",
				"GOWASM":                "For GOARCH=wasm, comma-separated list of experimental WebAssembly features to use",
				"GOWORK":                "In module aware mode, use the given go.work file as a workspace file",
				"PKG_CONFIG":            "Path to pkg-config tool",
			},
			VariableCompletion: map[string]carapace.Action{
				// TODO more flags
				"AR": bridge.ActionCarapaceBin().Split(),
				"CC": bridge.ActionCarapaceBin().Split(),
				"CGO_ENABLED": carapace.ActionStyledValuesDescribed(
					"0", "disabled", style.Carapace.KeywordNegative,
					"1", "enabled", style.Carapace.KeywordPositive,
				),
				"CXX":          bridge.ActionCarapaceBin().Split(),
				"FC":           bridge.ActionCarapaceBin().Split(),
				"GCCGO":        bridge.ActionCarapaceBin().Split(),
				"GCCGOTOOLDIR": carapace.ActionDirectories(),
				"GO111MODULE":  carapace.ActionValues("off", "on", "auto").StyleF(style.ForKeyword),
				"GOARCH":       golang.ActionArchitectures(),
				"GOAUTH": carapace.ActionMultiParts(";", func(c carapace.Context) carapace.Action {
					return carapace.Batch(
						carapace.ActionMultiPartsN(" ", 2, func(c carapace.Context) carapace.Action {
							switch len(c.Parts) {
							case 0:
								return carapace.ActionValuesDescribed("git", "Runs 'git credential fill' in dir and uses its credentials").Suffix(" ")
							default:
								return carapace.ActionDirectories()
							}
						}),
						// carapace.ActionExecutables(),
						// carapace.ActionFiles(),
						carapace.ActionStyledValuesDescribed(
							"off", "Disables authentication", style.Red,
							"netrc", "Uses credentials from NETRC or the .netrc file in your home directory", style.Default,
						),
					).ToA()
				}),
				"GOBIN":       carapace.ActionDirectories(),
				"GOCACHE":     carapace.ActionDirectories(),
				"GOCACHEPROG": bridge.ActionCarapaceBin().Split(),
				"GODEBUG":     golang.ActionGodebugKeyValues().List(","),
				"GOENV": carapace.Batch(
					carapace.ActionFiles(),
					carapace.ActionValues("off").StyleF(style.ForKeyword),
				).ToA(),
				"GOEXPERIMENT": golang.ActionExperiments().UniqueList(","),
				"GOFLAGS":      bridge.ActionCarapaceBin("go").Split(), // not entirely correct as it includes the subcommand but still helpful as that can be removed afterwards
				"GOMODCACHE":   carapace.ActionDirectories(),
				"GOOS":         golang.ActionOperatingSystems(),
				"GOSUMDB":      carapace.ActionValues("off").Style(style.Carapace.KeywordNegative),
				"GOTMPDIR":     carapace.ActionDirectories(),
				"GOTOOLCHAIN": carapace.ActionMultiPartsN("+", 2, func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return golang.ActionVersions().Prefix("go")
					default:
						return carapace.ActionValues("auto").StyleF(style.ForKeyword)
					}
				}),
				"GOVCS": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues("public", "private").StyleF(style.ForKeyword).Suffix(":")
					default:
						return carapace.ActionValuesDescribed(
							"all", "",
							"bzr", "Bazaar",
							"fossil", "Fossil",
							"git", "Git",
							"hg", "Mercurial",
							"off", "",
							"svn", "Subversion",
						).StyleF(style.ForKeyword).UniqueList("|")
					}
				}).List(","),
				"GOWORK":     carapace.ActionFiles(),
				"PKG_CONFIG": carapace.ActionFiles(),
			},
		}
	}
}
