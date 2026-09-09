package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_type_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume type properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_type_setCmd).Standalone()

	volume_type_setCmd.Flags().String("availability-zone", "", "Set an availability zone for this volume type (this is an alias for '--property RESKEY:availability_zones:<az>') (repeat option to set multiple availability zones)")
	volume_type_setCmd.Flags().Bool("cacheable", false, "Enable caching for this volume type (this is an alias for '--property cacheable=<is> True') (requires driver support)")
	volume_type_setCmd.Flags().String("description", "", "Set volume type description")
	volume_type_setCmd.Flags().String("encryption-cipher", "", "Set the encryption algorithm or mode for this volume type (e.g \"aes-xts-plain64\") (admin only)")
	volume_type_setCmd.Flags().String("encryption-control-location", "", "Set the notional service where the encryption is performed (\"front-end\" or \"back-end\") (admin only) (The default value for this option is \"front-end\" when setting encryption type of a volume for the first time.")
	volume_type_setCmd.Flags().String("encryption-key-size", "", "Set the size of the encryption key of this volume type (e.g \"128\" or \"256\") (admin only)")
	volume_type_setCmd.Flags().String("encryption-provider", "", "Set the encryption provider format for this volume type (e.g \"luks\" or \"plain\") (admin only) (This option is required when setting encryption type of a volume for the first time.")
	volume_type_setCmd.Flags().Bool("multiattach", false, "Enable multi-attach for this volume type (this is an alias for '--property multiattach=<is> True') (requires driver support)")
	volume_type_setCmd.Flags().String("name", "", "Set volume type name")
	volume_type_setCmd.Flags().Bool("private", false, "Volume type is not accessible to the public")
	volume_type_setCmd.Flags().String("project", "", "Set volume type access to project (name or ID) (admin only)")
	volume_type_setCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	volume_type_setCmd.Flags().String("property", "", "Set a property on this volume type (repeat option to set multiple properties)")
	volume_type_setCmd.Flags().Bool("public", false, "Volume type is accessible to the public")
	volume_type_setCmd.Flags().Bool("replicated", false, "Enabled replication for this volume type (this is an alias for '--property replication_enabled=<is> True') (requires driver support)")
	volume_typeCmd.AddCommand(volume_type_setCmd)
}
