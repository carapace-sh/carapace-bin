package kubeadm

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFeatureGates completes feature gates
//
//	ControlPlaneKubeletLocalMode=true
//	NodeLocalCRISocket=false
func ActionFeatureGates() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValues(
				"ControlPlaneKubeletLocalMode",
				"NodeLocalCRISocket",
				"PublicKeysECDSA",
				"RootlessControlPlane",
				"WaitForAllControlPlaneComponents",
			).Suffix("=")
		default:
			return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		}
	}).List(",")
}
