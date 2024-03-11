package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the list of proxy servers and retrieved certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	statusCmd.Flags().Bool("no-verbose", false, "Show extra status information after successful login")
	statusCmd.Flags().BoolP("verbose", "v", false, "Show extra status information after successful login")
	statusCmd.Flag("no-verbose").Hidden = true
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
