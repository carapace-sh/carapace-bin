package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Start a new beam, and optionally connect to it via SSH.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_addCmd).Standalone()

	beams_addCmd.Flags().Bool("console", true, "Connect to the beam via SSH after creation.")
	beams_addCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	beams_addCmd.Flags().Bool("no-console", false, "Connect to the beam via SSH after creation.")
	beams_addCmd.Flag("no-console").Hidden = true
	beamsCmd.AddCommand(beams_addCmd)
}
