package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/carapace-sh/carapace/pkg/util"
)

func main() {
	macros()
	conditions()

	goos := "common" // TODO support other goos
	names, descriptions := readCompleters(goos)

	imports := make([]string, 0, len(names))
	for _, name := range names {
		imports = append(imports, fmt.Sprintf(`	%v "github.com/carapace-sh/carapace-bin/completers/%v/%v_completer/cmd"`, varName(name), goos, name))
	}

	formattedNames := make([]string, 0)
	for _, name := range names {
		formattedNames = append(formattedNames, fmt.Sprintf("		\"%v\",", name))
	}

	maxlen := 0
	for _, name := range names {
		if l := len(name); l > maxlen {
			maxlen = l
		}
	}

	formattedDescriptions := make([]string, 0)
	for _, name := range names {
		formattedDescriptions = append(formattedDescriptions, fmt.Sprintf(`		%--`+strconv.Itoa(maxlen+4)+`v"%v",`, fmt.Sprintf(`"%v": `, name), descriptions[name]))
	}

	cases := make([]string, 0)
	for _, name := range names {
		cases = append(cases,
			fmt.Sprintf(`	case "%v":`, name),
			fmt.Sprintf(`		%v.Execute()`, varName(name)),
		)
	}

	content := fmt.Sprintf(`package cmd

import (
%v)

func executeCompleter(completer string) {
	switch completer {
%v	}
}
`,
		fmt.Sprintln(strings.Join(imports, "\n")),
		fmt.Sprintln(strings.Join(cases, "\n")),
	)

	root, err := rootDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Mkdir(root+"/cmd/carapace/cmd/completers", 0755)
	os.WriteFile(root+"/cmd/carapace/cmd/completers.go", []byte("//go:build !release\n\n"+content), 0644)
	os.WriteFile(root+"/cmd/carapace/cmd/completers/name.go", []byte(fmt.Sprintf("package completers\n\nfunc init() {\n	names = []string{\n%v\n\t}\n}\n", strings.Join(formattedNames, "\n"))), 0644)
	os.WriteFile(root+"/cmd/carapace/cmd/completers/description.go", []byte(fmt.Sprintf("package completers\n\nfunc init() {\n	descriptions = map[string]string{\n%v\n\t}\n}\n", strings.Join(formattedDescriptions, "\n"))), 0644)
	os.WriteFile(root+"/cmd/carapace/cmd/completers_release.go", []byte("//go:build release\n\n"+strings.Replace(content, "/completers/", "/completers_release/", -1)), 0644)
	os.RemoveAll(root + "/completers_release")
	execlog.Command("cp", "-r", root+"/completers", root+"/completers_release").Run()

	for _, name := range names {
		files, err := os.ReadDir(fmt.Sprintf("%v/completers_release/%v/%v_completer/cmd/", root, goos, name))
		if err == nil {
			initFuncs := make([]string, 0)
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
					path := fmt.Sprintf("%v/completers_release/%v/%v_completer/cmd/%v", root, goos, name, file.Name())
					content, err := os.ReadFile(path)
					if err == nil && strings.Contains(string(content), "func init() {") {
						patched := strings.Replace(string(content), "func init() {", fmt.Sprintf("func init_%v() {", strings.TrimSuffix(file.Name(), ".go")), 1)
						os.WriteFile(path, []byte(patched), os.ModePerm)
						initFuncs = append(initFuncs, fmt.Sprintf("	init_%v()", strings.TrimSuffix(file.Name(), ".go")))
					}
				}
			}

			path := fmt.Sprintf("%v/completers_release/%v/%v_completer/cmd/root.go", root, goos, name)
			content, err := os.ReadFile(path)
			if err == nil {
				patched := make([]string, 0)
				for _, line := range strings.Split(string(content), "\n") {
					patched = append(patched, line)
					if line == "func Execute() error {" {
						patched = append(patched, initFuncs...)
					}
				}
				os.WriteFile(path, []byte(strings.Join(patched, "\n")), os.ModePerm)
			}
		}
	}
}

func varName(name string) string {
	if name == "go" {
		return "_go"
	}
	return strings.NewReplacer(
		"-", "_",
		".", "_",
	).Replace(name)
}

func readCompleters(goos string) ([]string, map[string]string) {
	names := make([]string, 0)
	descriptions := make(map[string]string)
	if root, err := rootDir(); err == nil {
		if files, err := os.ReadDir(filepath.Join(root, "completers", goos)); err == nil {
			for _, file := range files {
				if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
					name := strings.TrimSuffix(file.Name(), "_completer")
					description := readDescription(root, goos, file.Name())
					names = append(names, name)
					descriptions[name] = description
				}
			}
		}
	}
	return names, descriptions
}

