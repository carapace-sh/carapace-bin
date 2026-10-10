package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugmanifestfulltextcacheCmd = &cobra.Command{
	Use:    "debugmanifestfulltextcache",
	Short:  "show, clear or amend the contents of the manifest fulltext cache",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugmanifestfulltextcacheCmd).Standalone()

	debugmanifestfulltextcacheCmd.Flags().StringArrayP("add", "a", nil, "add the given manifest nodes to the cache")
	debugmanifestfulltextcacheCmd.Flags().Bool("clear", false, "clear the cache")
	rootCmd.AddCommand(debugmanifestfulltextcacheCmd)
}
