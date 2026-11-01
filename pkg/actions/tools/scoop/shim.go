package scoop

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionShims completes installed shim names with their source as description
//
//	7z (7zip)
//	python (python)
func ActionShims() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cfg := loadConfig()
		root := cfg.scoopDir()
		if root == "" {
			return carapace.ActionValues()
		}
		global := cfg.globalDir()

		files := shimFiles(root, global)
		sort.Strings(files)

		sources := make(map[string]string)
		names := make([]string, 0)
		for _, file := range files {
			name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			if _, ok := sources[name]; ok {
				continue
			}
			names = append(names, name)
			sources[name] = shimSource(file, root, global)
		}

		sort.Strings(names)
		vals := make([]string, 0, len(names)*2)
		for _, name := range names {
			vals = append(vals, name, sources[name])
		}
		return carapace.ActionValuesDescribed(vals...).Tag("shims")
	})
}

func shimFiles(root, global string) []string {
	files := make([]string, 0)
	for _, dir := range []string{root, global} {
		if dir == "" {
			continue
		}
		_ = filepath.Walk(filepath.Join(dir, "shims"), func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			switch strings.ToLower(filepath.Ext(path)) {
			case ".shim", ".ps1":
				files = append(files, path)
			}
			return nil
		})
	}
	return files
}

func shimSource(path, root, global string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "External"
	}
	content := strings.ReplaceAll(string(data), "\\", "/")

	for _, dir := range []string{root, global} {
		if dir == "" {
			continue
		}
		appsDir := strings.ReplaceAll(filepath.Join(dir, "apps"), "\\", "/")
		pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(appsDir) + `/([^/\s"']+)`)
		if match := pattern.FindStringSubmatch(content); len(match) > 1 {
			return strings.ToLower(match[1])
		}
	}
	return "External"
}
