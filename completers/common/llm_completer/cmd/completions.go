package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/llm_completer/cmd/action"
)

func init() {
	applyFlagCompletion(promptCmd, carapace.ActionMap{
		"attachment": carapace.ActionFiles(),
		"database":   carapace.ActionFiles(),
		"fragment":   carapace.ActionFiles(),
		"functions":  carapace.ActionFiles(),
		"schema":     carapace.ActionFiles(".json"),
	})

	applyFlagCompletion(chatCmd, carapace.ActionMap{
		"attachment": carapace.ActionFiles(),
		"database":   carapace.ActionFiles(),
		"fragment":   carapace.ActionFiles(),
		"functions":  carapace.ActionFiles(),
		"schema":     carapace.ActionFiles(".json"),
	})

	applyFlagCompletion(similarCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(embedCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
		"format":   action.ActionEmbedFormats(),
		"input":    carapace.ActionFiles(),
	})

	applyFlagCompletion(embedMultiCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
		"format":   action.ActionEmbedMultiFormats(),
	})

	applyFlagCompletion(collectionsCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(collections_listCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(collections_deleteCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(logsCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(logs_listCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(logs_backupCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(keysCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(aliasesCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(modelsCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(embedModelsCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(fragmentsCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})

	applyFlagCompletion(templatesCmd, carapace.ActionMap{
		"database": carapace.ActionFiles(),
	})
}
