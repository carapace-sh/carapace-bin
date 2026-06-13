package completer

import (
	"maps"
	"os"
	"runtime"
	"slices"
	"sort"
	"strings"
)

type Completers []Completer

func (c Completers) Lookup(variant, group string) (*Completer, bool) {
	sort.Sort(c) // TODO this modifies c
	for _, completer := range c {
		if (variant == "" || completer.Variant == variant) && (group == "" || completer.Group == group) {
			return &completer, true
		}
	}
	return nil, false
}

func (c Completers) Len() int { return len(c) }
func (c Completers) Less(i, j int) bool { // TODO this needs testing (and likely some fixes)
	// TODO needs tests and is likely wrong
	if b := strings.Compare(c[i].Name, c[j].Name) != 0; b {
		return b
	}

	// choices have highest priority
	if a, b := c[i].Choice, c[j].Choice; a != b {
		return a > b
	}

	// TODO shells? more specific stuff
	groupPriority := map[string]int{
		"user":   -41, // user specs
		"system": -40, // system specs

		// CARAPACE_SHELL: -35 (see below)

		"bash":       -30,
		"bash-ble":   -29,
		"cmd-clink":  -28,
		"elvish":     -27,
		"fish":       -26,
		"nushell":    -25,
		"oil":        -24,
		"powershell": -23,
		"tcsh":       -22,
		"xonsh":      -21,
		"zsh":        -20,

		// runtime.GOOS: -15 (see below)

		"linux": -10,
		"unix":  -9,

		"darwin":  -8,
		"freebsd": -7,
		"netbsd":  -6,
		"openbsd": -5,
		"bsd":     -4,

		"windows": -3,
		"android": -2,
		"common":  -1,
		// TODO support pseudo os 'termux'?
		"bridge": 1, // lower priority than anything internal
	}

	switch b := os.Getenv("CARAPACE_SHELL"); b { // TODO public access of env in carapace
	case "bash",
		"bash-ble",
		"cmd-clink",
		"elvish",
		"fish",
		"nushell",
		"oil",
		"powershell",
		"tcsh",
		"xonsh",
		"zsh":
		groupPriority[b] = -35
	}

	// TODO urks - this is getting a bit out of hand. use copy map? needs to still support `force_all` build tag
	switch goos := runtime.GOOS; goos {
	case "android":
		maps.Copy(groupPriority, map[string]int{
			goos:    -15,
			"linux": -14,
			"unix":  -13,
		})
	case "darwin":
		maps.Copy(groupPriority, map[string]int{
			goos:   -15,
			"bsd":  -14,
			"unix": -13,
		})
	case "freebsd", "netbsd", "openbsd":
		maps.Copy(groupPriority, map[string]int{
			goos:   -15,
			"bsd":  -14,
			"unix": -13,
		})
	default:
		groupPriority[goos] = -15 // ensure runime.GOOS has highest priority betwees gooses
	}

	if diff := groupPriority[c[i].Group] - groupPriority[c[j].Group]; diff != 0 {
		return diff < 0 // TODO verify
	}

	if c[i].Group == "bridge" {
		// TODO this should be wrong when the implicit bridge is not configured (default value could have higher priority - or can it?)
		bridges := strings.Split(os.Getenv("CARAPACE_BRIDGES"), ",")
		return slices.Index(bridges, c[i].Variant) < slices.Index(bridges, c[j].Variant)
	}

	return strings.Compare(c[i].Variant, c[j].Variant) < 0
}
func (c Completers) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
