package completers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	"github.com/carapace-sh/carapace-bridge/pkg/choice"
	bridge_env "github.com/carapace-sh/carapace-bridge/pkg/env"
	"github.com/carapace-sh/carapace/pkg/xdg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// TODO rename parse - init? if false returned completers cannot be executed
func Completers(filter choice.Choice, parse bool) (completer.CompleterMap, error) {
	// TODO apply filter!!
	// TODO slow when just used for lookup
	m := completers.Filter(filter)

	switch filter.Group {
	case "", "user", "system":
		if err := AddSpecs(m, parse); err != nil {
			return nil, err
		}
	}

	switch filter.Group {
	case "", "bridge":
		if err := AddBridges(m, parse); err != nil {
			return nil, err
		}
	}

	if err := AddChoices(m, filter, parse); err != nil {
		return nil, err
	}

	if err := AddOverlays(m); err != nil {
		return nil, err
	}
	return m.Filter(filter), nil

}

func AddChoices(m completer.CompleterMap, filter choice.Choice, parse bool) error {
	if !parse {
		// add potentially unknown bridges (eg: `gh/cobra@bridge`)
		if filter.Group != "bridge" {
			return nil // anything other than custom bridges should already exist
		}

		choices, err := choice.List(false)
		if err != nil {
			return err
		}
		for _, choice := range choices {
			if filter.Name == "" || choice.Name == filter.Name {
				if _, ok := m[choice.Name]; !ok {
					// TODO
					m[choice.Name] = completer.Completers{
						{
							Name: choice.Name,
						},
					}
				}
			}
		}
		return nil
	}

	var choices []*choice.Choice
	switch {
	case filter.Name != "":
		c, err := choice.Get(filter.Name)
		switch {
		case os.IsNotExist(err):
			return nil
		case err != nil:
			return err
		default:
			choices = append(choices, c)
		}
	default:
		c, err := choice.List(parse)
		if err != nil {
			return err
		}
		choices = append(choices, c...)
	}

	// TODO sanity checks for name/variant
	for _, c := range choices {
		switch c.Group {
		case "bridge":
			matched := false
			for index, v := range m[c.Name] {
				if (c.Variant == "" || v.Variant == c.Variant) && v.Group == c.Group {
					m[c.Name][index].Choice = c.Format()
					matched = true
				}
			}

			if matched || c.Variant == "" {
				continue
			}

			// TODO should be prepended
			m[c.Name] = append(m[c.Name], completer.Completer{
				Name:    c.Name,
				Variant: c.Variant,
				Group:   c.Group,
				Choice:  c.Format(),
				Execute: func() error {
					if f, ok := bridge.Get(c.Variant); ok {
						return complete(c.Name, f(c.Name))()
					}
					return nil
				},
			})

		default:
			// mark matching variant
			for index, v := range m[c.Name] {
				if v.Variant == c.Variant && v.Group == c.Group { // TODO allow empty variant/group?
					m[c.Name][index].Choice = c.Format()
				}
			}
		}
	}
	return nil
}

func AddOverlays(m completer.CompleterMap) error {
	// TODO support overlays specific to variants
	dir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(filepath.Join(dir, "carapace", "overlays"))
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".yaml")
		path := filepath.Join(dir, entry.Name())
		if variants, ok := m[name]; ok {
			for index := range variants {
				variants[index].Overlay = path
			}
		}
	}
	return nil
}

