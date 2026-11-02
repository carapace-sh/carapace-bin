package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
)

func ActionExtensions() carapace.Action {
	return gh.ActionExtensions()
}

func ActionTopExtensions() carapace.Action {
	return gh.ActionTopExtensions(gh.HostOpts{})
}

func Extensions() ([]string, error) {
	return gh.Extensions()
}
