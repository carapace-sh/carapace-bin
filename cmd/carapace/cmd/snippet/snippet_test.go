package snippet

import (
	"testing"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
)

func TestFilterNativeCompletions(t *testing.T) {
	tests := map[string]struct {
		nativeNames []string
		completers  completer.CompleterMap
		chosenNames map[string]bool
		expected    map[string]bool
	}{
		"bridge only": {
			nativeNames: []string{"native"},
			completers: completer.CompleterMap{
				"native": {
					{Group: "bridge", Variant: "bash"},
					{Group: "bridge", Variant: "zsh"},
				},
			},
			expected: map[string]bool{"other": true},
		},
		"bridge and non-bridge": {
			nativeNames: []string{"common", "other-group", "platform", "system", "user"},
			completers: completer.CompleterMap{
				"common":      {{Group: "bridge"}, {Group: "common"}},
				"other-group": {{Group: "bridge"}, {Group: "custom"}},
				"platform":    {{Group: "bridge"}, {Group: "darwin"}},
				"system":      {{Group: "bridge"}, {Group: "system"}},
				"user":        {{Group: "bridge"}, {Group: "user"}},
			},
			expected: map[string]bool{
				"common":      true,
				"other":       true,
				"other-group": true,
				"platform":    true,
				"system":      true,
				"user":        true,
			},
		},
		"chosen bridge": {
			nativeNames: []string{"chosen", "unchosen"},
			completers: completer.CompleterMap{
				"chosen":   {{Group: "bridge"}},
				"unchosen": {{Group: "bridge"}},
			},
			chosenNames: map[string]bool{"chosen": true},
			expected:    map[string]bool{"chosen": true, "other": true},
		},
		"non-native bridge": {
			nativeNames: []string{"native"},
			completers: completer.CompleterMap{
				"native":     {{Group: "bridge"}},
				"non-native": {{Group: "bridge"}},
			},
			expected: map[string]bool{"non-native": true, "other": true},
		},
		"missing variants": {
			nativeNames: []string{"empty", "missing"},
			completers: completer.CompleterMap{
				"empty": {},
			},
			expected: map[string]bool{"other": true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			uniqueNames := map[string]bool{"other": true}
			for completerName := range test.completers {
				uniqueNames[completerName] = true
			}

			filterNativeCompletions(uniqueNames, test.nativeNames, test.completers, test.chosenNames)

			if !equalNames(uniqueNames, test.expected) {
				t.Fatalf("expected %#v, got %#v", test.expected, uniqueNames)
			}
		})
	}
}

func TestFilterNativeCompletionsShells(t *testing.T) {
	for _, shell := range []string{"bash", "fish", "zsh"} {
		t.Run(shell, func(t *testing.T) {
			uniqueNames := map[string]bool{
				"native":     true,
				"non-native": true,
			}
			m := completer.CompleterMap{
				"native":     {{Group: "bridge"}},
				"non-native": {{Group: "bridge"}},
			}

			filterNativeCompletions(uniqueNames, []string{"native"}, m, nil)

			expected := map[string]bool{"non-native": true}
			if !equalNames(uniqueNames, expected) {
				t.Fatalf("expected %#v, got %#v", expected, uniqueNames)
			}
		})
	}
}

func TestRemoveExcludesAfterFilterNativeCompletions(t *testing.T) {
	tests := map[string]struct {
		completers  completer.Completers
		chosenNames map[string]bool
	}{
		"chosen bridge": {
			completers:  completer.Completers{{Group: "bridge"}},
			chosenNames: map[string]bool{"native": true},
		},
		"mixed variants": {
			completers: completer.Completers{{Group: "bridge"}, {Group: "common"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Setenv("CARAPACE_EXCLUDES", "native")

			uniqueNames := map[string]bool{"native": true}
			m := completer.CompleterMap{"native": test.completers}

			filterNativeCompletions(uniqueNames, []string{"native"}, m, test.chosenNames)
			completers.RemoveExcludes(uniqueNames)

			if len(uniqueNames) != 0 {
				t.Fatalf("expected excluded native completer to be removed, got %#v", uniqueNames)
			}
		})
	}
}

func equalNames(actual, expected map[string]bool) bool {
	if len(actual) != len(expected) {
		return false
	}
	for name, value := range expected {
		if actual[name] != value {
			return false
		}
	}
	return true
}
