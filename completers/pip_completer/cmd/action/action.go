package action

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type pkg struct {
	Name    string
	Version string
}

func ActionInstalledPackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("pip", "list", "--format", "json").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			var packages []pkg
			json.Unmarshal(output, &packages)

			vals := make([]string, len(packages)*2)
			for index, p := range packages {
				vals[index*2] = p.Name
				vals[index*2+1] = p.Version
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionRemotePackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO pip search currently disabled due to load - basic workaround without paging
		if output, err := exec.Command("curl", fmt.Sprintf("https://pypi.org/search/?q=%v", c.CallbackValue)).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			re := regexp.MustCompile(`class="package-snippet__(name|description)">(?P<val>.*)<`)

			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					vals = append(vals, re.FindStringSubmatch(line)[2])
				}
			}

			if len(vals)%2 != 0 {
				return carapace.ActionMessage("failed to parse response")
			} else {
				return carapace.ActionValuesDescribed(vals...)
			}
		}
	})
}

func ActionConfigValues(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO fragile, also should be pflag but didn't work
			args := []string{"config"}
			if cmd.Parent().Flag("global").Changed {
				args = append(args, "--global")
			} else if cmd.Parent().Flag("site").Changed {
				args = append(args, "--site")
			} else if cmd.Parent().Flag("user").Changed {
				args = append(args, "--user")
			}
			args = append(args, "list")

			if output, err := exec.Command("pip", args...).Output(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				lines := strings.Split(string(output), "\n")

				vals := make([]string, (len(lines)-1)*2)
				for index, line := range lines[:len(lines)-1] {
					splitted := strings.SplitN(line, "=", 2)
					vals[index*2] = splitted[0]
					vals[index*2+1] = splitted[1]
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Invoke(c).ToMultiPartsA(".")
	})
}
