package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmutil",
	Short: "Time Machine utility",
	Long:  "https://keith.github.io/xcode-manpages/tmutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"addexclusion", "Configure an exclusion from Time Machine backups",
			"associatedisk", "Bind a volume store directory to a local disk",
			"calculatedrift", "Analyze backups and determine change between them",
			"compare", "Perform a backup diff",
			"delete", "Delete backups with the specified timestamp",
			"deleteinprogress", "Delete all in-progress backups",
			"deletelocalsnapshots", "Delete local Time Machine snapshots",
			"destinationinfo", "Print information about configured destinations",
			"disable", "Turn off automatic backups",
			"enable", "Turn on automatic backups",
			"inheritbackup", "Claim a machine directory or sparsebundle",
			"isexcluded", "Determine if an item is excluded from backups",
			"latestbackup", "List this computer's latest completed backup",
			"listbackups", "List all of this computer's completed backups",
			"listlocalsnapshotdates", "List creation dates of local snapshots",
			"listlocalsnapshots", "List local Time Machine snapshots",
			"localsnapshot", "Create new local Time Machine snapshots",
			"machinedirectory", "Print the path to the current machine directory",
			"removeexclusion", "Configure Time Machine to back up an item",
			"removedestination", "Remove a destination from configuration",
			"restore", "Restore an item from a backup",
			"setdestination", "Configure a backup destination",
			"setquota", "Set the quota for a destination",
			"startbackup", "Begin a backup",
			"stopbackup", "Cancel a backup in progress",
			"thinlocalsnapshots", "Thin local Time Machine snapshots",
			"uniquesize", "Analyze a path's unique size in backups",
			"verifychecksums", "Verify checksums of data in a backup",
		),
	)
}