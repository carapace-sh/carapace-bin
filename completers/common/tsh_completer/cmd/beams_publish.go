package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish an HTTP or TCP service running in a beam.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_publishCmd).Standalone()

	beams_publishCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	beams_publishCmd.Flags().Bool("no-tcp", false, "Publish as a TCP app instead of an HTTP app.")
	beams_publishCmd.Flags().Bool("tcp", false, "Publish as a TCP app instead of an HTTP app.")
	beams_publishCmd.Flag("no-tcp").Hidden = true
	beamsCmd.AddCommand(beams_publishCmd)
}
