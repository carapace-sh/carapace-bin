package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var buildUpdateRepoCmd = &cobra.Command{
	Use:     "build-update-repo [OPTION…] LOCATION",
	Short:   "Update repository metadata",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildUpdateRepoCmd).Standalone()

	buildUpdateRepoCmd.Flags().Bool("authenticator-install", false, "Autoinstall authenticator for this repository")
	buildUpdateRepoCmd.Flags().String("authenticator-name", "", "Name of authenticator for this repository")
	buildUpdateRepoCmd.Flags().String("authenticator-option", "", "Authenticator option")
	buildUpdateRepoCmd.Flags().String("collection-id", "", "Collection ID")
	buildUpdateRepoCmd.Flags().String("comment", "", "A one-line comment for this repository")
	buildUpdateRepoCmd.Flags().String("default-branch", "", "Default branch to use for this repository")
	buildUpdateRepoCmd.Flags().Bool("deploy-collection-id", false, "Permanently deploy collection ID to client remote configurations")
	buildUpdateRepoCmd.Flags().Bool("deploy-sideload-collection-id", false, "Permanently deploy collection ID to client remote configurations, only for sideload support")
	buildUpdateRepoCmd.Flags().String("description", "", "A full-paragraph description for this repository")
	buildUpdateRepoCmd.Flags().Bool("generate-static-deltas", false, "Generate delta files")
	buildUpdateRepoCmd.Flags().String("gpg-homedir", "", "GPG Homedir to use when looking for keyrings")
	buildUpdateRepoCmd.Flags().String("gpg-import", "", "Import new default GPG public key from FILE")
	buildUpdateRepoCmd.Flags().String("gpg-sign", "", "GPG Key ID to sign the summary with")
	buildUpdateRepoCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildUpdateRepoCmd.Flags().String("homepage", "", "URL for a website for this repository")
	buildUpdateRepoCmd.Flags().String("icon", "", "URL for an icon for this repository")
	buildUpdateRepoCmd.Flags().Bool("no-authenticator-install", false, "Don't autoinstall authenticator for this repository")
	buildUpdateRepoCmd.Flags().Bool("no-summary-index", false, "Don't generate a summary index")
	buildUpdateRepoCmd.Flags().Bool("no-update-appstream", false, "Don't update the appstream branch")
	buildUpdateRepoCmd.Flags().Bool("no-update-summary", false, "Don't update the summary")
	buildUpdateRepoCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildUpdateRepoCmd.Flags().Bool("prune", false, "Prune unused objects")
	buildUpdateRepoCmd.Flags().String("prune-depth", "", "Only traverse DEPTH parents for each commit (default: -1=infinite)")
	buildUpdateRepoCmd.Flags().Bool("prune-dry-run", false, "Prune but don't actually remove anything")
	buildUpdateRepoCmd.Flags().String("redirect-url", "", "Redirect this repo to a new URL")
	buildUpdateRepoCmd.Flags().String("static-delta-ignore-ref", "", "Don't create deltas matching refs")
	buildUpdateRepoCmd.Flags().String("static-delta-jobs", "", "Max parallel jobs when creating deltas (default: NUMCPUs)")
	buildUpdateRepoCmd.Flags().String("title", "", "A nice name to use for this repository")
	buildUpdateRepoCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildUpdateRepoCmd)

	// flag
	carapace.Gen(buildUpdateRepoCmd).FlagCompletion(carapace.ActionMap{
		// "authenticator-name":      carapace.ActionValues(),
		// "authenticator-option":    carapace.ActionValues(),
		// "collection-id":           carapace.ActionValues(),
		// "comment":                 carapace.ActionValues(),
		// "default-branch":          carapace.ActionValues(),
		// "description":             carapace.ActionValues(),
		"gpg-homedir": carapace.ActionDirectories(),
		"gpg-import":  carapace.ActionFiles(),
		"gpg-sign":    os.ActionGpgKeyIds(), // TODO handle gpg-homedir
		// "homepage":                carapace.ActionValues(),
		// "icon":                    carapace.ActionValues(),
		// "prune-depth":             carapace.ActionValues(),
		// "redirect-url":            carapace.ActionValues(),
		// "static-delta-ignore-ref": carapace.ActionValues(),
		// "static-delta-jobs":       carapace.ActionValues(),
		// "title":                   carapace.ActionValues(),
	})

	// TODO positional
}
