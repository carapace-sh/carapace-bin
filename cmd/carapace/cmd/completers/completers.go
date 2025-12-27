package completers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bin/pkg/env"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	bridge_env "github.com/carapace-sh/carapace-bridge/pkg/env"
	"github.com/carapace-sh/carapace/pkg/xdg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Completers(filter choices.Choice, parse bool) (completer.CompleterMap, error) { // TODO rename parse - init? if false returned completers cannot be executed
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

	RemoveExcludes(m)

	// TODO add pseudo carapace completer here instead of cmd
	AddCarapace(m)

	return m.Filter(filter), nil

}

func AddCarapace(m completer.CompleterMap) {
	m["carapace"] = append(m["carapace"], completer.Completer{
		Name:        "carapace",
		Description: "A multi-shell completion binary",
		Group:       "common",
		Package:     "github.com/carapace-sh/carapace-bin/cmd/carapace/cmd",
		Url:         "https://carapace.sh",
		Execute:     func() error { return nil }, // TODO verify - there shouldn't be any need for this
	})
}

func RemoveExcludes(m completer.CompleterMap) {
	for _, e := range env.Excludes() {
		if e == "*" {
			// TODO wildcard doesn't make much sense anymore as this does not only affect internal completers
			// TODO there's choices now which supersedes this
			// TODO excludes pretty much should now just prevent a command to be registered by carapace in general
			clear(m)
			return
		}
		delete(m, e)
	}

	if os.Getenv("CARAPACE_ENV") == "0" { // TODO util function
		delete(m, "get-env")
		delete(m, "set-env")
		delete(m, "unset-env")
	}
}

func AddChoices(m completer.CompleterMap, filter choices.Choice, parse bool) error {
	if !parse {
		// add potentially unknown bridges (eg: `gh/cobra@bridge`)
		if filter.Group != "bridge" {
			return nil // anything other than custom bridges should already exist
		}

		list, err := choices.List(false)
		if err != nil {
			return err
		}
		for _, choice := range list {
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

	var list []*choices.Choice
	switch {
	case filter.Name != "":
		c, err := choices.Get(filter.Name)
		switch {
		case os.IsNotExist(err):
			return nil
		case err != nil:
			return err
		default:
			list = append(list, c)
		}
	default:
		c, err := choices.List(parse)
		if err != nil {
			return err
		}
		list = append(list, c...)
	}

	// TODO sanity checks for name/variant
	for _, c := range list {
		switch c.Group {
		case "bridge":
			matched := false
			for index, v := range m[c.Name] {
				if (c.Variant == "" || v.Variant == c.Variant) && (c.Group == "" || v.Group == c.Group) {
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
				if (c.Variant == "" || v.Variant == c.Variant) && (c.Group == "" || v.Group == c.Group) { // TODO allow empty variant/group?
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
		if os.IsNotExist(err) {
			return nil
		}
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

func Lookup(nameVariantGroup string) (*completer.Completer, error) { // TODO choice as parameter?
	m, err := Completers(choices.Parse(nameVariantGroup), true) // TODO lookup needs to use a quick version (skip parsing of specs for descriptions)
	if err != nil {
		return nil, err
	}
	if c, ok := m.Lookup(nameVariantGroup); ok {
		return c, nil
	}
	return nil, fmt.Errorf("unknonw completer/variant: %#v", nameVariantGroup)
}

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
