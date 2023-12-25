package cmd

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/spf13/cobra"
)

var conditionCmd = &cobra.Command{
	Use:   "--condition [condition]",
	Short: "list or execute condition",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			printConditions()
		case 1:
			printCondition(strings.SplitN(args[0], "(", 2)[0])
		}
	},
}

func init() {
	carapace.Gen(conditionCmd).Standalone()
	macroCmd.Flags().SetInterspersed(false)

	carapace.Gen(conditionCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			vals := make([]string, 0)
			for name, condition := range conditions.MacroMap {
				vals = append(vals, name, condition.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			m, err := conditions.MacroMap.Lookup("$" + c.Args[0])
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			condition := m.Parse("$" + c.Args[0])

			return carapace.ActionValues(strconv.FormatBool(condition(c)))
		}),
	)
}

func printConditions() {
	maxlen := 0
	names := make([]string, 0)
	for name := range conditions.MacroMap {
		names = append(names, name)
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, conditions.MacroMap[name].Description)
	}
}

func printCondition(name string) {
	if m, ok := conditions.MacroMap[name]; ok {
		path := strings.Replace(name, ".", "/", -1)
		signature := ""
		if s := m.Signature(); s != "" {
			signature = fmt.Sprintf("(%v)", s)
		}

		fmt.Printf(`signature:   $%v%v
description: %v
reference:   https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/conditions/%v#Condition%v
`, name, signature, m.Description, filepath.Dir(path), filepath.Base(path))
	}
}
