package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_type_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new volume type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_type_createCmd).Standalone()

	volume_type_createCmd.Flags().String("availability-zone", "", "Set an availability zone for this volume type (this is an alias for '--property RESKEY:availability_zones:<az>') (repeat option to set multiple availability zones)")
	volume_type_createCmd.Flags().Bool("cacheable", false, "Enable caching for this volume type (this is an alias for '--property cacheable=<is> True') (requires driver support)")
	volume_type_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_type_createCmd.Flags().String("description", "", "Volume type description")
	volume_type_createCmd.Flags().String("encryption-cipher", "", "Set the encryption algorithm or mode for this volume type (e.g \"aes-xts-plain64\") (admin only)")
	volume_type_createCmd.Flags().String("encryption-control-location", "", "Set the notional service where the encryption is performed (\"front-end\" or \"back-end\") (admin only) (The default value for this option is \"front-end\" when setting encryption type of a volume.")
	volume_type_createCmd.Flags().String("encryption-key-size", "", "Set the size of the encryption key of this volume type (e.g \"128\" or \"256\") (admin only)")
	volume_type_createCmd.Flags().String("encryption-provider", "", "Set the encryption provider format for this volume type (e.g \"luks\" or \"plain\") (admin only) (this option is required when setting encryption type of a volume; consider using other encryption options such as: \"--encryption-cipher\", \"--encryption-key-size\" and \"--encryption-control-location\")")
	volume_type_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_type_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_type_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_type_createCmd.Flags().Bool("multiattach", false, "Enable multi-attach for this volume type (this is an alias for '--property multiattach=<is> True') (requires driver support)")
	volume_type_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_type_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_type_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_type_createCmd.Flags().Bool("private", false, "Volume type is not accessible to the public")
	volume_type_createCmd.Flags().String("project", "", "Allow <project> to access private type (name or ID) (must be used with --private option)")
	volume_type_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	volume_type_createCmd.Flags().String("property", "", "Set a property on this volume type (repeat option to set multiple properties)")
	volume_type_createCmd.Flags().Bool("public", false, "Volume type is accessible to the public")
	volume_type_createCmd.Flags().Bool("replicated", false, "Enabled replication for this volume type (this is an alias for '--property replication_enabled=<is> True') (requires driver support)")
	volume_type_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_typeCmd.AddCommand(volume_type_createCmd)
}
