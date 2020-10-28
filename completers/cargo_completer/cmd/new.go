package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new cargo package at <path>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	newCmd.Flags().Bool("bin", false, "Use a binary (application) template [default]")
	newCmd.Flags().String("color", "", "Coloring: auto, always, never")
	newCmd.Flags().String("edition", "", "Edition to set for the crate generated [possible values: 2015, 2018]")
	newCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	newCmd.Flags().BoolP("help", "h", false, "Prints help information")
	newCmd.Flags().Bool("lib", false, "Use a library template")
	newCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	newCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	newCmd.Flags().Bool("offline", false, "Run without accessing the network")
	newCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	newCmd.Flags().String("registry", "", "Registry to use")
	newCmd.Flags().String("vcs", "", "Initialize a new repository for the given version control system (git, hg, pijul, or")
	newCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