func readDescription(root, goos, completer string) string {
	if content, err := os.ReadFile(fmt.Sprintf("%v/completers/%v/%v/cmd/root.go", root, goos, completer)); err == nil {
		re := regexp.MustCompile("^\tShort: +\"(?P<description>.*)\",$")
		for _, line := range strings.Split(string(content), "\n") {
			if re.MatchString(line) {
				return re.FindStringSubmatch(line)[1]
			}
		}
	}
	return ""
}

func rootDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path, err := util.FindReverse(wd, "go.mod")
	if err != nil {
		return "", err
	}

	return filepath.ToSlash(filepath.Dir(path)), nil
}

func macros() {
	root, err := rootDir()
	if err != nil {
		panic(err.Error)
	}

	// hardcoded the bridge macros
	imports := map[string]bool{
		`bridge "github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"`: true,
	}
	macros := []string{
		`"bridge.Argcomplete":         spec.MacroV(bridge.ActionArgcomplete),`,
		`"bridge.ArgcompleteV1":       spec.MacroV(bridge.ActionArgcompleteV1), //lint:ignore SA1019 gcloud uses an old argcomplete version`,
		`"bridge.Bash":                spec.MacroV(bridge.ActionBash),`,
		`"bridge.Carapace":            spec.MacroV(bridge.ActionCarapace),`,
		`"bridge.CarapaceBin":         spec.MacroV(bridge.ActionCarapaceBin),`,
		`"bridge.Clap":                spec.MacroV(bridge.ActionClap),`,
		`"bridge.Click":               spec.MacroV(bridge.ActionClick),`,
		`"bridge.Cobra":               spec.MacroV(bridge.ActionCobra),`,
		`"bridge.Complete":            spec.MacroV(bridge.ActionComplete),`,
		`"bridge.Fish":                spec.MacroV(bridge.ActionFish),`,
		`"bridge.Inshellisense":       spec.MacroV(bridge.ActionInshellisense),`,
		`"bridge.Kingpin":             spec.MacroV(bridge.ActionKingpin),`,
		`"bridge.Macro":               spec.MacroV(bridge.ActionMacro),`,
		`"bridge.Powershell":          spec.MacroV(bridge.ActionPowershell),`,
		`"bridge.Urfavecli":           spec.MacroV(bridge.ActionUrfavecli),`,
		`"bridge.Yargs":               spec.MacroV(bridge.ActionYargs),`,
		`"bridge.Zsh":                 spec.MacroV(bridge.ActionZsh),`,
	}
	descriptions := map[string]string{
		"bridge.Argcomplete":   "bridges https://github.com/kislyuk/argcomplete",
		"bridge.ArgcompleteV1": "bridges https://github.com/kislyuk/argcomplete (v1)",
		"bridge.Bash":          "bridges https://www.gnu.org/software/bash/",
		"bridge.Carapace":      "bridges https://github.com/carapace-sh/carapace",
		"bridge.CarapaceBin":   "bridges https://github.com/carapace-sh/carapace-bin",
		"bridge.Clap":          "bridges https://github.com/clap-rs/clap",
		"bridge.Click":         "bridges https://github.com/pallets/click",
		"bridge.Cobra":         "bridges https://github.com/spf13/cobra",
		"bridge.Complete":      "bridges https://github.com/posener/complete",
		"bridge.Fish":          "bridges https://fishshell.com/",
		"bridge.Inshellisense": "bridges https://github.com/microsoft/inshellisense",
		"bridge.Kingpin":       "bridges https://github.com/alecthomas/kingpin",
		"bridge.Macro":         "bridges macros exposed with https://github.com/carapace-sh/carapace-spec",
		"bridge.Powershell":    "bridges https://microsoft.com/powershell",
		"bridge.Urfavecli":     "bridges https://github.com/urfave/cli",
		"bridge.Yargs":         "bridges https://github.com/yargs/yargs",
		"bridge.Zsh":           "bridges https://www.zsh.org/",
	}

	r := regexp.MustCompile(`^func Action(?P<name>[^(]+)\((?P<arg>[^(]*)\) carapace.Action {$`)
	filepath.WalkDir(root+"/pkg/actions", func(path string, d fs.DirEntry, err error) error {
		path = filepath.ToSlash(path)
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			pkg := strings.Replace(filepath.ToSlash(filepath.Dir(strings.TrimPrefix(path, root+"/pkg/actions/"))), "/", ".", -1)
			_import := fmt.Sprintf(`	%v "github.com/carapace-sh/carapace-bin/pkg/actions/%v"`, strings.Replace(pkg, ".", "_", -1), strings.Replace(pkg, ".", "/", -1))

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if t := scanner.Text(); strings.HasPrefix(t, "func Action") {
					if r.MatchString(t) {
						matches := r.FindStringSubmatch(t)

						_func := fmt.Sprintf("%v.Action%v", strings.Replace(pkg, ".", "_", -1), matches[1])

						disableflagParsing := ""
						if pkg == "bridge" {
							disableflagParsing = ".NoFlag()"
						}

						if arg := matches[2]; strings.Contains(arg, ",") {
							macros = append(macros, "// TODO unsupported signature: "+t)
							continue
						} else if arg == "" {
							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroN(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
						} else if strings.Contains(arg, "...") {
							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroV(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
						} else {
							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroI(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
						}
						imports[_import] = true
					}
				} else if strings.HasPrefix(t, "// Action") {
					if splitted := strings.SplitN(strings.TrimPrefix(t, "// Action"), " ", 2); len(splitted) > 1 {
						descriptions[pkg+"."+splitted[0]] = splitted[1]
					}
				}
			}

		}
		return nil
	})

	sortedImports := make([]string, 0)
	for i := range imports {
		sortedImports = append(sortedImports, i)
	}
	sort.Strings(sortedImports)

	sortedDescriptions := make([]string, 0)
	for key, value := range descriptions {
		sortedDescriptions = append(sortedDescriptions, fmt.Sprintf(`%#v: %#v,`, key, value))
	}
	sort.Strings(sortedDescriptions)

	content := fmt.Sprintf(`package actions

import (
%v
	spec "github.com/carapace-sh/carapace-spec"
)

func init() {
	MacroMap = map[string]spec.Macro{
%v
	}

	MacroDescriptions = map[string]string {
%v
	}
}
`, strings.Join(sortedImports, "\n"), strings.Join(macros, "\n"), strings.Join(sortedDescriptions, "\n"))

	os.WriteFile(root+"/pkg/actions/actions_generated.go", []byte(content), 0644)
	execlog.Command("go", "fmt", root+"/pkg/actions/actions_generated.go").Run()

}

func conditions() {
	root, err := rootDir()
	if err != nil {
		panic(err.Error)
	}

	macros := make([]string, 0)
	descriptions := make(map[string]string, 0)

	r := regexp.MustCompile(`^func Condition(?P<name>[^(]+)\((?P<arg>[^(]*)\) condition.Condition {$`)
	filepath.WalkDir(root+"/pkg/conditions", func(path string, d fs.DirEntry, err error) error { // TODO walkdir not necessary
		path = filepath.ToSlash(path)
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if t := scanner.Text(); strings.HasPrefix(t, "func Condition") {
					if r.MatchString(t) {
						matches := r.FindStringSubmatch(t)

						_func := fmt.Sprintf("Condition%v", matches[1])

						if arg := matches[2]; strings.Contains(arg, ",") {
							macros = append(macros, "// TODO unsupported signature: "+t)
							continue
						} else if arg == "" {
							macros = append(macros, fmt.Sprintf(`"%v": condition.MacroN(%v).WithDescription(%#v),`, matches[1], _func, descriptions[matches[1]]))
						} else if strings.Contains(arg, "...") {
							macros = append(macros, fmt.Sprintf(`"%v": condition.MacroV(%v).WithDescription(%#v),`, matches[1], _func, descriptions[matches[1]]))
						} else {
							macros = append(macros, fmt.Sprintf(`"%v": condition.MacroI(%v).WithDescription(%#v),`, matches[1], _func, descriptions[matches[1]]))
						}
					}
				} else if strings.HasPrefix(t, "// Condition") {
					if splitted := strings.SplitN(strings.TrimPrefix(t, "// Condition"), " ", 2); len(splitted) > 1 {
						descriptions[splitted[0]] = splitted[1]
					}
				}
			}

		}
		return nil
	})

	content := fmt.Sprintf(`package conditions

import (
	"github.com/carapace-sh/carapace-bin/internal/condition"
	"github.com/carapace-sh/carapace-spec/pkg/macro"
)

func init() {
	MacroMap = macro.MacroMap[condition.Macro]{
%v
	}
}
`, strings.Join(macros, "\n"))

	os.WriteFile(root+"/pkg/conditions/conditions_generated.go", []byte(content), 0644)
	execlog.Command("go", "fmt", root+"/pkg/conditions/conditions_generated.go").Run()

}
