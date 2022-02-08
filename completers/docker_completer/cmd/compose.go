package cmd

import (
	compose "github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd"
)

func init() {
	// TODO verify this has no unwanted side-effects
	// TODO probably best to finally add a way to invoke and reuse a completer like `repo clone -- <git clone completion>`
	compose.AddTo(rootCmd)
}
