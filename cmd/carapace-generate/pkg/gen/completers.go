package gen

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type Completers map[string]map[string]Completer

func (c Completers) Format() string {
	return ""
}

type Completer struct {
	Name        string
	Description string
	Group       string
	Package     string
	Variant     string
}

func ReadCompleters(dir, goos string) (Completers, error) {
	// TODO shell specific completers
	// TODO distro specific completers (arch,ubuntu,...)
	// TODO variants (tldr like tealdear)
	groups := map[string][]string{
		"linux":   {"common", "unix", "linux"},
		"darwin":  {"common", "unix", "darwin"},
		"windows": {"common", "windows"},
	}

	completers := make(map[string]map[string]Completer)
	for _, group := range groups[goos] {
		groupCompleters, err := readCompleters(dir, group)
		if err != nil {
			return nil, err
		}
		maps.Copy(completers, groupCompleters) // TODO variants might get lost here, need to be handled explicitly
	}
	// TODO map value should be a list of completers (variants)
	return completers, nil
}

func readCompleters(dir, group string) (map[string]map[string]Completer, error) {
	prefix, err := packagePrefix(dir)
	if err != nil {
		return nil, err
	}

	completers := make(map[string]map[string]Completer)
	if files, err := os.ReadDir(filepath.Join(dir, group)); err == nil {
		for _, file := range files {
			if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
				name := strings.TrimSuffix(file.Name(), "_completer")
				name, variant, _ := strings.Cut(name, ".")
				description, err := readDescription(dir, group, file.Name())
				if err != nil {
					return nil, err
				}
				if completers[name] == nil {
					completers[name] = map[string]Completer{}
				}
				completers[name][variant] = Completer{
					Name:        name,
					Description: description,
					Group:       group,
					Package:     filepath.Join(prefix, group, file.Name(), "cmd"),
					Variant:     variant,
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

	dir = filepath.Dir(dir)
	if dir == "." {
		dir = ""
	}
	return gomod.Module.Path + "/" + filepath.ToSlash(dir), nil
}
