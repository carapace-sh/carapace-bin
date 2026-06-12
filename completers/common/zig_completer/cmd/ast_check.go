package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var astCheckCmd = &cobra.Command{
	Use:     "ast-check [options] [files]",
	Short:   "Look for simple compile errors in any set of files",
	GroupID: "source",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(astCheckCmd).Standalone()

	rootCmd.AddCommand(astCheckCmd)

	astCheckCmd.Flags().String("color", "", "Enable or disable colored error messages")
	astCheckCmd.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(astCheckCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
	})

	carapace.Gen(astCheckCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig").FilterArgs(),
	)
}
