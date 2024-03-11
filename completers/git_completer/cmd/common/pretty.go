package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/text"
	"github.com/spf13/cobra"
)

func AddPrettyFlags(cmd *cobra.Command) {
	// TODO move stuff to revlist options
	// TODO check if pretty is stil relevant
	// TODO add missing flags fix completions
	cmd.Flags().Bool("no-notes", false, "do not show notes")
	cmd.Flags().String("notes", "", "show the notes that annotate the commit")
	cmd.Flags().Bool("show-notes-by-default", false, "show the default notes")

	cmd.Flag("notes").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"encoding": text.ActionEncodings(),
		"format":   carapace.ActionValues(), // TODO formats
		"notes":    carapace.ActionValues(), // TODO complete refs
		"pretty": carapace.ActionValues(
			"email",
			"format:",
			"full",
			"fuller",
			"medium",
			"oneline",
			"raw",
			"reference",
			"short",
			"tformat:",
		), // TODO format: and tformat:
	})
}
