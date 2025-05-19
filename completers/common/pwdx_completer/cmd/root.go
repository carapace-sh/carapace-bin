package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwdx",
	Short: "report current working directory of a process",
	Long:  "https://man7.org/linux/man-pages/man1/pwdx.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
