package git

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"gopkg.in/ini.v1"
)

type submodule struct {
	Path   string `ini:"path"`
	Url    string `ini:"url"`
	Branch string `ini:"branch"`
}

func loadSubmodules(c carapace.Context) (submodules map[string]*submodule, err error) {
	var root string
	if root, err = rootDir(c); err == nil {
		var cfg *ini.File
		if cfg, err = ini.Load(root + "/.gitmodules"); err == nil {
			r := regexp.MustCompile(`^submodule "(?P<name>.*)"$`)
			submodules = make(map[string]*submodule)
			for _, section := range cfg.Sections() {
				if r.MatchString(section.Name()) {
					matches := r.FindStringSubmatch(section.Name())
					s := new(submodule)
					section.MapTo(s)
					submodules[matches[1]] = s
				}
			}
		}
	}
	return
}

// ActionSubmoduleNames completes submodule names
//
//	cobra (https://github.com/spf13/cobra.git)
//	pflag (https://github.com/spf13/pflag.git)
func ActionSubmoduleNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		submodules, err := loadSubmodules(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for name, s := range submodules {
			vals = append(vals, name, s.Url)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionSubmoduleBranches completes brances and tags of submodules (filtered by name)
// TODO verify and add example
func ActionSubmoduleBranches(filter ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		submodules, err := loadSubmodules(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		actions := make([]carapace.Action, 0)
		for name, s := range submodules {
			if len(filter) == 0 || contains(filter, name) {
				if s.Url != "" {
					actions = append(actions, ActionLsRemoteRefs(LsRemoteRefOption{Url: s.Url, Branches: true, Tags: true}))
				}
			}
		}
		return carapace.Batch(actions...).Invoke(c).Merge().ToA()
	})
}

func contains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}

// ActionSubmodulePaths completes submodules (relative path)
//
//	../cli (v1.12.1-12-gaed8966f)
//	../pflag (heads/forOgier)
func ActionSubmodulePaths() carapace.Action {
	// TODO use loadSubmodules
	return carapace.ActionExecCommand("git", "submodule")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^.(?P<commit>[^ ]+) (?P<path>.*) \((?P<ref>[^()]+)\)$`)

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[2], matches[3])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
