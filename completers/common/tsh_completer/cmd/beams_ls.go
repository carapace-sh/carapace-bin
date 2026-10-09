package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List beam instances.",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_lsCmd).Standalone()

	beams_lsCmd.Flags().Bool("all", false, "List all beams. By default, filters to show only beams belonging to the current user.")
	beams_lsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	beams_lsCmd.Flags().Bool("no-all", false, "List all beams. By default, filters to show only beams belonging to the current user.")
	beams_lsCmd.Flag("no-all").Hidden = true
	beamsCmd.AddCommand(beams_lsCmd)
}
