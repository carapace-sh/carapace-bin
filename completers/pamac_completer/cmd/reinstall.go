package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/pacman"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("as-deps", false, "mark all packages installed as a dependency")
	reinstallCmd.Flags().Bool("as-explicit", false, "mark all packages explicitly installed")
	reinstallCmd.Flags().BoolP("download-only", "w", false, "download all packages but do not install/upgrade")
	reinstallCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	reinstallCmd.Flags().String("overwrite", "", "overwrite conflicting files, multiple patterns can be")
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(carapace.ActionCallback(func(args []string) carapace.Action {
		return pacman.ActionPackages(pacman.PackageOption{Explicit: true}).Invoke(args).Filter(args).ToA() // TODO groups as well
	}))
}
