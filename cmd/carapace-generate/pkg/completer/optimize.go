package completer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Optimize(dir string) error {
	source, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	target := source + "_release"
	if err := os.RemoveAll(target); err != nil { // TODO dangerous, but limited to `_release suffix`
		return err
	}

	if err := os.Mkdir(target, os.ModePerm); err != nil {
		return err
	}

	prefix, err := packagePrefix(source)
	if err != nil {
		return err
	}
	importReplacer := strings.NewReplacer(
		prefix,
		filepath.ToSlash(filepath.Join(filepath.Dir(prefix), filepath.Base(target))), // TODO not always correct
	)
	return filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// TODO update import
		content, err := os.ReadFile(path) // TODO this fully reads file into memory
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".go" {
			patched := importReplacer.Replace(string(content)) // TODO potential issues when prefix is used otherwise than import
			if strings.HasSuffix(filepath.ToSlash(filepath.Dir(path)), "/cmd") {
				switch filepath.Base(path) {
				case "root.go":
					// TODO assumes all go files are subcommands and have an init function
					entries, err := os.ReadDir(filepath.Dir(path))
					if err != nil {
						return err
					}

					initFuncs := make([]string, 0)
					for _, entry := range entries {
						if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
							initFuncs = append(initFuncs, fmt.Sprintf("	init_%v()", strings.TrimSuffix(entry.Name(), ".go")))
						}
					}
					patched = strings.Replace(patched, "\nfunc Execute() error {", "\nfunc Execute() error {\n"+strings.Join(initFuncs, "\n"), 1)
					patched = strings.Replace(patched, "\nfunc init() {", fmt.Sprintf("\nfunc init_%v() {", strings.TrimSuffix(info.Name(), ".go")), 1)

				default:
					patched = strings.Replace(patched, "\nfunc init() {", fmt.Sprintf("\nfunc init_%v() {", strings.TrimSuffix(info.Name(), ".go")), 1)
					// initFuncs = append(initFuncs, fmt.Sprintf("	init_%v()", strings.TrimSuffix(file.Name(), ".go")))
				}
			}
			content = []byte(patched)
		}

		targetPath := filepath.Join(target, strings.TrimPrefix(path, source))
		if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
			return err
		}
		return os.WriteFile(targetPath, content, info.Mode())
	})
}
