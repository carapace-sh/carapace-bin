package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List available Teleport clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clustersCmd).Standalone()

	clustersCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	clustersCmd.Flags().Bool("no-quiet", false, "Quiet mode")
	clustersCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	clustersCmd.Flag("no-quiet").Hidden = true
	rootCmd.AddCommand(clustersCmd)

	carapace.Gen(clustersCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
