package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	bat "github.com/rsteube/carapace-completers/completers/bat_completer/cmd"
	chmod "github.com/rsteube/carapace-completers/completers/chmod_completer/cmd"
	curl "github.com/rsteube/carapace-completers/completers/curl_completer/cmd"
	exa "github.com/rsteube/carapace-completers/completers/exa_completer/cmd"
	git "github.com/rsteube/carapace-completers/completers/git_completer/cmd"
	gradle "github.com/rsteube/carapace-completers/completers/gradle_completer/cmd"
	mkdir "github.com/rsteube/carapace-completers/completers/mkdir_completer/cmd"
	pkill "github.com/rsteube/carapace-completers/completers/pkill_completer/cmd"
	sha1sum "github.com/rsteube/carapace-completers/completers/sha1sum_completer/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:       "carapace-completers",
	Short:     "",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"bat", "chmod", "curl", "exa", "git", "gradle", "mkdir", "pkill", "sha1sum"},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO if len < thans sth (script generation)
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		if len(args) > 0 {
			switch args[0] {
			case "bat":
				os.Args[1] = "_carapace"
				bat.Execute()
			case "chmod":
				os.Args[1] = "_carapace"
				chmod.Execute()
			case "curl":
				os.Args[1] = "_carapace"
				curl.Execute()
			case "exa":
				os.Args[1] = "_carapace"
				exa.Execute()
			case "git":
				os.Args[1] = "_carapace"
				git.Execute()
			case "gradle":
				os.Args[1] = "_carapace"
				gradle.Execute()
			case "mkdir":
				os.Args[1] = "_carapace"
				mkdir.Execute()
			case "pkill":
				os.Args[1] = "_carapace"
				pkill.Execute()
			case "sha1sum":
				os.Args[1] = "_carapace"
				sha1sum.Execute()
			}
		}

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = old
		fmt.Println(strings.Replace(string(out), "carapace-completers _carapace", "carapace-completers "+args[0], -1))
	},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd)
}
