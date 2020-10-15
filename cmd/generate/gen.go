package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	completers := readCompleters()

	imports := make([]string, len(completers))
	for index, c := range readCompleters() {
		imports[index] = fmt.Sprintf(`	%v "github.com/rsteube/carapace-bin/completers/%v_completer/cmd"`, varName(c), c)
	}

	vars := make([]string, len(completers))
	for index, c := range readCompleters() {
		vars[index] = fmt.Sprintf(`	"%v",`, c)
	}

	cases := make([]string, len(completers)*2)
	for index, c := range readCompleters() {
		cases[index*2] = fmt.Sprintf(`	case "%v":`, c)
		cases[(index*2)+1] = fmt.Sprintf(`		%v.Execute()`, varName(c))
	}

	content := fmt.Sprintf(`package cmd

import (
%v)

var completers = []string{
%v}

func executeCompleter(completer string) {
	switch completer {
%v	}
}
`,
		fmt.Sprintln(strings.Join(imports, "\n")),
		fmt.Sprintln(strings.Join(vars, "\n")),
		fmt.Sprintln(strings.Join(cases, "\n")),
	)
	if root, err := rootDir(); err == nil {
		ioutil.WriteFile(root+"/cmd/carapace/cmd/completers.go", []byte(content), 0644)
	}
}

func varName(name string) string {
    if name == "go" {
        return "_go"
    }
	return strings.Replace(name, "-", "_", -1)
}

func readCompleters() []string {
	completers := make([]string, 0)
	if root, err := rootDir(); err == nil {
		if files, err := ioutil.ReadDir(root + "/completers/"); err == nil {
			for _, file := range files {
				if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
					completers = append(completers, strings.TrimSuffix(file.Name(), "_completer"))
				}
			}
		}
	}
	return completers
}

func rootDir() (string, error) {
	if output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}
