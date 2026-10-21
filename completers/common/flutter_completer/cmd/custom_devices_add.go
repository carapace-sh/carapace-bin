package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var custom_devices_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new device the custom devices config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(custom_devices_addCmd).Standalone()

	custom_devices_addCmd.Flags().Bool("check", false, "Make sure the config actually works. This will execute some of the commands in the config (if necessary with dummy arguments). This flag is enabled by default when \"--json\" is not specified. If \"--json\" is given, it is disabled by default. For example, a config with \"null\" as the \"runDebug\" command is invalid. If the \"runDebug\" command is valid (so it is an array of strings) but the command is not found (because you have a typo, for example), the config won't work and \"--check\" will spot that.")
	custom_devices_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	custom_devices_addCmd.Flags().String("js", "", "Add the custom device described by this JSON-encoded string to the list of custom-devices instead of using the normal, interactive way of configuring. Useful if you want to use the \"flutter custom-devices add\" command from a script, or use it non-interactively for some other reason. By default, this won't check whether the passed in config actually works. For more info see the \"--check\" option.")
	custom_devices_addCmd.Flags().String("json", "", "Add the custom device described by this JSON-encoded string to the list of custom-devices instead of using the normal, interactive way of configuring. Useful if you want to use the \"flutter custom-devices add\" command from a script, or use it non-interactively for some other reason. By default, this won't check whether the passed in config actually works. For more info see the \"--check\" option.")
	custom_devices_addCmd.Flags().Bool("no-check", false, "Make sure the config actually works. This will execute some of the commands in the config (if necessary with dummy arguments). This flag is enabled by default when \"--json\" is not specified. If \"--json\" is given, it is disabled by default. For example, a config with \"null\" as the \"runDebug\" command is invalid. If the \"runDebug\" command is valid (so it is an array of strings) but the command is not found (because you have a typo, for example), the config won't work and \"--check\" will spot that.")
	custom_devices_addCmd.Flags().Bool("ssh", false, "Add a ssh-device. This will automatically fill out some of the config options for you with good defaults, and in other cases save you some typing. So you'll only need to enter some things like hostname and username of the remote device instead of entering each individual command.")
	custom_devicesCmd.AddCommand(custom_devices_addCmd)
}
