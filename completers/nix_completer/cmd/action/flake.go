package action

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
)

// ActionFlakeRefs completes a full flake reference
// It will only complete attributes for local flakes or flakes
// in the flake registry.
// It takes in the subcommand being completed
//
// nixpkgs#hello
// .#foo
func ActionFlakeRefs(fullCmd []string) carapace.Action {
	return carapace.ActionMultiPartsN("#", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return nix.ActionFlakes().Suffix("#")
		default:
			return ActionFlakeAttributes(fullCmd, c.Parts[0], c.Value)
		}
	})
}

// ActionFlakeAttributes completes attributes on a flake
// Completions are only supplied for local flakes or flakes
// in the registry.
//
// hello
// packages.x86_64-linux.hello
func ActionFlakeAttributes(fullCmd []string, flake, flakePath string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// Get completions from the flake registry
		c.Setenv("NIX_GET_COMPLETIONS", fmt.Sprint(len(fullCmd)))
		completionArgs := append(fullCmd[1:], flake)

		var stdout, stderr bytes.Buffer
		cmd := c.Command("nix", completionArgs...)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return carapace.ActionValues()
		}

		// Check if the given flake is in the registry or is a local path
		// For some reason, nix adds \b to completion items
		rawLines := strings.ReplaceAll(stdout.String(), "\r\b", "")
		lines := strings.Split(rawLines, "\n")

		flakeInRegistry := slices.ContainsFunc(lines[1:], func(s string) bool {
			return strings.Trim(s, " \t") == flake
		})
		flakeIsLocal := directoryExists(flake)

		if !flakeInRegistry && !flakeIsLocal {
			return carapace.ActionValues()
		}

		// Get completions using the nix cli
		stdout.Reset()
		stderr.Reset()

		completionArgs = append(fullCmd[1:], flake+"#"+flakePath)
		cmd = c.Command("nix", completionArgs...)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return carapace.ActionValues()
		}

		// Remove \r for consistency with line endings
		// For some reason, nix adds \b to completion items
		rawLines = strings.ReplaceAll(stdout.String(), "\r\b", "")
		lines = strings.Split(rawLines, "\n")

		// The first line is always "attrs"
		completionResults := lines[1:]

		// Remove the flake# prefix
		for i := range completionResults {
			completionResults[i] = strings.Trim(completionResults[i], " \t")
			completionResults[i] = strings.TrimPrefix(completionResults[i], flake+"#")
		}

		return carapace.ActionValues(completionResults...)
	})
}

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}
