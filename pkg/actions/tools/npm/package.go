package npm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPackageSearch(registry string) carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackageNames(registry)
		case 1:
			return ActionPackageVersions(registry, c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionPackageNames(registry string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"search", "--parseable", "--searchlimit", "250", fmt.Sprintf(`/^%v`, c.CallbackValue)}
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

func ActionPackageVersions(registry, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", pkg, "versions", "--json"}
		if registry != "" {
			args = append(args, "--registry", registry)
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

func ActionPackageTags(registry, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", pkg, "dist-tags", "--json"}
		if registry != "" {
			args = append(args, "--registry", registry)
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
