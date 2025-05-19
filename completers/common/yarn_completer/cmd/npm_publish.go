package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the active workspace to the npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_publishCmd).Standalone()

	npm_publishCmd.Flags().String("access", "", "The access for the published package (public or restricted)")
	npm_publishCmd.Flags().String("otp", "", "The OTP token to use with the command")
	npm_publishCmd.Flags().String("tag", "", "The tag on the registry that the package should be attached to")
	npm_publishCmd.Flags().Bool("tolerate-republish", false, "Warn and exit when republishing an already existing version of a package")
	npmCmd.AddCommand(npm_publishCmd)

	carapace.Gen(npm_publishCmd).FlagCompletion(carapace.ActionMap{
		"access": carapace.ActionValues("public", "restricted"),
	})
}
