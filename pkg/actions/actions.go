package actions

import (
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/carapace-spec/pkg/macro"
)

var (
	MacroMap          = make(macro.MacroMap[spec.Macro])
	MacroDescriptions = make(map[string]string)
)
