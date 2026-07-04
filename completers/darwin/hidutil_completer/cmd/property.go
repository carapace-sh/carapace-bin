package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var propertyCmd = &cobra.Command{
	Use:   "property",
	Short: "Read/write HID Event System property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(propertyCmd).Standalone()

	propertyCmd.Flags().StringP("get", "g", "", "Get property with key")
	propertyCmd.Flags().StringP("matching", "m", "", "Set matching services/devices (JSON dictionary or device string)")
	propertyCmd.Flags().StringP("set", "s", "", "Set property (key/value pair JSON dictionary)")
	rootCmd.AddCommand(propertyCmd)
}
