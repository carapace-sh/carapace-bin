package completer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace/pkg/traverse"
)

func ReadCompleters(dir, goos string) (completer.CompleterMap, error) {
	// TODO shell specific completers
	// TODO distro specific completers (arch,ubuntu,...)
	// TODO variants (tldr like tealdear)
	groups := map[string][]string{
		"linux":     {"common", "unix", "linux"},
		"darwin":    {"common", "unix", "darwin"},
		"windows":   {"common", "windows"},
		"force_all": {"common", "unix", "linux", "darwin", "windows"},
	}

	completers := make(completer.CompleterMap)
	for _, group := range groups[goos] {
		groupCompleters, err := readCompleters(dir, group)
		if err != nil {
			return nil, err
		}
		completers.Merge(groupCompleters)
	}
	return completers, nil
}

func readCompleters(dir, group string) (completer.CompleterMap, error) {
	prefix, err := packagePrefix(dir)
	if err != nil {
		return nil, err
	}

	completers := make(completer.CompleterMap)
	if files, err := os.ReadDir(filepath.Join(dir, group)); err == nil {
		for _, file := range files {
			if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
				name := strings.TrimSuffix(file.Name(), "_completer")
				if _, err := os.Stat(filepath.Join(dir, group, file.Name(), "cmd/root.go")); err == nil {
					description, err := readDescription(dir, group, file.Name())
					if err != nil {
						return nil, err
					}
					url, err := readUrl(dir, group, file.Name())
					if err != nil {
						return nil, err
					}
					completers[name] = append(completers[name], completer.Completer{
						Name:        name,
						Description: description,
						Group:       group,
						Package:     filepath.Join(prefix, group, file.Name(), "cmd"),
						Url:         url,
					})
				}

				if variants, err := os.ReadDir(filepath.Join(dir, group, file.Name())); err == nil {
					for _, variant := range variants {
						if _, err := os.Stat(filepath.Join(dir, group, file.Name(), variant.Name(), "cmd/root.go")); err == nil {
							// TODO path for variants not correct (description/url)
							description, err := readDescription(dir, group, file.Name()+"/"+variant.Name()) // TODO just pass the root.go
							if err != nil {
								return nil, err
							}
							url, err := readUrl(dir, group, file.Name()+"/"+variant.Name())
							if err != nil {
								return nil, err
							}
							completers[name] = append(completers[name], completer.Completer{
								Name:        name,
								Description: description,
								Group:       group,
								Package:     filepath.Join(prefix, group, file.Name(), variant.Name(), "cmd"),
								Url:         url,
								Variant:     variant.Name(),
							})
						}
					}
				}
			}
		}
	}
	return completers, nil
}

func readDescription(root, goos, completer string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("%v/%v/%v/cmd/root.go", root, goos, completer))
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("^\tShort: +\"(?P<description>.*)\",$")
	for _, line := range strings.Split(string(content), "\n") {
		if re.MatchString(line) {
			return re.FindStringSubmatch(line)[1], nil
		}
	}
	return "", nil // TODO error?
}

func readUrl(root, goos, completer string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("%v/%v/%v/cmd/root.go", root, goos, completer))
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("^\tLong: +\"(?P<url>.*)\",$")
	for _, line := range strings.Split(string(content), "\n") {
		if re.MatchString(line) {
			return re.FindStringSubmatch(line)[1], nil
		}
	}
	return "", nil // TODO error?
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
