package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print commands to set Teleport session environment variables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	envCmd.Flags().Bool("no-unset", false, "Print commands to clear Teleport session environment variables")
	envCmd.Flags().Bool("unset", false, "Print commands to clear Teleport session environment variables")
	envCmd.Flag("no-unset").Hidden = true
	rootCmd.AddCommand(envCmd)

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
