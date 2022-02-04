package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_createSetterCmd = &cobra.Command{
	Use:   "create-setter",
	Short: "[Alpha] Create a custom setter for a Resource field",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_createSetterCmd).Standalone()
	cfg_createSetterCmd.Flags().String("description", "", "record a description for the current setter value.")
	cfg_createSetterCmd.Flags().String("field", "", "name of the field to set, a suffix of the path to the field, or the full path to the field. Default is to match all fields.")
	cfg_createSetterCmd.Flags().BoolP("recurse-subpackages", "R", false, "creates setter recursively in all the nested subpackages")
	cfg_createSetterCmd.Flags().Bool("required", false, "indicates that this setter must be set by package consumer before live apply/preview")
	cfg_createSetterCmd.Flags().String("schema-path", "", "openAPI schema file path for setter constraints -- file content e.g. {\"type\": \"string\", \"maxLength\": 15, \"enum\": [\"allowedValue1\", \"allowedValue2\"]}")
	cfg_createSetterCmd.Flags().String("set-by", "", "record who the field was default by.")
	cfg_createSetterCmd.Flags().String("type", "", "OpenAPI field type for the setter -- e.g. integer,boolean,string.")
	cfg_createSetterCmd.Flags().String("value", "", "optional flag, alternative to specifying the value as an argument. e.g. used to specify values that start with '-'")
	cfgCmd.AddCommand(cfg_createSetterCmd)

	carapace.Gen(cfg_createSetterCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("string", "number", "integer", "boolean", "array", "object"),
	})

	carapace.Gen(cfg_createSetterCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
