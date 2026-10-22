package cmd

import (
	"os"

	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
)

func init() {
	action.SetGlobalCloud(globalCloud)
	action.SetGlobalEnv(globalEnv)
}

func globalCloud() string {
	if flag := rootCmd.Flags().Lookup("os-cloud"); flag != nil && flag.Changed {
		return flag.Value.String()
	}
	return os.Getenv("OS_CLOUD")
}

func globalEnv() []string {
	env := []string{}
	addVar := func(name, flagName string) {
		if flag := rootCmd.Flags().Lookup(flagName); flag != nil && flag.Changed {
			env = append(env, name+"="+flag.Value.String())
		}
	}
	addBool := func(name, flagName string) {
		if flag := rootCmd.Flags().Lookup(flagName); flag != nil && flag.Changed {
			env = append(env, name+"="+flag.Value.String())
		}
	}

	addVar("OS_CLOUD", "os-cloud")
	addVar("OS_REGION_NAME", "os-region-name")
	addVar("OS_INTERFACE", "os-interface")
	addVar("OS_AUTH_TYPE", "os-auth-type")
	addVar("OS_AUTH_URL", "os-auth-url")
	addVar("OS_IDENTITY_API_VERSION", "os-identity-api-version")
	addVar("OS_USERNAME", "os-username")
	addVar("OS_USER_ID", "os-user-id")
	addVar("OS_PASSWORD", "os-password")
	addVar("OS_TOKEN", "os-token")
	addVar("OS_PROJECT_NAME", "os-project-name")
	addVar("OS_PROJECT_ID", "os-project-id")
	addVar("OS_PROJECT_DOMAIN_NAME", "os-project-domain-name")
	addVar("OS_PROJECT_DOMAIN_ID", "os-project-domain-id")
	addVar("OS_USER_DOMAIN_NAME", "os-user-domain-name")
	addVar("OS_USER_DOMAIN_ID", "os-user-domain-id")
	addVar("OS_DOMAIN_NAME", "os-domain-name")
	addVar("OS_DOMAIN_ID", "os-domain-id")
	addVar("OS_DEFAULT_DOMAIN", "os-default-domain")
	addVar("OS_DEFAULT_DOMAIN_NAME", "os-default-domain-name")
	addVar("OS_DEFAULT_DOMAIN_ID", "os-default-domain-id")
	addVar("OS_TENANT_NAME", "os-tenant-name")
	addVar("OS_TENANT_ID", "os-tenant-id")
	addVar("OS_TRUST_ID", "os-trust-id")
	addVar("OS_APPLICATION_CREDENTIAL_ID", "os-application-credential-id")
	addVar("OS_APPLICATION_CREDENTIAL_NAME", "os-application-credential-name")
	addVar("OS_APPLICATION_CREDENTIAL_SECRET", "os-application-credential-secret")
	addVar("OS_ACCESS_TOKEN", "os-access-token")
	addVar("OS_SYSTEM_SCOPE", "os-system-scope")
	addVar("OS_CACERT", "os-cacert")
	addVar("OS_CERT", "os-cert")
	addVar("OS_KEY", "os-key")
	addVar("OS_ENDPOINT_OVERRIDE", "os-endpoint-override")
	addVar("OS_SERVICE_PROVIDER", "os-service-provider")
	addVar("OS_REMOTE_PROJECT_ID", "os-remote-project-id")
	addVar("OS_REMOTE_PROJECT_NAME", "os-remote-project-name")
	addVar("OS_REMOTE_PROJECT_DOMAIN_ID", "os-remote-project-domain-id")
	addVar("OS_REMOTE_PROJECT_DOMAIN_NAME", "os-remote-project-domain-name")
	addVar("OS_COMPUTE_API_VERSION", "os-compute-api-version")
	addVar("OS_NETWORK_API_VERSION", "os-network-api-version")
	addVar("OS_VOLUME_API_VERSION", "os-volume-api-version")
	addVar("OS_IMAGE_API_VERSION", "os-image-api-version")
	addVar("OS_OBJECT_API_VERSION", "os-object-api-version")
	addVar("OS_SHARE_API_VERSION", "os-share-api-version")
	addVar("OS_ACCELERATOR_API_VERSION", "os-accelerator-api-version")
	addBool("OS_INSECURE", "insecure")
	addBool("OS_VERIFY", "verify")

	return env
}
