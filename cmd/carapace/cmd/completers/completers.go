package completers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Names() []string {
	return names
}

func Description(name string) string {
	if d, err := specDescription(name); err == nil {
		return d
	}
	return descriptions[name]
}

func specDescription(name string) (string, error) {
	confDir, err := os.UserConfigDir()
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
