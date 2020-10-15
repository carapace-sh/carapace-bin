package pacman

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"gopkg.in/ini.v1"
)

func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
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

func (o PackageOption) Format() []string {
	options := []string{"-Qq"}
	if o.Explicit {
		options[0] = options[0] + "e"
	} else {
		options = []string{options[0] + "s", fmt.Sprintf("^%v", carapace.CallbackValue)}
	}
	return options
}

func ActionPackages(option PackageOption) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("pacman", option.Format()...).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}

func ActionPackageGroups() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("sh", "-c", "pacman -Qg | cut -d' ' -f1 | sort |  uniq").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}
