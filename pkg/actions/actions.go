package actions

import (
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace-spec/pkg/macro"
)

var (
	MacroMap          = make(macro.MacroMap[spec.Macro])
	MacroDescriptions = make(map[string]string)
)
