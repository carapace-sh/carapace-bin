package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Interact with devbox services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicesCmd).Standalone()

	servicesCmd.PersistentFlags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	servicesCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	servicesCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	servicesCmd.PersistentFlags().String("environment", "", "environment to use, when supported (e.g. secrets support dev, prod, preview.)")
	servicesCmd.PersistentFlags().Bool("run-in-current-shell", false, "run the command in the current shell instead of a new shell")
	servicesCmd.Flag("run-in-current-shell").Hidden = true
	rootCmd.AddCommand(servicesCmd)

	// TODO environment
	carapace.Gen(servicesCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})
}
