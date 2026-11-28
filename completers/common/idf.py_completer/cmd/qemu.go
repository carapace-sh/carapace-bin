package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qemuCmd = &cobra.Command{
	Use:   "qemu",
	Short: "Run QEMU",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qemuCmd).Standalone()
	qemuCmd.Flags().String("efuse-file", "", "File used to store efuse values")
	qemuCmd.Flags().String("flash-file", "", "File used as the qemu flash image")
	qemuCmd.Flags().BoolP("gdb", "d", false, "Wait for gdb to connect")
	qemuCmd.Flags().BoolP("graphics", "g", false, "Enable graphical window")
	qemuCmd.Flags().String("qemu-extra-args", "", "Extra arguments to pass to QEMU")
	rootCmd.AddCommand(qemuCmd)
	carapace.Gen(qemuCmd).FlagCompletion(carapace.ActionMap{
		"efuse-file": carapace.ActionFiles(),
		"flash-file": carapace.ActionFiles(),
	})
}
