package git

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
)

// ActionTrailers completes trailer key:value pairs
//
//	Co-authored-by: (author)
//	Signed-off-by: (author)
func ActionTrailers() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValues(
				"Co-authored-by",
				"Signed-off-by",
				"Helped-by",
			).Suffix(":")
		default:
			return carapace.Batch(
				gh.ActionOwners(gh.HostOpts{}), // TODO include email
				ActionAuthors(),
			).ToA()
		}
	}).Tag("trailers")
}
