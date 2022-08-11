// package pip contains pip related actions
package pip

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

type pkg struct {
	Name    string
	Version string
}

// ActionInstalledPackages completes installed packages
//   Automat (20.2.0)
//   Beaker (1.11.0)
func ActionInstalledPackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pip", "list", "--format", "json")(func(output []byte) carapace.Action {
			var packages []pkg
			json.Unmarshal(output, &packages)

			vals := make([]string, len(packages)*2)
			for index, p := range packages {
				vals[index*2] = p.Name
				vals[index*2+1] = p.Version
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionPackageSearch completes remote packages
//   git-gopher (Improving the Git CLI experience with fzf)
//   git-lint (Git Lint)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO pip search currently disabled due to load - basic workaround without paging
		return carapace.ActionExecCommand("curl", fmt.Sprintf("https://pypi.org/search/?q=%v", c.CallbackValue))(func(output []byte) carapace.Action {
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
		})
	})
}

type ConfigOpts struct {
	Global bool
	Site   bool
	User   bool
}

// ActionConfigValues completes config values
//   first ('1'
//   second ('2')
func ActionConfigValues(opts ConfigOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO fragile, also should be pflag but didn't work
			args := []string{"config"}
			if opts.Global {
				args = append(args, "--global")
			} else if opts.Site {
				args = append(args, "--site")
			} else if opts.User {
				args = append(args, "--user")
			}
			args = append(args, "list")

			return carapace.ActionExecCommand("pip", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")

				vals := make([]string, (len(lines)-1)*2)
				for index, line := range lines[:len(lines)-1] {
					splitted := strings.SplitN(line, "=", 2)
					vals[index*2] = splitted[0]
					vals[index*2+1] = splitted[1]
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		}).Invoke(c).ToMultiPartsA(".")
	})
}
