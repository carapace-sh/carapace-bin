package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_device_profile_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Register a new device profile with the accelerator service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_device_profile_createCmd).Standalone()

	accelerator_device_profile_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_device_profile_createCmd.Flags().String("description", "", "Description for the device profile")
	accelerator_device_profile_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_device_profile_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_device_profile_createCmd.Flags().String("group", "", "Device profile group (repeat option to specify multiple groups).")
	accelerator_device_profile_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_device_profile_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_device_profile_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_device_profile_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_device_profile_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_device_profileCmd.AddCommand(accelerator_device_profile_createCmd)
}
