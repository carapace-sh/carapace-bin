package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Installs a package and any packages that it depends on",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	addCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	addCmd.Flags().Bool("color", false, "Controls colors in the output")
	addCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	addCmd.Flags().String("filter", "", "set filter")
	addCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	addCmd.Flags().BoolP("global", "g", false, "Install as a global package")
	addCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	addCmd.Flags().BoolP("help", "h", false, "Output usage information")
	addCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts")
	addCmd.Flags().String("loglevel", "", "What level of logs to report")
	addCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	addCmd.Flags().Bool("no-save-exact", false, "Do not install exact version")
	addCmd.Flags().Bool("no-save-workspace-protocol", false, "Do not save packages from the workspace with a \"workspace:\" protocol")
	addCmd.Flags().Bool("offline", false, "Trigger an error if any required dependencies are not available")
	addCmd.Flags().Bool("prefer-offline", false, "Skip staleness checks for cached data")
	addCmd.Flags().BoolP("recursive", "r", false, "Run installation recursively in every package found in subdirectories")
	addCmd.Flags().BoolP("save-dev", "D", false, "Save package to your `devDependencies`")
	addCmd.Flags().BoolP("save-exact", "E", false, "Install exact version")
	addCmd.Flags().BoolP("save-optional", "O", false, "Save package to your `optionalDependencies`")
	addCmd.Flags().Bool("save-peer", false, "Save package to your `peerDependencies` and `devDependencies`")
	addCmd.Flags().BoolP("save-prod", "P", false, "Save package to your `dependencies`")
	addCmd.Flags().Bool("save-workspace-protocol", false, "Save packages from the workspace with a \"workspace:\" protocol")
	addCmd.Flags().Bool("silent", false, "turn off all logging")
	addCmd.Flags().String("store-dir", "", "The directory in which all the packages are saved on the disk")
	addCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	addCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	addCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	addCmd.Flags().String("virtual-store-dir", "", "The directory with links to the store")
	addCmd.Flags().Bool("workspace", false, "Only adds the new dependency if it is found in the workspace")
	addCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"dir":               carapace.ActionDirectories(),
		"filter":            pnpm.ActionFilter(),
		"loglevel":          pnpm.ActionLoglevel(),
		"store-dir":         carapace.ActionDirectories(),
		"virtual-store-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "https://") {
				return git.ActionRepositorySearch(git.SearchOpts{}.Default())
			}

			return carapace.Batch(
				carapace.ActionFiles(),
				npm.ActionPackageSearch("").UnlessF(condition.CompletingPath),
				git.ActionRepositorySearch(git.SearchOpts{}.Default()).UnlessF(condition.CompletingPath),
			).ToA()
		}),
	)
}
