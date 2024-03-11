package vagrant

import "github.com/carapace-sh/carapace"

// ActionProviders completes providers
func ActionProviders() carapace.Action {
	return carapace.ActionValues(
		"aws",
		"cloudstack",
		"digitalocean",
		"docker",
		"google",
		"hyperv",
		"libvirt",
		"lxc",
		"openstack",
		"parallels",
		"qemu",
		"rackspace",
		"softlayer",
		"veertu",
		"virtualbox",
		"vmware",
		"vmware_desktop",
		"vmware_fusion",
		"vmware_ovf",
		"vmware_workstation",
		"vsphere",
		"xenserver",
	)
}
