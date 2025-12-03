package gen

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Completer struct {
	Name        string
	Description string
	Group       string
	Package     string
}

func Completers(dir, goos string) (map[string]Completer, error) {
	// TODO shell specific completers
	groups := map[string][]string{
		"linux":   {"common", "unix", "linux"},
		"darwin":  {"common", "unix", "darwin"},
		"windows": {"common", "windows"},
	}

	completers := make(map[string]Completer)
	for _, group := range groups[goos] {
		groupCompleters, err := readCompleters(dir, group)
		if err != nil {
			return nil, err
		}
		maps.Copy(completers, groupCompleters)
	}
	return completers, nil
}

func readCompleters(dir, group string) (map[string]Completer, error) {
	completers := make(map[string]Completer)

	if files, err := os.ReadDir(filepath.Join(dir, group)); err == nil {
		for _, file := range files {
			if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
				name := strings.TrimSuffix(file.Name(), "_completer")
				description, err := readDescription(dir, group, file.Name())
				if err != nil {
					return nil, err
				}
				completers[name] = Completer{
					Name:        name,
					Description: description,
					Group:       group,
					// TODO package
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
