package completer

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	"github.com/carapace-sh/carapace/pkg/traverse"
)

func ReadCompleters(dir, goos string) (completer.CompleterMap, error) {
	// TODO shell specific completers
	// TODO distro specific completers (arch,ubuntu,...)
	groups := map[string][]string{
		"android": {"common", "unix", "linux", "android"},
		"darwin":  {"common", "unix", "bsd", "darwin"},
		"freebsd": {"common", "unix", "bsd", "freebsd"},
		"linux":   {"common", "unix", "linux"},
		"netbsd":  {"common", "unix", "bsd", "netbsd"},
		"openbsd": {"common", "unix", "bsd", "openbsd"},
		"windows": {"common", "windows"},

		"force_all": {"common", "unix", "linux", "bsd", "darwin", "android", "windows", "freebsd", "netbsd", "openbsd"},
	}

	completers, err := readCompleters(dir)
	if err != nil {
		return nil, err
	}

	filtered := make(completer.CompleterMap)
	for _, group := range groups[goos] {
		filtered.Merge(completers.Filter(choices.Choice{Group: group}))
	}
	return filtered, nil
}

func readCompleters(dir string) (completer.CompleterMap, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	prefix, err := packagePrefix(abs)
	if err != nil {
		return nil, err
	}

	completers := make(completer.CompleterMap)

	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if name, ok := strings.CutSuffix(d.Name(), "_completer"); ok {
			group := filepath.Base(filepath.Dir(path))

			rootPath := filepath.Join(path, "cmd", "root.go")
			if _, err := os.Stat(rootPath); err == nil {
				description, url, err := readRoot(rootPath)
				if err != nil {
					return err
				}
				completers[name] = append(completers[name], completer.Completer{
					Name:        name,
					Description: description,
					Group:       group,
					Package:     filepath.ToSlash(filepath.Join(prefix, group, d.Name(), "cmd")),
					Url:         url,
				})

				return filepath.SkipDir
			}

			if variants, err := os.ReadDir(path); err == nil {
				for _, variant := range variants {
					if !variant.IsDir() {
						continue
					}

					rootPath = filepath.Join(path, variant.Name(), "cmd", "root.go")
					if _, err := os.Stat(path); err == nil {
						// TODO path for variants not correct (description/url)
						description, url, err := readRoot(rootPath)
						if err != nil {
							return err
						}
						completers[name] = append(completers[name], completer.Completer{
							Name:        name,
							Description: description,
							Group:       group,
							Package:     filepath.ToSlash(filepath.Join(prefix, group, d.Name(), variant.Name(), "cmd")),
							Url:         url,
							Variant:     variant.Name(),
						})
					}
				}
			}

			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return completers, nil
}

func readRoot(path string) (string, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	rDescription := regexp.MustCompile("^\tShort: +\"(?P<description>.*)\",$")
	rUrl := regexp.MustCompile("^\tLong: +\"(?P<url>.*)\",$")

	var description, url string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if description != "" && url != "" {
			break
		}
		if description == "" {
			if matches := rDescription.FindStringSubmatch(scanner.Text()); matches != nil {
				description = matches[1]
				continue
			}
		}
		if url == "" {
			if matches := rUrl.FindStringSubmatch(scanner.Text()); matches != nil {
				url = matches[1]
				continue
			}
		}
	}
	return description, url, nil
}

func packagePrefix(dir string) (string, error) {
	content, err := exec.Command("go", "mod", "edit", "-json").Output()
	if err != nil {
		return "", err
	}

	var gomod struct {
		Module struct {
			Path string
		}
	}
	if err := json.Unmarshal(content, &gomod); err != nil {
		return "", err
	}

	c := carapace.NewContext()
	c.Dir, err = c.Abs(dir)
	if err != nil {
		return "", err
	}

	rootDir, err := traverse.Parent("go.mod")(c)
	if err != nil {
		return "", err
	}

	dir, err = filepath.Rel(rootDir, c.Dir)
	if err != nil {
		return "", err
	}
	if dir == "." {
		dir = ""
	}
	return gomod.Module.Path + "/" + filepath.ToSlash(dir), nil
}
