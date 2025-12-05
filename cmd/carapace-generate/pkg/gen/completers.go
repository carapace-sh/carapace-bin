package gen

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
)

type Completers map[string][]Completer

func (c Completers) Format(pkg string) string {
	return fmt.Sprintf(`package %s

import (
%s)

%s`, pkg, c.FormatImports(), c.FormateExecute())
}

func (c Completers) FormatImports() string {
	s := make([]string, 0)
	for _, group := range c {
		for _, completer := range group {
			s = append(s, completer.FormatImport())
		}
	}
	sort.Strings(s)
	return strings.Join(s, "\n")
}

func (c Completers) FormateExecute() string {
	cases := make([]string, 0)
	for name, variants := range c {
		switch len(variants) {
		case 0:
		case 1:
			for _, completer := range variants {
				cases = append(cases, fmt.Sprintf("\t\tcase %q:\n\t\t\t%s.Execute()", name, completer.Variable()))
			}
		default:
			s := make([]string, 0)
			s = append(s, fmt.Sprintf("\t\tcase %q:\n\t\t\tswitch variant {", name))
			for _, completer := range variants {
				switch completer.Variant { // TODO incorrect since it's a slice now (TODO how to handle multiple completer with empty variant)
				case "":
					s = append(s, fmt.Sprintf("\t\t\tdefault:\n\t\t\t\t%s.Execute()", completer.Variable()))
				default:
					s = append(s, fmt.Sprintf("\t\t\tcase %q:\n\t\t\t\t%s.Execute()", completer.Variant, completer.Variable()))
				}
			}
			s = append(s, "\t\t\t}")
			cases = append(cases, strings.Join(s, "\n"))
		}
	}
	sort.Strings(cases)

	return fmt.Sprintf(`func executeCompleter(completer, variant string) {
	switch completer {
%s
	}
}`, strings.Join(cases, "\n"))
}

type Completer struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Package     string `json:"package,omitempty"`
	Spec        string `json:"spec,omitempty"`
	Variant     string `json:"variant,omitempty"`
}

func (c Completer) Variable() string {
	if c.Variant == "" {
		return varName(strings.Join([]string{c.Group, c.Name}, "__"))
	}
	return varName(strings.Join([]string{c.Group, c.Name, c.Variant}, "__"))
}

func (c Completer) FormatImport() string {
	return fmt.Sprintf("\t%s %q", c.Variable(), c.Package)
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

	completers := make(map[string][]Completer)
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

func readCompleters(dir, group string) (map[string][]Completer, error) {
	prefix, err := packagePrefix(dir)
	if err != nil {
		return nil, err
	}

	completers := make(map[string][]Completer)
	if files, err := os.ReadDir(filepath.Join(dir, group)); err == nil {
		for _, file := range files {
			if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
				name := strings.TrimSuffix(file.Name(), "_completer")
				if _, err := os.Stat(filepath.Join(dir, group, file.Name(), "cmd/root.go")); err == nil {
					description, err := readDescription(dir, group, file.Name())
					if err != nil {
						return nil, err
					}
					completers[name] = append(completers[name], Completer{
						Name:        name,
						Description: description,
						Group:       group,
						Package:     filepath.Join(prefix, group, file.Name(), "cmd"),
					})
				}

				if variants, err := os.ReadDir(filepath.Join(dir, group, file.Name())); err == nil {
					for _, variant := range variants {
						if _, err := os.Stat(filepath.Join(dir, group, file.Name(), variant.Name(), "cmd/root.go")); err == nil {
							description, err := readDescription(dir, group, file.Name()+"/"+variant.Name()) // TODO just pass the root.go
							if err != nil {
								return nil, err
							}
							completers[name] = append(completers[name], Completer{
								Name:        name,
								Description: description,
								Group:       group,
								Package:     filepath.Join(prefix, group, file.Name(), variant.Name(), "cmd"),
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

func varName(name string) string {
	if name == "go" {
		return "_go"
	}
	if unicode.IsDigit([]rune(name)[0]) {
		name = "_" + name
	}
	return strings.NewReplacer(
		"-", "_",
		".", "_",
	).Replace(name)
}
