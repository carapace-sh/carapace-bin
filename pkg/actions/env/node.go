package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	_bool := carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)
	knownVariables["node"] = variables{
		Condition: checkPath("node"),
		Names: map[string]string{
			"NODE_DEBUG":                   "Comma-separated list of core modules that should print debug information",
			"NODE_DEBUG_NATIVE":            "Comma-separated list of C++ core modules that should print debug information",
			"NODE_DISABLE_COLORS":          "When set to 1, colors will not be used in the REPL",
			"NODE_EXTRA_CA_CERTS":          "When set, the well-known “root” CAs will be extended with the extra certificates in file",
			"NODE_ICU_DATA":                "Data path for ICU (Intl object) data",
			"NODE_NO_WARNINGS":             "When set to 1, process warnings are silenced",
			"NODE_OPTIONS":                 "A space-separated list of command-line options",
			"NODE_PATH":                    "A colon-separated list of directories prefixed to the module search path",
			"NODE_PENDING_DEPRECATION":     "When set to 1, emit pending deprecation warnings",
			"NODE_PRESERVE_SYMLINKS":       "When set to 1, the module loader preserves symbolic links when resolving and caching modules",
			"NODE_REDIRECT_WARNINGS":       "Write process warnings to the given file instead of printing to stderr",
			"NODE_REPL_HISTORY":            "Path to the file used to store persistent REPL history",
			"NODE_REPL_EXTERNAL_MODULE":    "Path to a Node.js module which will be loaded in place of the built-in REPL",
			"NODE_SKIP_PLATFORM_CHECK":     "When set to 1, the check for a supported platform is skipped during Node.js startup",
			"NODE_TLS_REJECT_UNAUTHORIZED": "When set to 0, TLS certificate validation is disabled",
			"NODE_V8_COVERAGE":             "When set, Node.js writes JavaScript  code  coverage information to dir",
		},
		Completion: map[string]carapace.Action{
			// TODO more completions
			"NODE_DISABLE_COLORS":          _bool,
			"NODE_NO_WARNINGS":             _bool,
			"NODE_PATH":                    carapace.ActionDirectories().List(":"),
			"NODE_PENDING_DEPRECATION":     _bool,
			"NODE_SKIP_PLATFORM_CHECK":     _bool,
			"NODE_TLS_REJECT_UNAUTHORIZED": _bool,
		},
	}

}
