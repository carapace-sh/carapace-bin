// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_taskStubsCmd = &cobra.Command{
	Use:   "task-stubs",
	Short: "Generate scripts that run mise tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_taskStubsCmd).Standalone()

	generate_taskStubsCmd.Flags().StringP("dir", "d", "bin", "Directory to create task stubs inside of")
	generate_taskStubsCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_taskStubsCmd.Flags().StringP("mise-bin", "m", "mise", "mise executable the stubs run; defaults to mise from PATH")
	generate_taskStubsCmd.Flags().String("windows-launcher", "cmd", "What to write beside each stub for Windows to launch")
	generateCmd.AddCommand(generate_taskStubsCmd)

	carapace.Gen(generate_taskStubsCmd).FlagCompletion(carapace.ActionMap{
		"dir":              carapace.ActionDirectories(),
		"windows-launcher": carapace.ActionValues("cmd", "exe"),
	})
}
