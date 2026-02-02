package git

import (
	"github.com/carapace-sh/carapace"
)

// ActionObjectFilters completes object filters
//
//	blob:none (omits all blobs)
//	blob:limit (omits blobs of size at least <n> bytes or units)
func ActionObjectFilters() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValuesDescribed(
					"blob:none", "omits all blobs",
				).Uid("git", "object-filter"),
				carapace.ActionValuesDescribed(
					"tree", "omits all blobs and trees whose depth from the root tree is >= <depth>",
				).Uid("git", "object-filter").Suffix(":").NoSpace(':'),
				carapace.ActionValuesDescribed(
					"blob:limit", "omits blobs of size at least <n> bytes or units",
					"object:type", "omits all objects which are not of the requested type",
					"sparse:oid", "omit blobs that would not be required for a sparse checkout",
				).Uid("git", "object-filter").Suffix("="),
			).ToA()
		default:
			switch c.Parts[0] {
			case "object:type":
				return carapace.ActionValues("tag", "commit", "tree", "blob")
			default:
				return carapace.ActionValues()
			}
		}
	}).Tag("object filters")
}
