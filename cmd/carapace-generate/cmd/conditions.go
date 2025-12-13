package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var conditionsCmd = &cobra.Command{
	Use:   "conditions",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := conditions(args[0])
		if err != nil {
			return err
		}

		if f := cmd.Flag("output"); f.Changed {
			if err := os.WriteFile(f.Value.String(), []byte(s), 0644); err != nil {
				return err
			}
			return exec.Command("go", "fmt", f.Value.String()).Run()
		}
		fmt.Println(s)
		return nil

	},
}

func init() {
	carapace.Gen(conditionsCmd).Standalone()
	conditionsCmd.Flags().String("output", "", "file to write to")

	rootCmd.AddCommand(conditionsCmd)
}

func conditions(dir string) (string, error) {
	macros := make([]string, 0)
	descriptions := make(map[string]string, 0)

	r := regexp.MustCompile(`^func Condition(?P<name>[^(]+)\((?P<arg>[^(]*)\) condition.Condition {$`)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error { // TODO walkdir not necessary
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
				} else if after, ok := strings.CutPrefix(t, "// Condition"); ok {
					if splitted := strings.SplitN(after, " ", 2); len(splitted) > 1 {
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

	return content, nil
}
