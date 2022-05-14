// package pacman contains arch package manager related actions
package pacman

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"gopkg.in/ini.v1"
)

// ActionRepositories completes package repositories
//   extra
//   multilib
func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		loadOptions := ini.LoadOptions{}
		loadOptions.AllowBooleanKeys = true
		loadOptions.UnparseableSections = []string{"options"}
		if cfg, err := ini.LoadSources(loadOptions, "/etc/pacman.conf"); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			repos := make([]string, 0)
			for _, section := range cfg.SectionStrings() {
				if section != "DEFAULT" && section != "options" {
					repos = append(repos, section)
				}
			}
			return carapace.ActionValues(repos...)
		}
	})
}

type PackageOption struct {
	Explicit bool
}

func (o PackageOption) format(callbackValue string) []string {
	options := []string{"-Qq"}
	if o.Explicit {
		options[0] = options[0] + "e"
	} else {
		options = []string{options[0] + "s", fmt.Sprintf("^%v", callbackValue)}
	}
	return options
}

// ActionPackages completes package names
//   gnome-common
//   gstreamer
func ActionPackages(option PackageOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pacman", option.format(c.CallbackValue)...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

// ActionPackageGroups completes package group names
//   i3-manjaro
//   linux49-extramodules
func ActionPackageGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("sh", "-c", "pacman -Qg | cut -d' ' -f1 | sort |  uniq")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
