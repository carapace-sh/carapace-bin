package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(localsnapshotCmd)
	rootCmd.AddCommand(listlocalsnapshotsCmd)
	rootCmd.AddCommand(listlocalsnapshotdatesCmd)
	rootCmd.AddCommand(deletelocalsnapshotsCmd)
	rootCmd.AddCommand(thinlocalsnapshotsCmd)
}

var localsnapshotCmd = &cobra.Command{
	Use:   "localsnapshot",
	Short: "create new local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listlocalsnapshotsCmd = &cobra.Command{
	Use:   "listlocalsnapshots",
	Short: "list local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listlocalsnapshotdatesCmd = &cobra.Command{
	Use:   "listlocalsnapshotdates",
	Short: "list creation dates of local snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deletelocalsnapshotsCmd = &cobra.Command{
	Use:   "deletelocalsnapshots",
	Short: "delete local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var thinlocalsnapshotsCmd = &cobra.Command{
	Use:   "thinlocalsnapshots",
	Short: "thin local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localsnapshotCmd).Standalone()
	carapace.Gen(listlocalsnapshotsCmd).Standalone()
	carapace.Gen(listlocalsnapshotdatesCmd).Standalone()
	carapace.Gen(deletelocalsnapshotsCmd).Standalone()
	carapace.Gen(thinlocalsnapshotsCmd).Standalone()
}