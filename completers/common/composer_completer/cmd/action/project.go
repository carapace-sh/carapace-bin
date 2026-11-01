package action

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// project is a parsed composer.json (or config.json).
// Fields are decoded on demand so that an unexpected type in one of them does not affect the others.
type project map[string]json.RawMessage

// object returns the object stored under given key (nil if missing or not an object).
func (p project) object(key string) map[string]any {
	var v map[string]any
	if raw, ok := p[key]; ok {
		_ = json.Unmarshal(raw, &v)
	}
	return v
}

// strings returns the list of strings stored under given key.
func (p project) strings(key string) []string {
	var v []string
	if raw, ok := p[key]; ok {
		_ = json.Unmarshal(raw, &v)
	}
	return v
}

// text returns the string stored under given key.
func (p project) text(key string) string {
	var v string
	if raw, ok := p[key]; ok {
		_ = json.Unmarshal(raw, &v)
	}
	return v
}

// config returns the "config" object.
func (p project) config() map[string]any {
	return p.object("config")
}

// workingDir returns the absolute path of the directory composer operates in.
func workingDir(c carapace.Context, cmd *cobra.Command) (string, error) {
	dir := "."
	if flag := cmd.Flag("working-dir"); flag != nil && flag.Value.String() != "" {
		dir = flag.Value.String()
	}
	return c.Abs(dir)
}

// composerFile returns the path of the composer.json (honouring $COMPOSER).
func composerFile(c carapace.Context, cmd *cobra.Command) (string, error) {
	dir, err := workingDir(c, cmd)
	if err != nil {
		return "", err
	}

	name := c.Getenv("COMPOSER")
	if name == "" {
		name = "composer.json"
	}
	if filepath.IsAbs(name) {
		return name, nil
	}
	return filepath.Join(dir, name), nil
}

// lockFile returns the path of the lock file belonging to given composer.json.
func lockFile(composerFile string) string {
	if filepath.Ext(composerFile) == ".json" {
		return strings.TrimSuffix(composerFile, "json") + "lock"
	}
	return composerFile + ".lock"
}

// configFile returns the file that `config`, `repository` and `policy` operate on.
func configFile(c carapace.Context, cmd *cobra.Command) (string, error) {
	if flag := cmd.Flag("global"); flag != nil && flag.Value.String() == "true" {
		home, err := composerHome(c)
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "config.json"), nil
	}

	if flag := cmd.Flag("file"); flag != nil && flag.Value.String() != "" {
		return c.Abs(flag.Value.String())
	}
	return composerFile(c, cmd)
}

// composerHome returns the global composer directory.
func composerHome(c carapace.Context) (string, error) {
	if home := c.Getenv("COMPOSER_HOME"); home != "" {
		return home, nil
	}

	if runtime.GOOS == "windows" {
		appData := c.Getenv("APPDATA")
		if appData == "" {
			return "", errors.New("APPDATA or COMPOSER_HOME must be set")
		}
		return filepath.Join(appData, "Composer"), nil
	}

	userDir := strings.TrimRight(c.Getenv("HOME"), "/")
	if userDir == "" {
		return "", errors.New("HOME or COMPOSER_HOME must be set")
	}

	dirs := make([]string, 0, 2)
	if useXdg(c) {
		xdgConfig := c.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			xdgConfig = userDir + "/.config"
		}
		dirs = append(dirs, xdgConfig+"/composer")
	}
	dirs = append(dirs, userDir+"/.composer")

	for _, dir := range dirs {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir, nil
		}
	}
	return dirs[0], nil
}

func useXdg(c carapace.Context) bool {
	for _, env := range c.Env {
		if strings.HasPrefix(env, "XDG_") {
			return true
		}
	}
	info, err := os.Stat("/etc/xdg")
	return err == nil && info.IsDir()
}

// readProject reads and parses given composer.json.
func readProject(path string) (project, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(filepath.Base(path) + " not found")
		}
		return nil, err
	}

	var p project
	if err := json.Unmarshal(content, &p); err != nil {
		return nil, errors.New(filepath.Base(path) + ": " + err.Error())
	}
	return p, nil
}

// loadProject reads the composer.json composer would use.
func loadProject(c carapace.Context, cmd *cobra.Command) (project, error) {
	path, err := composerFile(c, cmd)
	if err != nil {
		return nil, err
	}
	return readProject(path)
}

// vendorDir returns the absolute vendor directory.
func vendorDir(c carapace.Context, cmd *cobra.Command, p project) (string, error) {
	dir, err := workingDir(c, cmd)
	if err != nil {
		return "", err
	}

	vendor := c.Getenv("COMPOSER_VENDOR_DIR")
	if vendor == "" {
		vendor, _ = p.config()["vendor-dir"].(string)
	}
	if vendor == "" {
		vendor = "vendor"
	}
	if filepath.IsAbs(vendor) {
		return vendor, nil
	}
	return filepath.Join(dir, vendor), nil
}

// oneLine reduces a text to its first line.
func oneLine(s string) string {
	s, _, _ = strings.Cut(strings.TrimSpace(s), "\n")
	return strings.TrimSpace(s)
}
