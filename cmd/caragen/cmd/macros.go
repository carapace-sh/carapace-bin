package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "list macros",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := json.Marshal(macros())
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(m))
		return nil
	},
}

func init() {
	carapace.Gen(macrosCmd).Standalone()

	rootCmd.AddCommand(macrosCmd)
}

type macro struct {
	Path        string
	ImportPath  string
	Package     string
	Function    string
	Description string
}

func (m macro) FlagParsingDisabled() bool {
	return m.Package == "bridge"
}

func macros() []macro {
	root, err := rootDir()
	if err != nil {
		panic(err.Error)
	}

	macros := make([]macro, 0)
	descriptions := make(map[string]string)

	r := regexp.MustCompile(`^func Action(?P<name>[^(]+)\((?P<arg>[^(]*)\) carapace.Action {$`)
	filepath.WalkDir(root+"/pkg/actions", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				_macro := &macro{}
				_macro.Package = strings.Replace(filepath.Dir(strings.TrimPrefix(path, root+"/pkg/actions/")), "/", ".", -1)

				if t := scanner.Text(); strings.HasPrefix(t, "func Action") {
					if r.MatchString(t) {
						matches := r.FindStringSubmatch(t)
						//_macro.ImportPath = fmt.Sprintf(`	%v "github.com/rsteube/carapace-bin/pkg/actions/%v"`, strings.Replace(_macro.Package, ".", "_", -1), strings.Replace(_macro.Package, ".", "/", -1))
						_macro.ImportPath = fmt.Sprintf("github.com/rsteube/carapace-bin/pkg/actions/%v", strings.Replace(_macro.Package, ".", "/", -1))
						_macro.Function = fmt.Sprintf("%v.Action%v", strings.Replace(_macro.Package, ".", "_", -1), matches[1])
						macros = append(macros, *_macro)
					}
				} else if strings.HasPrefix(t, "// Action") {
					if splitted := strings.SplitN(strings.TrimPrefix(t, "// Action"), " ", 2); len(splitted) > 1 {
						descriptions[_macro.Package+"."+splitted[0]] = splitted[1]
					}
				}
			}

		}
		return nil
	})
	return macros
}
