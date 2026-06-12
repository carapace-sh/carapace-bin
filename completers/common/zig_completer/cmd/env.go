package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print lib path, std path, cache directory, and version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	rootCmd.AddCommand(envCmd)

	envCmd.Flags().String("color", "", "Enable or disable colored error messages")
	envCmd.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
	})
}
