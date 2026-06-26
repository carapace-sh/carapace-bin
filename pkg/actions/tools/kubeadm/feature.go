package kubeadm

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFeatureGates completes feature gates
//
//	NodeLocalCRISocket=true
//	PublicKeysECDSA=false
//	RootlessControlPlane=false
func ActionFeatureGates() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"NodeLocalCRISocket", "default=true",
				"PublicKeysECDSA", "DEPRECATED - default=false",
				"RootlessControlPlane", "ALPHA - default=false",
			).Suffix("=")
		default:
			return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		}
	}).List(",")
}
