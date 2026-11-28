package action

import (
	"github.com/carapace-sh/carapace"
)

// ActionEmbedFormats completes output formats of `llm embed`
func ActionEmbedFormats() carapace.Action {
	return carapace.ActionValues("base64", "blob", "hex", "json").Tag("formats")
}

// ActionEmbedMultiFormats completes output formats of `llm embed-multi`
func ActionEmbedMultiFormats() carapace.Action {
	return carapace.ActionValues("csv", "json", "nl", "tsv").Tag("formats")
}
