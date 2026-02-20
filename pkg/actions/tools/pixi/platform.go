package pixi

import "github.com/carapace-sh/carapace"

// ActionPlatforms completes platform names
//
//	linux-64
//	osx-arm64
func ActionPlatforms() carapace.Action {
	return actionExecPixi("info", "--json")(func(output []byte) carapace.Action {
		info, err := parseInfo(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		seen := make(map[string]bool)
		vals := make([]string, 0)
		for _, env := range info.EnvironmentsInfo {
			for _, p := range env.Platforms {
				if !seen[p] {
					seen[p] = true
					vals = append(vals, p)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("platforms")
}

// ActionKnownPlatforms completes known platform names
//
//	linux-64
//	osx-arm64
func ActionKnownPlatforms() carapace.Action {
	return carapace.ActionValues(
		"emscripten-wasm32",
		"freebsd-64",
		"linux-32",
		"linux-64",
		"linux-aarch64",
		"linux-armv6l",
		"linux-armv7l",
		"linux-ppc64",
		"linux-ppc64le",
		"linux-riscv32",
		"linux-riscv64",
		"linux-s390x",
		"osx-64",
		"osx-arm64",
		"wasi-wasm32",
		"win-32",
		"win-64",
		"win-arm64",
	).Tag("platforms")
}
