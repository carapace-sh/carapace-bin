package completer

import (
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
	if b := strings.Compare(c[i].Name, c[j].Name) != 0; b {
		return b
	}

	// choices have highest priority
	if a, b := c[i].Choice, c[j].Choice; a != b {
		return a > b
	}

	groupPriority := map[string]int{
		// user specs
		"user": -7,
		// system specs
		"system": -6,
		// TODO shells? more specific stuff
		// runtime.GOOS: -5 (see below)
		"linux":   -4,
		"darwin":  -3,
		"windows": -2,
		// TODO support pseudo os 'termux'?
		"bridge": 1, // lower priority than anything internal
	}
	groupPriority[runtime.GOOS] = -5 // ensure runime.GOOS has highest priority betwees gooses

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
