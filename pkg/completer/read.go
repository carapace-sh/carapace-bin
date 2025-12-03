package completer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func ReadSpecs(dir, group string, parse bool) (CompleterMap, error) {
	var err error
	if dir, err = filepath.Abs(dir); err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`[^.]+\.yaml`)

	m := make(CompleterMap)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var _spec struct {
			Name        string
			Description string
		}

		switch {
		case r.MatchString(entry.Name()):
			if !parse {
				m[_spec.Name] = append(m[_spec.Name], Completer{
					Name:    _spec.Name,
					Group:   group,
					Spec:    filepath.Join(dir, entry.Name()),
					Execute: complete(_spec.Name, spec.ActionSpec(filepath.Join(dir, entry.Name()))),
				})
				continue
			}

			content, err := os.ReadFile(filepath.Join(dir, entry.Name())) // TODO slow - only need to read name/description
			if err != nil {
				return nil, err
			}

			if err := yaml.Unmarshal(content, &_spec); err != nil {
				return nil, err
			}

			if entry.Name() != _spec.Name+".yaml" {
				return nil, fmt.Errorf("spec filename %#v does not match name %#v", entry.Name(), _spec.Name)
			}

			m[_spec.Name] = append(m[_spec.Name], Completer{
				Name:        _spec.Name,
				Description: _spec.Description,
				Group:       group,
				Spec:        filepath.Join(dir, entry.Name()),
				Execute:     complete(_spec.Name, spec.ActionSpec(filepath.Join(dir, entry.Name()))),
			})

		case entry.IsDir():
			variants, err := os.ReadDir(filepath.Join(dir, entry.Name()))
			if err != nil {
				return nil, err
			}

			for _, variant := range variants {
				if !r.MatchString(variant.Name()) {
					continue
				}
				if !parse {
					m[_spec.Name] = append(m[_spec.Name], Completer{
						Name:    _spec.Name,
						Group:   group,
						Spec:    filepath.Join(dir, entry.Name(), variant.Name()),
						Variant: strings.TrimSuffix(variant.Name(), ".yaml"),
						Execute: func() error {
							// TODO load/execute spec
							return nil
						},
					})
					continue
				}

				content, err := os.ReadFile(filepath.Join(dir, entry.Name(), variant.Name())) // TODO slow - only need to read name/description
				if err != nil {
					return nil, err
				}

				if err := yaml.Unmarshal(content, &_spec); err != nil {
					return nil, err
				}

				m[_spec.Name] = append(m[_spec.Name], Completer{
					Name:        _spec.Name,
					Description: _spec.Description,
					Group:       group,
					Spec:        filepath.Join(dir, entry.Name(), variant.Name()),
					Variant:     strings.TrimSuffix(variant.Name(), ".yaml"),
					Execute: func() error {
						// TODO load/execute spec
						return nil
					},
				})
			}
		}
	}
	return m, nil
}

// TODO duplicated from cmd/carapace/cmd/completer/
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
