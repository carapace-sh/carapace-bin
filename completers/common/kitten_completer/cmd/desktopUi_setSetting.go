package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setSettingCmd = &cobra.Command{
	Use:   "set-setting",
	Short: "Change an arbitrary setting",
}

func init() {
	setSettingCmd.AddCommand(desktopUiCmd)
	carapace.Gen(setSettingCmd).Standalone()

	setSettingCmd.Flags().String("data-type", "", "The DBUS data type signature of the value")
	setSettingCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setSettingCmd.Flags().StringP("namespace", "n", "org.freedesktop.appearance", "The namespace in which to change the setting")

	carapace.Gen(setSettingCmd).FlagCompletion(carapace.ActionMap{})
}
