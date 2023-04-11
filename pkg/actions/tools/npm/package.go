package npm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionPackageSearch completes packages@version for given registry
func ActionPackageSearch(registry string) carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackageNames(registry).NoSpace()
		case 1:
			return ActionPackageVersions(PackageOpts{Registry: registry, Package: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionPackageNames completes package names for given registry
func ActionPackageNames(registry string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"search", "--parseable", "--searchlimit", "250", fmt.Sprintf(`/^%v`, c.Value)}
		if registry != "" {
			args = append(args, "--registry", registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				fields := strings.Split(line, "\t")
				vals = append(vals, fields[0], fields[1])
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type PackageOpts struct {
	Registry string
	Package  string
}

// ActionPackageVersions completes versions for given package
func ActionPackageVersions(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", opts.Package, "versions", "--json"}
		if opts.Registry != "" {
			args = append(args, "--registry", opts.Registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var versions []string
			if err := json.Unmarshal(output, &versions); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValues(versions...)
		})
	})
}

// ActionPackageTags completes tags for given package
func ActionPackageTags(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", opts.Package, "dist-tags", "--json"}
		if opts.Registry != "" {
			args = append(args, "--registry", opts.Registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var tags map[string]string
			if err := json.Unmarshal(output, &tags); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(tags)*2)
			for tag, version := range tags {
				vals = append(vals, tag, version)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
