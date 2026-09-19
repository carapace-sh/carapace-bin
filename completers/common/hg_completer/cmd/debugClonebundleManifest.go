package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugClonebundleManifestCmd = &cobra.Command{
	Use:    "debug::clonebundle-manifest",
	Short:  "REPO",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugClonebundleManifestCmd).Standalone()

	debugClonebundleManifestCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	debugClonebundleManifestCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	debugClonebundleManifestCmd.Flags().Bool("raw", false, "show the raw manifest instead of parsing and filtering it")
	debugClonebundleManifestCmd.Flags().Bool("stream", false, "disply stream bundle only")
	rootCmd.AddCommand(debugClonebundleManifestCmd)

	carapace.Gen(debugClonebundleManifestCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})
}
