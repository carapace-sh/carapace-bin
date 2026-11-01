package action

import "github.com/carapace-sh/carapace"

func ActionSecretsProvider() carapace.Action {
	return carapace.ActionValues("awskms", "azurekeyvault", "default", "gcpkms", "hashivault", "passphrase")
}
