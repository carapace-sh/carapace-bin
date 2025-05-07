package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golangcilint"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func ActionLinters(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return golangcilint.ActionLinters(cmd.Flag("config").Value.String())
	})
}

func ActionOutputPaths() carapace.Action {
	return carapace.Batch(
		carapace.ActionStyledValues("stdout", style.Carapace.KeywordPositive),
		carapace.ActionStyledValues("stderr", style.Carapace.KeywordNegative),
		carapace.ActionFiles(),
	).ToA()
}
