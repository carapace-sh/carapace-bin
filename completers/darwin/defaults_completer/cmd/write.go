package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/defaults"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write the value for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(writeCmd).Standalone()
	writeCmd.Flags().Bool("array", false, "Write an array")
	writeCmd.Flags().Bool("array-add", false, "Append to an array")
	writeCmd.Flags().Bool("bool", false, "Write a boolean value")
	writeCmd.Flags().Bool("data", false, "Write raw data (hex)")
	writeCmd.Flags().Bool("date", false, "Write a date")
	writeCmd.Flags().Bool("float", false, "Write a float value")
	writeCmd.Flags().Bool("int", false, "Write an integer value")
	writeCmd.Flags().Bool("string", false, "Write a string value")
	rootCmd.AddCommand(writeCmd)
	carapace.Gen(writeCmd).PositionalCompletion(
		defaults.ActionDomains(),
	)
}
