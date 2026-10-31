package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update wrap files from WrapDB (Since 0.63.0)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_updateCmd).Standalone()

	wrap_updateCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	wrap_updateCmd.Flags().Bool("force", false, "Update wraps that does not seems to come from WrapDB")
	wrap_updateCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	wrap_updateCmd.Flags().String("sourcedir", "", "Path to source directory")
	wrap_updateCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(wrap_updateCmd)
	wrapCmd.AddCommand(wrap_updateCmd)
}
