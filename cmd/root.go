package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:       "carapace",
	Short:     "",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: completers,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Args[1] == "--list" {
			fmt.Println(strings.Join(completers, "\n"))
			return
		}
		// TODO if len < thans sth (script generation)
		if len(args) > 0 {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			completer := args[0]
			os.Args[1] = "_carapace"
			executeCompleter(completer)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = old
			fmt.Print(strings.Replace(string(out), "carapace _carapace", "carapace "+completer, -1))
		}
	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.Flags().Bool("list", false, "list completers")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(completers...),
		carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"),
	)
}
