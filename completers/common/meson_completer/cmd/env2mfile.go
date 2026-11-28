package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var env2mfileCmd = &cobra.Command{
	Use:   "env2mfile",
	Short: "Convert current environment to a cross or native file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env2mfileCmd).Standalone()

	env2mfileCmd.Flags().Bool("android", false, "Generate cross files for Android toolchains.")
	env2mfileCmd.Flags().String("cpu", "", "Define cpu for cross compilation.")
	env2mfileCmd.Flags().String("cpu-family", "", "Define cpu family for cross compilation.")
	env2mfileCmd.Flags().Bool("cross", false, "Generate a cross compilation file.")
	env2mfileCmd.Flags().String("debarch", "", "The dpkg architecture to generate.")
	env2mfileCmd.Flags().String("endian", "", "Define endianness for cross compilation.")
	env2mfileCmd.Flags().String("gccsuffix", "", "A particular gcc version suffix if necessary.")
	env2mfileCmd.Flags().String("kernel", "", "Define kernel for cross compilation.")
	env2mfileCmd.Flags().Bool("native", false, "Generate a native compilation file.")
	env2mfileCmd.Flags().StringS("o", "o", "", "The output file or directory (for Android).")
	env2mfileCmd.Flags().String("subsystem", "", "Define subsystem for cross compilation.")
	env2mfileCmd.Flags().String("system", "", "Define system for cross compilation.")
	env2mfileCmd.Flags().Bool("use-for-build", false, "Use _FOR_BUILD envvars.")
	applyCommonFlagCompletion(env2mfileCmd)
	rootCmd.AddCommand(env2mfileCmd)
}
