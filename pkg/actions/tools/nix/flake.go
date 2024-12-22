package nix

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFlakes completes flakes currently available
//
// nixpkgs
// .
func ActionFlakes() carapace.Action {
	return carapace.ActionExecCommand("nix", "registry", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		// `nix registry list` outputs a blank line after the listings
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			name := strings.Split(fields[1], ":")
			if name[0] == "flake" {
				vals = append(vals, name[1], fields[2], styleForRegistry(fields[0]))
			}
		}
		// TODO add directory completion externally
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Cache(time.Minute).Tag("flakes")
}

func styleForRegistry(s string) string {
	switch s {
	case "global":
		return style.Blue
	case "system":
		return style.Yellow
	default:
		return style.Default
	}
}

// ActionFlakeAttributes completes attributes on a flake
// Completions are only supplied for local flakes or flakes
// in the registry.
//
//	hello
//	packages.x86_64-linux.hello
func ActionFlakeAttributes(flake string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// Get completions from the flake registry
		c.Setenv("NIX_GET_COMPLETIONS", "2")

		var stdout, stderr bytes.Buffer
		cmd := c.Command("nix", "build", flake)
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

		cmd = c.Command("nix", "build", fmt.Sprintf("%v#%v", flake, c.Value))
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

// ActionFlakeRefs completes a full flake reference
// It will only complete attributes for local flakes or flakes
// in the flake registry.
// It takes in the subcommand being completed
//
//	nixpkgs#hello
//	.#foo
func ActionFlakeRefs() carapace.Action {
	return carapace.ActionMultiPartsN("#", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			// TODO add directory completion externally
			return ActionFlakes().Suffix("#")
		default:
			return ActionFlakeAttributes(c.Parts[0])
		}
	})
}
