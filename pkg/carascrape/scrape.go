package carascrape

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var tmpDir string = getTmpDir()

func getTmpDir() string {
	dir, err := os.MkdirTemp(os.TempDir(), "carascrape")
	if err != nil {
		panic(err.Error())
	}
	return dir
}
func Scrape(cmd *cobra.Command) {
	out := &bytes.Buffer{}

	if len(cmd.Aliases) > 0 {
		fmt.Fprintf(out, `package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var %vCmd = &cobra.Command{
	Use:     "%v",
	Short:   "%v",
	Aliases: []string{"%v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

`, cmdVarName(cmd), cmd.Name(), cmd.Short, strings.Join(cmd.Aliases, `", "`))
	} else {

		fmt.Fprintf(out, `package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var %vCmd = &cobra.Command{
	Use:   "%v",
	Short: "%v",
	Run:   func(cmd *cobra.Command, args []string) {},
}

`, cmdVarName(cmd), cmd.Name(), cmd.Short)
	}
	if !cmd.HasParent() {
		fmt.Fprintf(out, `func Execute() error {
	return rootCmd.Execute()
}

`)
	}

	fmt.Fprintf(out, `func init() {
	carapace.Gen(%vCmd).Standalone()
`, cmdVarName(cmd))

	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if f.Deprecated != "" {
			return
		}

		isShortHand := f.Shorthand != ""

		persistentPrefix := ""
		if cmd.PersistentFlags().Lookup(f.Name) != nil {
			persistentPrefix = "Persistent"
		}

		if isShortHand {
			fmt.Fprintf(out, `	%vCmd.%vFlags().%vP("%v", "%v", %v, "%v")`+"\n", cmdVarName(cmd), persistentPrefix, flagType(f), f.Name, f.Shorthand, flagValue(f), formatUsage(f.Usage))
		} else {
			fmt.Fprintf(out, `	%vCmd.%vFlags().%v("%v", %v, "%v")`+"\n", cmdVarName(cmd), persistentPrefix, flagType(f), f.Name, flagValue(f), formatUsage(f.Usage))
		}
	})

	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if f.Deprecated != "" {
			return
		}

		if f.Value.Type() != "bool" && f.NoOptDefVal != "" {
			fmt.Fprintf(out, `    %vCmd.Flag("%v").NoOptDefVal = "%v"`+"\n", cmdVarName(cmd), f.Name, f.NoOptDefVal)
		}
	})

	if cmd.HasParent() {
		fmt.Fprintf(out, `	%vCmd.AddCommand(%vCmd)`+"\n", cmdVarName(cmd.Parent()), cmdVarName(cmd))
	}

	fmt.Fprintln(out, "}")

	filename := fmt.Sprintf(`%v/%v.go`, tmpDir, cmdVarName(cmd))
	println(filename)
	ioutil.WriteFile(filename, out.Bytes(), 0644)

	for _, subcmd := range cmd.Commands() {
		if !subcmd.Hidden && subcmd.Deprecated == "" {
			Scrape(subcmd)
		}
	}
}

func formatUsage(usage string) string {
	return strings.Replace(strings.Split(usage, "\n")[0], `"`, `\"`, -1)
}

func cmdVarName(cmd *cobra.Command) string {
	if !cmd.HasParent() {
		return "root"
	}
	return strings.TrimPrefix(fmt.Sprintf(`%v_%v`, cmdVarName(cmd.Parent()), normaliceVarName(cmd.Name())), "root_")
}

func normaliceVarName(s string) string {
	normalized := make([]string, 0)
	capitalize := false

	for _, c := range s {
		if c == '-' {
			capitalize = true
			continue
		} else if capitalize {
			normalized = append(normalized, strings.ToUpper(string(c)))
			capitalize = false
		} else {
			normalized = append(normalized, string(c))
		}
	}
	return strings.Join(normalized, "")
}

func flagType(f *pflag.Flag) string {
	return strings.ToUpper(f.Value.Type()[:1]) + f.Value.Type()[1:]
}

func flagValue(f *pflag.Flag) string {
	if strings.HasSuffix(f.Value.Type(), "Slice") ||
		strings.HasSuffix(f.Value.Type(), "Array") {
		if strings.HasPrefix(f.Value.Type(), "string") {
			vals, _ := csv.NewReader(strings.NewReader(f.Value.String()[1 : len(f.Value.String())-1])).Read()
			formatted := strings.Join(vals, `", "`)
			if len(formatted) > 0 {
				formatted = fmt.Sprintf(`"%v"`, formatted)
			}
			return fmt.Sprintf(`[]string{%v}`, formatted)
		}
		return fmt.Sprintf(`[]%v{%v}`, strings.TrimSuffix(strings.TrimSuffix(f.Value.Type(), "Slice"), "Array"), f.Value.String()[1:len(f.Value.String())-1])
	}

	if f.Value.Type() == "string" {
		return fmt.Sprintf(`"%v"`, f.Value.String())
	}

	return f.Value.String()
}
