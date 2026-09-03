package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Show information about an operation and its view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_objectCmd).Standalone()

	debug_objectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_objectCmd)
}

var debug_object_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_commitCmd).Standalone()

	debug_object_commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_objectCmd.AddCommand(debug_object_commitCmd)
}

var debug_object_fileCmd = &cobra.Command{
	Use:   "file",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_fileCmd).Standalone()

	debug_object_fileCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_fileCmd.Flags().StringP("revision", "r", "", "")
	debug_objectCmd.AddCommand(debug_object_fileCmd)

	carapace.Gen(debug_object_fileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}

var debug_object_operationCmd = &cobra.Command{
	Use:   "operation",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_operationCmd).Standalone()

	debug_object_operationCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_objectCmd.AddCommand(debug_object_operationCmd)
}

var debug_object_symlinkCmd = &cobra.Command{
	Use:   "symlink",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_symlinkCmd).Standalone()

	debug_object_symlinkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_symlinkCmd.Flags().StringP("revision", "r", "", "")
	debug_objectCmd.AddCommand(debug_object_symlinkCmd)

	carapace.Gen(debug_object_symlinkCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}

var debug_object_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_treeCmd).Standalone()

	debug_object_treeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_treeCmd.Flags().StringP("revision", "r", "", "")
	debug_objectCmd.AddCommand(debug_object_treeCmd)

	carapace.Gen(debug_object_treeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}

var debug_object_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_viewCmd).Standalone()

	debug_object_viewCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_viewCmd.Flags().String("op", "", "")
	debug_objectCmd.AddCommand(debug_object_viewCmd)
}