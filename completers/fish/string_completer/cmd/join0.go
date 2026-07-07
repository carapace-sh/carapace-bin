package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var join0Cmd = &cobra.Command{
	Use:   "join0",
	Short: "Join strings with NUL bytes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join0Cmd).Standalone()

	join0Cmd.Flags().BoolP("no-empty", "n", false, "exclude empty strings")
	join0Cmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(join0Cmd)
}