func AddBridges(m completer.CompleterMap, parse bool) error {
	description := func(name string) string {
		if !parse {
			return ""
		}
		return Description(name)
	}

	// TODO description is unclear for bridges and should probably be omitted
	// TODO remove - replace by choices (do the same in carapace-bin)
	// for name, b := range bridges.Config() {
	// 	// TODO differentiate explicit bridge from the rest (currently overlaps are possible)
	// 	m[name] = append(m[name], completer.Completer{
	// 		Name:        name,
	// 		Description: description(name), // TODO
	// 		Group:       "bridge.config",
	// 		Variant:     b,
	// 		Execute: func() error {
	// 			// TODO execute the corresponding bridge
	// 			if f, ok := bridge.Get(b); ok {
	// 				complete(name, f(name))()
	// 			}
	// 			return nil
	// 		},
	// 	})
	// }

	for _, b := range bridge_env.Bridges() {
		switch b {
		case "bash":
			for _, name := range bridges.Bash() {
				m[name] = append(m[name], completer.Completer{
					Name:        name,
					Description: description(name), // TODO
					Group:       "bridge",
					Variant:     "bash",
					Execute:     complete(name, bridge.ActionBash(name)),
				})
			}
		case "fish":
			for _, name := range bridges.Fish() {
				m[name] = append(m[name], completer.Completer{
					Name:        name,
					Description: description(name), // TODO
					Group:       "bridge",
					Variant:     "fish",
					Execute:     complete(name, bridge.ActionFish(name)),
				})
			}
		case "inshellisense":
			for _, name := range bridges.Inshellisense() {
				m[name] = append(m[name], completer.Completer{
					Name:        name,
					Description: description(name), // TODO
					Group:       "bridge",
					Variant:     "inshellisense",
					Execute:     complete(name, bridge.ActionInshellisense(name)),
				})
			}
		case "zsh":
			for _, name := range bridges.Zsh() {
				m[name] = append(m[name], completer.Completer{
					Name:        name,
					Description: description(name), // TODO
					Group:       "bridge",
					Variant:     "zsh",
					Execute:     complete(name, bridge.ActionZsh(name)),
				})
			}
		}
	}
	return nil
}

func AddSpecs(m completer.CompleterMap, parse bool) error {
	dir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	specs, err := completer.ReadSpecs(filepath.Join(dir, "carapace", "specs"), "user", parse)
	if err != nil {
		return err
	}
	m.Merge(specs)
	return nil
}

// TODO choice as parameter?
func Lookup(nameVariantGroup string) (*completer.Completer, error) {
	m, err := Completers(choice.Parse(nameVariantGroup), true) // TODO lookup needs to use a quick version (skip parsing of specs for descriptions)
	if err != nil {
		return nil, err
	}
	if c, ok := m.Lookup(nameVariantGroup); ok {
		return c, nil
	}
	return nil, fmt.Errorf("unknonw completer/variant: %#v", nameVariantGroup)
}

// TODO adopt conditional names
// func Names() []string {
// 	excludes := make(map[string]bool)
// 	for _, e := range env.Excludes() {
// 		excludes[e] = true
// 	}

// 	unique := map[string]bool{
// 		"carapace": true,
// 	}

// 	if os.Getenv("CARAPACE_ENV") != "0" {
// 		unique["get-env"] = true
// 		unique["set-env"] = true
// 		unique["unset-env"] = true
// 	}

// 	if _, ok := excludes["*"]; !ok {
// 		for name := range completers {
// 			if _, ok := excludes[name]; !ok {
// 				unique[name] = true
// 			}
// 		}
// 	}
// 	//if specNames, err := Specs(); err == nil {
// 	if specNames, _ := Specs(); true { // TODO use and handle err
// 		for _, name := range specNames {
// 			unique[name] = true
// 		}
// 	}

// 	combined := make([]string, 0, len(unique))
// 	for name := range unique {
// 		combined = append(combined, name)
// 	}
// 	sort.Strings(combined)
// 	return combined
// }

func Description(name string) string {
	if d, err := specDescription(name); err == nil {
		return d
	}
	if completer, ok := completers.Lookup(name); ok {
		return completer.Description
	}
	return ""
}

func specDescription(name string) (string, error) {
	confDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(fmt.Sprintf("%v/carapace/specs/%v.yaml", confDir, name))
	if err != nil {
		return "", err
	}
	var s struct {
		Description string
	}

	err = yaml.Unmarshal(content, &s)
	if err != nil {
		return "", err
	}
	return s.Description, nil
}

func complete(name string, action carapace.Action) func() error {
	return func() error {
		cmd := &cobra.Command{
			Use:                name,
			DisableFlagParsing: true,
		}
		carapace.Gen(cmd).Standalone()
		carapace.Gen(cmd).PositionalAnyCompletion(
			action,
		)
		return cmd.Execute()
	}
}
