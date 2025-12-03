package cmd

import (
	"os"
	"path/filepath"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

func loadSpec(path string) (string, *spec.Command, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", nil, err
	}

	content, err := os.ReadFile(abs)
	if err != nil {
		return "", nil, err
	}

	var specCmd spec.Command
	if err := yaml.Unmarshal(content, &specCmd); err != nil {
		return "", nil, err
	}
	return abs, &specCmd, nil
}

func codegen(path string) error {
	// TODO yuck - all this needs some cleanup
	_, spec, err := loadSpec(path)
	if err != nil {
		return err
	}
	return spec.Codegen()
}
