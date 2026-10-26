package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dumpAutoloadCmd = &cobra.Command{
	Use:     "dump-autoload",
	Short:   "Dumps the autoloader",
	Aliases: []string{"dumpautoload"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpAutoloadCmd).Standalone()

	dumpAutoloadCmd.Flags().Bool("apcu", false, "Use APCu to cache found/not-found classes.")
	dumpAutoloadCmd.Flags().String("apcu-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu")
	dumpAutoloadCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize`.")
	dumpAutoloadCmd.Flags().Bool("dev", false, "Enables autoload-dev rules. Composer will by default infer this automatically according to the last install or update --no-dev state.")
	dumpAutoloadCmd.Flags().Bool("dry-run", false, "Outputs the operations but will not execute anything.")
	dumpAutoloadCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	dumpAutoloadCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	dumpAutoloadCmd.Flags().Bool("no-dev", false, "Disables autoload-dev rules. Composer will by default infer this automatically according to the last install or update --no-dev state.")
	dumpAutoloadCmd.Flags().BoolP("optimize", "o", false, "Optimizes PSR0 and PSR4 packages to be loaded with classmaps too, good for production.")
	dumpAutoloadCmd.Flags().Bool("strict-ambiguous", false, "Return a failed status code (2) if the same class is found in multiple files. Requires --optimize to work.")
	dumpAutoloadCmd.Flags().Bool("strict-psr", false, "Return a failed status code (1) if PSR-4 or PSR-0 mapping errors are present. Requires --optimize to work.")

	rootCmd.AddCommand(dumpAutoloadCmd)

	carapace.Gen(dumpAutoloadCmd).FlagCompletion(carapace.ActionMap{
		"ignore-platform-req": action.ActionPlatformPackages(dumpAutoloadCmd),
	})
}
