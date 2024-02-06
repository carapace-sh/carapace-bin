package cmd

import (
	"github.com/rsteube/carapace"
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

	auditCmd.Flags().Bool("arch", false, "Audit the given CPU architecture. (Pass `all` to audit all architectures.)")
	auditCmd.Flags().Bool("audit-debug", false, "Enable debugging and profiling of audit methods.")
	auditCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	auditCmd.Flags().Bool("debug", false, "Display any debugging information.")
	auditCmd.Flags().Bool("display-filename", false, "Prefix every line of output with the file or formula name being audited, to make output easy to grep.")
	auditCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to audit them. Implied if `HOMEBREW_EVAL_ALL` is set.")
	auditCmd.Flags().Bool("except", false, "Specify a comma-separated <method> list to skip running the methods named `audit_`<method>.")
	auditCmd.Flags().Bool("except-cops", false, "Specify a comma-separated <cops> list to skip checking for violations of the listed RuboCop cops.")
	auditCmd.Flags().Bool("fix", false, "Fix style violations automatically using RuboCop's auto-correct feature.")
	auditCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	auditCmd.Flags().Bool("git", false, "Run additional, slower style checks that navigate the Git repository.")
	auditCmd.Flags().Bool("help", false, "Show this message.")
	auditCmd.Flags().Bool("installed", false, "Only check formulae and casks that are currently installed.")
	auditCmd.Flags().Bool("new", false, "Run various additional style checks to determine if a new formula or cask is eligible for Homebrew. This should be used when creating new formulae or casks and implies `--strict` and `--online`.")
	auditCmd.Flags().Bool("no-signing", false, "Audit for signed apps, which are required on ARM")
	auditCmd.Flags().Bool("online", false, "Run additional, slower style checks that require a network connection.")
	auditCmd.Flags().Bool("only", false, "Specify a comma-separated <method> list to only run the methods named `audit_`<method>.")
	auditCmd.Flags().Bool("only-cops", false, "Specify a comma-separated <cops> list to check for violations of only the listed RuboCop cops.")
	auditCmd.Flags().Bool("os", false, "Audit the given operating system. (Pass `all` to audit all operating systems.)")
	auditCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	auditCmd.Flags().Bool("signing", false, "Audit for signed apps, which are required on ARM")
	auditCmd.Flags().Bool("skip-style", false, "Skip running non-RuboCop style checks. Useful if you plan on running `brew style` separately. Enabled by default unless a formula is specified by name.")
	auditCmd.Flags().Bool("strict", false, "Run additional, stricter style checks.")
	auditCmd.Flags().Bool("tap", false, "Check the formulae within the given tap, specified as <user>`/`<repo>.")
	auditCmd.Flags().Bool("token-conflicts", false, "Audit for token conflicts.")
	auditCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(auditCmd)
}
