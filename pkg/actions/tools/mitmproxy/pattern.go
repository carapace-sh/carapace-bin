package mitmproxy

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
)

// ActionModifyBodyPattern completes body modification patterns
//
//	_~all_@file
//	_~s_.*_@file
func ActionModifyBodyPattern() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) > 1 {
			return carapace.ActionMultiParts(c.Value[:1], func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 1:
					return actionOptionFlowFilters().NoSpace()
				case 2:
					return actionOptionFiles().NoSpace()

				case 3:
					return actionOptionFiles()
				}
				return carapace.ActionValues()
			})
		}
		return carapace.ActionValues()
	})
}

// ActionModifyHeaderPattern completes header modification patterns
//
//	_Accept-Language_av,af,am_@file
//	_Accept-Encoding_zstd_@dist/
func ActionModifyHeaderPattern() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) > 1 {
			return carapace.ActionMultiParts(c.Value[:1], func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 1:
					return actionOptionFlowFilters().NoSpace()
				case 2:
					return http.ActionRequestHeaderNames().NoSpace()
				case 3:
					return carapace.Batch(
						http.ActionRequestHeaderValues(c.Parts[2]),
						actionOptionFiles(),
					).ToA()
				}
				return carapace.ActionValues()
			})
		}
		return carapace.ActionValues()
	})
}
