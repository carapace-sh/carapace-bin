package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List available Teleport clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clustersCmd).Standalone()

	clustersCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	clustersCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	clustersCmd.Flags().Bool("no-verbose", false, "Verbose table output, shows full label output.")
	clustersCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	clustersCmd.Flags().BoolP("verbose", "v", false, "Verbose table output, shows full label output.")
	clustersCmd.Flag("no-quiet").Hidden = true
	clustersCmd.Flag("no-verbose").Hidden = true
	rootCmd.AddCommand(clustersCmd)
}
