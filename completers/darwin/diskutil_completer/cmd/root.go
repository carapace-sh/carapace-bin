package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diskutil",
	Short: "modify, verify and repair local disks",
	Long:  "https://keith.github.io/xcode-manpages/diskutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("acls", "a", false, "Enable ACL support on the volume")
	rootCmd.Flags().BoolP("force", "f", false, "Force operation")
	rootCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")
	rootCmd.Flags().BoolP("plist", "p", false, "Output in plist format")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"list", "List partitions on all disks or a specified disk",
			"info", "Get information on a specific disk or partition",
			"infoFilesystem", "Get filesystem information on a specific volume",
			"verifyDisk", "Verify the partition map of a disk",
			"repairDisk", "Repair the partition map of a disk",
			"verifyVolume", "Verify the filesystem of a volume",
			"repairVolume", "Repair the filesystem of a volume",
			"eraseDisk", "Erase and format a disk",
			"eraseVolume", "Erase and format a volume",
			"eraseFreespace", "Erase the free space on a volume",
			"reformat", "Reformat an existing volume",
			"resizeVolume", "Resize a volume",
			"mergePartitions", "Merge two or more partitions into one",
			"splitPartition", "Split an existing partition",
			"addPartition", "Add a new partition",
			"enableOwnership", "Enable ownership for a volume",
			"disableOwnership", "Disable ownership for a volume",
			"enableJournaling", "Enable journaling on an HFS+ volume",
			"disableJournaling", "Disable journaling on an HFS+ volume",
			"mount", "Mount a volume",
			"unmount", "Unmount a volume",
			"unmountDisk", "Unmount all volumes on a disk",
			"eject", "Eject a disk",
			"secureErase", "Securely erase a disk or volume",
			"apfs", "APFS operations",
			"appleRAID", "Apple RAID operations",
			"coreStorage", "CoreStorage operations",
			"imageVolume", "Create and attach a disk image to a volume",
			"partitionDisk", "Partition a disk",
		),
		fs.ActionBlockDevices(),
	)
}
