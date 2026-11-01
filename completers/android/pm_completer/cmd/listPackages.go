package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var listPackagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Print packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPackagesCmd).Standalone()

	listPackagesCmd.Flags().BoolP("all", "a", false, "all known packages (but excluding APEXes)")
	listPackagesCmd.Flags().Bool("apex-only", false, "only show APEX packages")
	listPackagesCmd.Flags().BoolP("disabled", "d", false, "filter to only show disabled packages")
	listPackagesCmd.Flags().BoolP("enabled", "e", false, "filter to only show enabled packages")
	listPackagesCmd.Flags().Bool("factory-only", false, "only show system packages excluding updates")
	listPackagesCmd.Flags().BoolP("file", "f", false, "see their associated file")
	listPackagesCmd.Flags().BoolP("ignored", "l", false, "ignored (used for compatibility with older releases)")
	listPackagesCmd.Flags().BoolP("installer", "i", false, "see the installer for the packages")
	listPackagesCmd.Flags().Bool("match-libraries", false, "include packages that declare static shared and SDK libraries")
	listPackagesCmd.Flags().BoolP("quarantined", "q", false, "filter to only show quarantined packages")
	listPackagesCmd.Flags().BoolP("show-uid", "U", false, "also show the package UID")
	listPackagesCmd.Flags().Bool("show-versioncode", false, "also show the version code")
	listPackagesCmd.Flags().BoolP("system", "s", false, "filter to only show system packages")
	listPackagesCmd.Flags().BoolP("third-party", "3", false, "filter to only show third party packages")
	listPackagesCmd.Flags().String("uid", "", "filter to only show packages with the given `UID`")
	listPackagesCmd.Flags().BoolP("uninstalled", "u", false, "also include uninstalled packages")
	listPackagesCmd.Flags().String("user", "", "only list packages belonging to the given `USER_ID`")

	listCmd.AddCommand(listPackagesCmd)

	carapace.Gen(listPackagesCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(listPackagesCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
