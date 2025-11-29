package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var global_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Interact with devbox services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_servicesCmd).Standalone()

	global_servicesCmd.PersistentFlags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_servicesCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	global_servicesCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	global_servicesCmd.PersistentFlags().String("environment", "", "environment to use, when supported (e.g. secrets support dev, prod, preview.)")
	global_servicesCmd.PersistentFlags().Bool("run-in-current-shell", false, "run the command in the current shell instead of a new shell")
	global_servicesCmd.Flag("run-in-current-shell").Hidden = true
	globalCmd.AddCommand(global_servicesCmd)

	// TODO environment
	carapace.Gen(global_servicesCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})
}
