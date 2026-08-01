package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:     "import [-DfFlmNns] [-c cachefile] [-d dir|device]... [-o property=value]... [-R root] [-T txg] [-t] [pool|id] [newpool]",
	Short:   "import a pool",
	GroupID: "io",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().BoolS("D", "D", false, "import destroyed pools")
	importCmd.Flags().BoolS("F", "F", false, "recovery mode")
	importCmd.Flags().BoolS("N", "N", false, "do not mount filesystems")
	importCmd.Flags().StringS("R", "R", "", "alternate root directory")
	importCmd.Flags().StringS("T", "T", "", "transaction group for rollback")
	importCmd.Flags().BoolS("X", "X", false, "extreme recovery")
	importCmd.Flags().BoolS("a", "a", false, "import all pools")
	importCmd.Flags().StringS("c", "c", "", "cachefile")
	importCmd.Flags().StringArrayS("d", "d", nil, "search directory or device")
	importCmd.Flags().BoolS("f", "f", false, "force import")
	importCmd.Flags().BoolS("l", "l", false, "load encryption keys")
	importCmd.Flags().BoolS("m", "m", false, "allow missing log device")
	importCmd.Flags().BoolS("n", "n", false, "dry-run with -F")
	importCmd.Flags().StringArrayS("o", "o", nil, "set mount option or pool property")
	importCmd.Flags().Bool("rewind-to-checkpoint", false, "rewind pool to the checkpointed state")
	importCmd.Flags().BoolS("s", "s", false, "scan default paths")
	importCmd.Flags().BoolS("t", "t", false, "temporary pool name")

	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionDirectories(),
		"c": carapace.ActionFiles(),
		"d": carapace.Batch(
			carapace.ActionDirectories(),
			carapace.ActionFiles(),
		).ToA(),
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(importCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
