package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_scpCmd = &cobra.Command{
	Use:     "scp",
	Short:   "Copy files between a beam and the local filesystem.",
	Aliases: []string{"cp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_scpCmd).Standalone()

	beams_scpCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	beams_scpCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	beams_scpCmd.Flags().Bool("no-recursive", false, "Recursive copy of subdirectories.")
	beams_scpCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	beams_scpCmd.Flags().BoolP("recursive", "r", false, "Recursive copy of subdirectories.")
	beams_scpCmd.Flag("no-quiet").Hidden = true
	beams_scpCmd.Flag("no-recursive").Hidden = true
	beamsCmd.AddCommand(beams_scpCmd)
}
