package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_setSettingCmd = &cobra.Command{
	Use:   "set-setting",
	Short: "Change an arbitrary setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_setSettingCmd).Standalone()

	desktopUi_setSettingCmd.Flags().String("data-type", "", "The DBUS data type signature of the value")
	desktopUi_setSettingCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUi_setSettingCmd.Flags().StringP("namespace", "n", "org.freedesktop.appearance", "The namespace in which to change the setting")
	desktopUiCmd.AddCommand(desktopUi_setSettingCmd)

}
