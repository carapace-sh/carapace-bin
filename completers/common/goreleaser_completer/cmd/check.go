package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:     "check [configuration files]",
	Short:   "Checks if configuration is valid",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringP("config", "f", "", "Configuration file(s) to check")
	checkCmd.Flags().Bool("deprecated", false, "Force print the deprecation message - tests only")
	checkCmd.Flags().BoolP("quiet", "q", false, "Quiet mode: no output")
	checkCmd.Flag("config").Hidden = true
	checkCmd.Flag("deprecated").Hidden = true
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
