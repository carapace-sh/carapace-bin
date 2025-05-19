package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new profile",
	Aliases: []string{"configure"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_addCmd).Standalone()
	profile_addCmd.Flags().BoolP("acceptDefaults", "y", false, "suppress prompts and accept default values")
	profile_addCmd.Flags().Int("accountId", 0, "your account ID")
	profile_addCmd.Flags().String("apiKey", "", "your personal API key")
	profile_addCmd.Flags().String("licenseKey", "", "your license key")
	profile_addCmd.Flags().StringP("region", "r", "", "the US or EU region")
	profileCmd.AddCommand(profile_addCmd)

	carapace.Gen(profile_addCmd).FlagCompletion(carapace.ActionMap{
		"region": carapace.ActionValues("US", "EU"),
	})
}
