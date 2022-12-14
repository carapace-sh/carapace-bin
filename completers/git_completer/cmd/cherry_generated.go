package cmd

import (
	"github.com/spf13/cobra"
)

var cherryCmd = &cobra.Command{
	Use:     "cherry",
	Short:   "Find commits yet to be applied to upstream",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	cherryCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	cherryCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(cherryCmd)
}
