package action

import "github.com/carapace-sh/carapace"

func ActionSecretsProvider() carapace.Action {
	return carapace.ActionValues("default", "passphrase", "awskms", "azurekeyvault", "gcpkms", "hashivault")
}
