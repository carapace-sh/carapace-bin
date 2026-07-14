package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var cvsimportCmd = &cobra.Command{
	Use:     "cvsimport",
	Short:   "Salvage your data out of another SCM people love to hate",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(cvsimportCmd).Standalone()

	cvsimportCmd.Flags().StringS("A", "A", "", "CVS by default uses the Unix username when writing its commit logs")
	cvsimportCmd.Flags().StringS("C", "C", "", "the Git repository to import to")
	cvsimportCmd.Flags().StringS("L", "L", "", "limit the number of commits imported")
	cvsimportCmd.Flags().StringS("M", "M", "", "attempt to detect merges based on the commit message with a custom regex")
	cvsimportCmd.Flags().StringS("P", "P", "", "instead of calling cvsps, read the provided cvsps output file")
	cvsimportCmd.Flags().BoolS("R", "R", false, "generate a $GIT_DIR/cvs-revisions file containing a mapping from CVS revision numbers to newly-created Git commit IDs")
	cvsimportCmd.Flags().StringS("S", "S", "", "skip paths matching the regex")
	cvsimportCmd.Flags().BoolS("a", "a", false, "import all commits, including recent ones")
	cvsimportCmd.Flags().StringS("d", "d", "", "the root of the CVS archive")
	cvsimportCmd.Flags().BoolS("h", "h", false, "print a short usage message and exit")
	cvsimportCmd.Flags().BoolS("i", "i", false, "import-only: don’t perform a checkout after importing")
	cvsimportCmd.Flags().BoolS("k", "k", false, "kill keywords: will extract files with -kk from the CVS archive to avoid noisy changesets")
	cvsimportCmd.Flags().BoolS("m", "m", false, "attempt to detect merges based on the commit message")
	cvsimportCmd.Flags().StringS("o", "o", "", "import into a different branch")
	cvsimportCmd.Flags().StringS("p", "p", "", "additional options for cvsps")
	cvsimportCmd.Flags().StringS("r", "r", "", "the Git remote to import this CVS repository into")
	cvsimportCmd.Flags().StringS("s", "s", "", "substitute the character \"/\" in branch names with <subst>")
	cvsimportCmd.Flags().BoolS("u", "u", false, "convert underscores in tag and branch names to dots")
	cvsimportCmd.Flags().BoolS("v", "v", false, "verbosity: let cvsimport report what it is doing")
	cvsimportCmd.Flags().StringS("z", "z", "", "pass the timestamp fuzz factor to cvsps, in seconds")
	rootCmd.AddCommand(cvsimportCmd)

	carapace.Gen(cvsimportCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionFiles(),
		"C": carapace.ActionDirectories(),
		"P": carapace.ActionFiles(),
		"d": carapace.ActionDirectories(),
		"o": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}).ChdirF(traverse.Flag(cvsimportCmd.Flag("C"))),
		"p": carapace.ActionValues(), // TODO cvsps options
		"r": git.ActionRemotes().ChdirF(traverse.Flag(cvsimportCmd.Flag("C"))),
	})
}
