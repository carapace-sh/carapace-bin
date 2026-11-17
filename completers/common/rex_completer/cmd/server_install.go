package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverInstallCmd = &cobra.Command{
	Use:   "install [flags] [ssh://[USER@]HOST[:PORT]] [-- SERVER-FLAG...]",
	Short: "Start your server at boot and keep it running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverInstallCmd).Standalone()

	serverInstallCmd.Flags().SetInterspersed(false)

	serverInstallCmd.Flags().String("channel", "", "install from the release channel NAME")
	serverInstallCmd.Flags().Bool("json", false, "output what the command did as JSON")
	serverInstallCmd.Flags().String("label", "", "add the machine to your hosts as LABEL")
	serverInstallCmd.Flags().Bool("package", false, "install the Rex package for the whole machine through sudo")
	serverInstallCmd.Flags().Bool("user", false, "start the server when you sign in and stop it when you sign out")
	serverInstallCmd.Flags().Bool("yes", false, "answer yes to every question")

	serverCmd.AddCommand(serverInstallCmd)

	carapace.Gen(serverInstallCmd).FlagCompletion(carapace.ActionMap{
		"channel": carapace.ActionValues("preview", "snapshot", "stable"),
	})
}
