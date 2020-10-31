package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

type manifestJson struct {
	Dependencies []struct {
		Name string
		Req  string
	}
	Features map[string][]string

	Targets []struct {
		Name string
		Kind []string
	}
}

type manifestToml struct {
	Profile   map[string]struct{}
	Workspace struct {
		Members []string
	}
}

func manifestLocation(cmd *cobra.Command) (string, error) {
	if f := cmd.Flag("manifest-path"); f != nil && f.Changed {
		return f.Value.String(), nil
	}
	return util.FindReverse("", "Cargo.toml")
}

func readManifest(cmd *cobra.Command) (m manifestJson, err error) {
	var output []byte
	var path string
	if path, err = manifestLocation(cmd); err == nil {
		if output, err = exec.Command("cargo", "read-manifest", "--offline", "--manifest-path", path).Output(); err == nil {
			err = json.Unmarshal(output, &m)
		}
	}
	return
}

func parseManifest(cmd *cobra.Command) (m manifestToml, err error) {
	var content []byte
	var path string
	if path, err = manifestLocation(cmd); err == nil {
		if content, err = ioutil.ReadFile(path); err == nil {
			err = toml.Unmarshal(content, &m)
		}
	}
	return
}

func ActionColorModes() carapace.Action {
	return carapace.ActionValues("auto", "always", "never")
}

type TargetOpts struct {
	Bench   bool
	Bin     bool
	Example bool
	Lib     bool
	Test    bool
}

func (t *TargetOpts) Includes(kinds []string) bool {
	var match bool
	for _, kind := range kinds {
		switch kind {
		case "bench":
			match = t.Bench
		case "bin":
			match = t.Bin
		case "example":
			match = t.Example
		case "lib":
			match = t.Lib
		case "test":
			match = t.Test
		}
		if match {
			return true
		}
	}
	return false
}

func readManifestAction(cmd *cobra.Command, f func(m manifestJson, args []string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if m, err := readManifest(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return f(m, args)
		}
	})
}

func parseManifestAction(cmd *cobra.Command, f func(m manifestToml, args []string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if m, err := parseManifest(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return f(m, args)
		}
	})
}

func ActionTargets(cmd *cobra.Command, opts TargetOpts) carapace.Action {
	return readManifestAction(cmd, func(m manifestJson, args []string) carapace.Action {
		vals := make([]string, 0)
		for _, target := range m.Targets {
			if opts.Includes(target.Kind) {
				vals = append(vals, target.Name, strings.Join(target.Kind, ", "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionDependencies(cmd *cobra.Command) carapace.Action {
	return readManifestAction(cmd, func(m manifestJson, args []string) carapace.Action {
		vals := make([]string, len(m.Dependencies)*2)
		for index, dependency := range m.Dependencies {
			vals[index*2] = fmt.Sprintf("%v:%v", dependency.Name, strings.TrimLeft(dependency.Req, "^"))
			vals[(index*2)+1] = dependency.Req
		}
		return carapace.ActionValuesDescribed(vals...)

	})
}

func ActionWorkspaceMembers(cmd *cobra.Command) carapace.Action {
	return parseManifestAction(cmd, func(m manifestToml, args []string) carapace.Action {
		return carapace.ActionValues(m.Workspace.Members...)
	})
}

func ActionMessageFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"human", "Display in a human-readable text format",
		"short", "Emit shorter, human-readable text messages",
		"json", "Emit JSON messages to stdout",
		"json-diagnostic-short", "Ensure the rendered field of JSON messages contains the 'short' rendering from rustc",
		"json-diagnostic-rendered-ansi", "Ensure the rendered field of JSON messages contains embedded ANSI color codes",
		"json-render-diagnostics", "Instruct Cargo to not include rustc diagnostics in in JSON messages printed",
	)
}

func ActionFeatures(cmd *cobra.Command) carapace.Action {
	return readManifestAction(cmd, func(m manifestJson, args []string) carapace.Action {
		vals := make([]string, 0)
		for name, packages := range m.Features {
			vals = append(vals, name, strings.Join(packages, ", "))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionProfiles(cmd *cobra.Command) carapace.Action {
	return parseManifestAction(cmd, func(m manifestToml, args []string) carapace.Action {
		profiles := make([]string, 0)
		for key := range m.Profile {
			profiles = append(profiles, key)
		}
		return carapace.ActionValues(profiles...)
	})
}
