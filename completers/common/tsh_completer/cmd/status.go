package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the list of proxy servers and retrieved certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("client", false, "Show client information only (no server required).")
	statusCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	statusCmd.Flags().Bool("no-client", false, "Show client information only (no server required).")
	statusCmd.Flags().Bool("no-verbose", false, "Show extra status information after successful login.")
	statusCmd.Flags().BoolP("verbose", "v", false, "Show extra status information after successful login.")
	statusCmd.Flag("no-client").Hidden = true
	statusCmd.Flag("no-verbose").Hidden = true
	rootCmd.AddCommand(statusCmd)
}
