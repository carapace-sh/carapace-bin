package completer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadSpecs(dir, group string) (CompleterMap, error) {
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
		var spec struct {
			Name        string
			Description string
		}

		switch {
		case r.MatchString(entry.Name()):
			content, err := os.ReadFile(filepath.Join(dir, entry.Name())) // TODO slow - only need to read name/description
			if err != nil {
				return nil, err
			}

			if err := yaml.Unmarshal(content, &spec); err != nil {
				return nil, err
			}

			if entry.Name() != spec.Name+".yaml" {
				return nil, fmt.Errorf("spec filename %#v does not match name %#v", entry.Name(), spec.Name)
			}

			m[spec.Name] = append(m[spec.Name], Completer{
				Name:        spec.Name,
				Description: spec.Description,
				Group:       group,
				Spec:        filepath.Join(dir, entry.Name()),
				Execute: func() error {
					// TODO load/execute spec
					return nil
				},
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

				content, err := os.ReadFile(filepath.Join(dir, entry.Name(), variant.Name())) // TODO slow - only need to read name/description
				if err != nil {
					return nil, err
				}

				if err := yaml.Unmarshal(content, &spec); err != nil {
					return nil, err
				}

				m[spec.Name] = append(m[spec.Name], Completer{
					Name:        spec.Name,
					Description: spec.Description,
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
