package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umask",
	Short: "Display or set file mode mask",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-umask",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "makes the output symbolic; otherwise an octal number is output")
	rootCmd.Flags().BoolS("p", "p", false, "output in a form that may be reused as input")

	carapace.Gen(rootCmd).PositionalCompletion(
		fs.ActionFileModes(),
	)
}
