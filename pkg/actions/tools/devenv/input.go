package devenv

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

// ActionInputs completes inputs of the current environment
//
//	devenv (github:cachix/devenv)
//	nixpkgs (github:cachix/devenv-nixpkgs/rolling)
func ActionInputs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		file, err := configFile(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		root := filepath.Dir(file)

		urls := inputUrls(filepath.Join(root, "devenv.yaml"))
		names := lockedInputs(filepath.Join(root, "devenv.lock"))
		if names == nil {
			for name := range urls {
				names = append(names, name)
			}
		}

		vals := make([]string, 0)
		for _, name := range names {
			vals = append(vals, name, urls[name])
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("inputs")
}

// inputUrls returns the url of each input declared in devenv.yaml
func inputUrls(path string) map[string]string {
	content, err := os.ReadFile(path)
	if err != nil {
		return map[string]string{}
	}

	var config struct {
		Inputs map[string]struct {
			URL string `yaml:"url"`
		} `yaml:"inputs"`
	}
	if err := yaml.Unmarshal(content, &config); err != nil {
		return map[string]string{}
	}

	urls := make(map[string]string)
	for name, input := range config.Inputs {
		urls[name] = input.URL
	}
	return urls
}

// lockedInputs returns the top-level inputs of devenv.lock which also contains implicit ones like `devenv`
func lockedInputs(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var lock struct {
		Nodes map[string]struct {
			Inputs map[string]json.RawMessage `json:"inputs"`
		} `json:"nodes"`
		Root string `json:"root"`
	}
	if err := json.Unmarshal(content, &lock); err != nil {
		return nil
	}

	names := make([]string, 0)
	for name := range lock.Nodes[lock.Root].Inputs {
		names = append(names, name)
	}
	if len(names) == 0 {
		return nil
	}
	return names
}
