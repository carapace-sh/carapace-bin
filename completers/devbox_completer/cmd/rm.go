package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/devbox"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a package from your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			dir := rmCmd.Flag("config").Value.String()
			return devbox.ActionInstalledPackages().Chdir(dir).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
