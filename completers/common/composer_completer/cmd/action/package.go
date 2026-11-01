package action

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/spf13/cobra"
)

type installedPackage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// installedPackages returns the packages composer considers installed:
// the ones in the lock file or, without one, the ones in vendor/composer/installed.json.
func installedPackages(c carapace.Context, cmd *cobra.Command) ([]installedPackage, error) {
	path, err := composerFile(c, cmd)
	if err != nil {
		return nil, err
	}

	if content, err := os.ReadFile(lockFile(path)); err == nil {
		var lock struct {
			Packages    []installedPackage `json:"packages"`
			PackagesDev []installedPackage `json:"packages-dev"`
		}
		if err := json.Unmarshal(content, &lock); err != nil {
			return nil, fmt.Errorf("%s: %w", filepath.Base(lockFile(path)), err)
		}
		return append(lock.Packages, lock.PackagesDev...), nil
	}

	p, err := readProject(path)
	if err != nil {
		p = project{} // vendor directory might still exist
	}

	vendor, err := vendorDir(c, cmd, p)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(filepath.Join(vendor, "composer", "installed.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no lock file or installed packages found")
		}
		return nil, err
	}

	var installed struct {
		Packages []installedPackage `json:"packages"`
	}
	if err := json.Unmarshal(content, &installed); err != nil {
		var packages []installedPackage // composer 1 format
		if err := json.Unmarshal(content, &packages); err != nil {
			return nil, fmt.Errorf("installed.json: %w", err)
		}
		installed.Packages = packages
	}
	return installed.Packages, nil
}

// ActionInstalledPackages completes the installed packages (as listed in the lock file or vendor/composer/installed.json)
func ActionInstalledPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		packages, err := installedPackages(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(packages)*2)
		for _, p := range packages {
			vals = append(vals, p.Name, oneLine(p.Description))
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("installed packages")
}

// ActionInstalledPackageTypes completes the types of the installed packages
func ActionInstalledPackageTypes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		packages, err := installedPackages(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		types := make(map[string]bool)
		for _, p := range packages {
			switch p.Type {
			case "":
				types["library"] = true
			default:
				types[p.Type] = true
			}
		}

		vals := make([]string, 0, len(types))
		for t := range types {
			vals = append(vals, t)
		}
		sort.Strings(vals)
		return carapace.ActionValues(vals...)
	}).Tag("package types")
}

// ActionRootRequirements completes the packages required by the project (require and require-dev)
func ActionRootRequirements(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadProject(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, section := range []string{"require", "require-dev"} {
			for name, constraint := range p.object(section) {
				vals = append(vals, name, fmt.Sprint(constraint))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("required packages")
}

// ActionPlatformPackages completes the platform packages (php, extensions and libraries)
func ActionPlatformPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"show", "--platform", "--format=json", "--no-ansi", "--no-interaction", "--no-plugins", "--no-scripts"}
		dir := ""
		if flag := cmd.Flag("working-dir"); flag != nil && flag.Value.String() != "" {
			dir = flag.Value.String()
			args = append(args, "--working-dir="+dir)
		}

		return carapace.ActionExecCommand("composer", args...)(func(output []byte) carapace.Action {
			var platform struct {
				Platform []struct {
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"platform"`
			}
			if err := json.Unmarshal(output, &platform); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(platform.Platform)*2)
			for _, p := range platform.Platform {
				vals = append(vals, p.Name, oneLine(p.Description))
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Cache(time.Hour, key.String(c.Dir, dir))
	}).Tag("platform packages")
}

// ActionPackagesByMode completes the packages `composer show` is able to inspect (depending on the given flags)
func ActionPackagesByMode(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		for _, name := range []string{"all", "available"} {
			if flag := cmd.Flag(name); flag != nil && flag.Value.String() == "true" {
				return carapace.ActionValues() // would require querying the repositories
			}
		}

		if flag := cmd.Flag("platform"); flag != nil && flag.Value.String() == "true" {
			return ActionPlatformPackages(cmd)
		}
		return ActionInstalledPackages(cmd)
	})
}
