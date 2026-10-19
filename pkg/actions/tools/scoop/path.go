package scoop

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type config map[string]json.RawMessage

func loadConfig() config {
	paths := make([]string, 0, 2)
	if root := os.Getenv("SCOOP"); root != "" {
		paths = append(paths, filepath.Join(root, "config.json"))
	}
	if home := configHome(); home != "" {
		paths = append(paths, filepath.Join(home, "scoop", "config.json"))
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var cfg config
		if err := json.Unmarshal(bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF}), &cfg); err != nil {
			continue
		}
		return cfg
	}
	return nil
}

func configHome() string {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return dir
	}
	if home, err := os.UserHomeDir(); err == nil {
		return filepath.Join(home, ".config")
	}
	return ""
}

func (c config) string(key string) string {
	for k, raw := range c {
		if strings.EqualFold(k, key) {
			var value string
			if err := json.Unmarshal(raw, &value); err == nil {
				return value
			}
		}
	}
	return ""
}

func (c config) scoopDir() string {
	if dir := os.Getenv("SCOOP"); dir != "" {
		return dir
	}
	if dir := c.string("root_path"); dir != "" {
		return dir
	}
	if home, err := os.UserHomeDir(); err == nil {
		return filepath.Join(home, "scoop")
	}
	return ""
}

func (c config) globalDir() string {
	if dir := os.Getenv("SCOOP_GLOBAL"); dir != "" {
		return dir
	}
	if dir := c.string("global_path"); dir != "" {
		return dir
	}
	if dir := os.Getenv("ProgramData"); dir != "" {
		return filepath.Join(dir, "scoop")
	}
	return ""
}

func (c config) cacheDir() string {
	if dir := os.Getenv("SCOOP_CACHE"); dir != "" {
		return dir
	}
	if dir := c.string("cache_path"); dir != "" {
		return dir
	}
	root := c.scoopDir()
	if root == "" {
		return ""
	}
	return filepath.Join(root, "cache")
}
