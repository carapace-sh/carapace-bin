package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Check <formula> for Homebrew coding style violations",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().Bool("arch", false, "Audit the given CPU architecture")
	auditCmd.Flags().Bool("audit-debug", false, "Enable debugging and profiling of audit methods")
	auditCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	auditCmd.Flags().Bool("debug", false, "Display any debugging information")
	auditCmd.Flags().Bool("display-filename", false, "Prefix every line of output with the file or formula name being audited")
	auditCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks")
	auditCmd.Flags().String("except", "", "Specify a comma-separated <method> list to skip")
	auditCmd.Flags().String("except-cops", "", "Specify a comma-separated <cops> list to skip checking")
	auditCmd.Flags().Bool("fix", false, "Fix style violations automatically using RuboCop's auto-correct feature")
	auditCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	auditCmd.Flags().Bool("git", false, "Run additional, slower style checks that navigate the Git repository")
	auditCmd.Flags().Bool("help", false, "Show this message")
	auditCmd.Flags().Bool("installed", false, "Only check formulae and casks that are currently installed")
	auditCmd.Flags().Bool("new", false, "Run various additional style checks")
	auditCmd.Flags().Bool("no-signing", false, "Audit for signed apps, which are required on ARM")
	auditCmd.Flags().Bool("online", false, "Run additional, slower style checks that require a network connection")
	auditCmd.Flags().String("only", "", "Specify a comma-separated <method> list to only run")
	auditCmd.Flags().String("only-cops", "", "Specify a comma-separated <cops> list to check")
	auditCmd.Flags().String("os", "", "Audit the given operating system")
	auditCmd.Flags().Bool("quiet", false, "Make some output more quiet")
	auditCmd.Flags().Bool("signing", false, "Audit for signed apps, which are required on ARM")
	auditCmd.Flags().Bool("skip-style", false, "Skip running non-RuboCop style checks")
	auditCmd.Flags().Bool("strict", false, "Run additional, stricter style checks")
	auditCmd.Flags().String("tap", "", "Check the formulae within the given tap, specified as <user>`/`<repo>")
	auditCmd.Flags().Bool("token-conflicts", false, "Audit for token conflicts")
	auditCmd.Flags().Bool("verbose", false, "Make some output more verbose")
	rootCmd.AddCommand(auditCmd)

	// TODO flag completion
	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"tap": gh.ActionOwnerRepositories(gh.HostOpts{}),
	})

	carapace.Gen(auditCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := auditCmd.Flag("tap"); f.Changed {
				if owner, repo, ok := strings.Cut(f.Value.String(), "/"); ok {
					return gh.ActionContents(gh.ContentOpts{Owner: owner, Name: repo}) // TODO list remote formulae
				}
			}

			return carapace.Batch(
				brew.ActionAllCasks().Unless(auditCmd.Flag("formula").Changed),
				brew.ActionAllFormulae().Unless(auditCmd.Flag("cask").Changed),
			).ToA().FilterArgs()
		}),
	)
}
