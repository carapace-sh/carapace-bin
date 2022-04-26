package action

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackageNames(cmd)
		case 1:
			return ActionPackageVersions(cmd, c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionPackageNames(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"search", "--parseable", "--searchlimit", "250", fmt.Sprintf(`/^%v`, c.CallbackValue)}
		if flag := cmd.Flag("registry"); flag.Changed {
			args = append(args, "--registry", flag.Value.String())
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

func ActionPackageVersions(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", pkg, "versions", "--json"}
		if flag := cmd.Flag("registry"); flag.Changed {
			args = append(args, "--registry", flag.Value.String())
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

func ActionPackageTags(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", pkg, "dist-tags", "--json"}
		if flag := cmd.Flag("registry"); flag.Changed {
			args = append(args, "--registry", flag.Value.String())
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

type packageJson struct {
	Scripts    map[string]string
	Workspaces []string
}

func loadPackageJson() (pj packageJson, err error) {
	var wd, packageFile string
	if wd, err = os.Getwd(); err == nil {
		if packageFile, err = util.FindReverse(wd, "package.json"); err == nil {
			var content []byte
			if content, err = os.ReadFile(packageFile); err == nil {
				err = json.Unmarshal(content, &pj)
			}
		}
	}
	return
}

func ActionScripts(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if pj, err := loadPackageJson(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for name := range pj.Scripts {
				vals = append(vals, name)
			}
			return carapace.ActionValues(vals...)
		}
	})
}
