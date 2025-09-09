package completers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/env"
	"github.com/carapace-sh/carapace/pkg/xdg"
	"gopkg.in/yaml.v3"
)

var (
	descriptions = make(map[string]string)
	names        = make([]string, 0)
)

func Names() []string {
	excludes := make(map[string]bool)
	for _, e := range env.Excludes() {
		excludes[e] = true
	}

	unique := map[string]bool{
		"carapace": true,
	}

	if os.Getenv("CARAPACE_ENV") != "0" {
		unique["get-env"] = true
		unique["set-env"] = true
		unique["unset-env"] = true
	}

	if _, ok := excludes["*"]; !ok {
		for _, name := range append(names, "python/http.server") { // TODO extensions
			if _, ok := excludes[name]; !ok {
				unique[name] = true
			}
		}
	}
	//if specNames, err := Specs(); err == nil {
	if specNames, _ := Specs(); true { // TODO use and handle err
		for _, name := range specNames {
			unique[name] = true
		}
	}

	combined := make([]string, 0, len(unique))
	for name := range unique {
		combined = append(combined, name)
	}
	sort.Strings(combined)
	return combined
}

func SpecPath(name string) (string, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%v/carapace/specs/%v.yaml", configDir, name)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func OverlayPath(name string) (string, error) {
	// TODO code duplication
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%v/carapace/overlays/%v.yaml", configDir, name)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func Specs() (specs []string, dir string) {
	r := regexp.MustCompile(`^[0-9a-zA-Z_\-.]+\.yaml$`) // sanity check
	specs = make([]string, 0)
	if configDir, err := xdg.UserConfigDir(); err == nil {
		dir = configDir + "/carapace/specs"
		if entries, err := os.ReadDir(dir); err == nil {
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") && r.MatchString(entry.Name()) {
					specs = append(specs, strings.TrimSuffix(entry.Name(), ".yaml"))
				}
			}
		} else if os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
	return
}

func Description(name string) string {
	if d, err := specDescription(name); err == nil {
		return d
	}
	return descriptions[name]
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
