package env

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/aws"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["aws"] = func() variables {
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		return variables{
			Condition: conditions.ConditionPath("aws"),
			Variables: map[string]string{
				"AWS_ACCESS_KEY_ID":         "Specifies an AWS access key associated with an IAM account",
				"AWS_CA_BUNDLE":             "Specifies the path to a certificate bundle to use for HTTPS certificate validation",
				"AWS_CLI_AUTO_PROMPT":       "Enables the auto-prompt for the AWS CLI version 2",
				"AWS_CLI_FILE_ENCODING":     "Specifies the encoding used for text files",
				"AWS_CONFIG_FILE":           "Specifies the location of the file that the AWS CLI uses to store configuration profiles",
				"AWS_DATA_PATH":             "A list of additional directories to check outside of the built-in search path of ~/.aws/models",
				"AWS_DEFAULT_OUTPUT":        "Specifies the output format to use",
				"AWS_DEFAULT_REGION":        "The Default region name",
				"AWS_EC2_METADATA_DISABLED": "Disables the use of the Amazon EC2 instance metadata service (IMDS)",
				// "AWS_ENDPOINT_URL_<SERVICE>": "pecifies a custom endpoint that is used for a specific service, where <SERVICE> is replaced with the AWS service identifier. For example, Amazon DynamoDB has a serviceId of DynamoDB. For this service, the endpoint URL environment variable is AWS_ENDPOINT_URL_DYNAMODB.
				"AWS_ENDPOINT_URL":                    "Specifies the endpoint that is used for all service requests",
				"AWS_IGNORE_CONFIGURED_ENDPOINT_URLS": "If enabled, the AWS CLI ignores all custom endpoint configurations",
				"AWS_MAX_ATTEMPTS":                    "Specifies a value of maximum retry attempts the AWS CLI retry handler uses",
				"AWS_METADATA_SERVICE_NUM_ATTEMPTS":   "retry multiple times before giving up",
				"AWS_METADATA_SERVICE_TIMEOUT":        "The number of seconds before a connection to the instance metadata service should time out",
				"AWS_PAGER":                           "Specifies the pager program used for output",
				"AWS_PROFILE":                         "Specifies the name of the AWS CLI profile with the credentials and options to use",
				"AWS_REGION":                          "The AWS SDK compatible environment variable that specifies the AWS Region to send the request to",
				"AWS_RETRY_MODE":                      "Specifies which retry mode AWS CLI uses",
				"AWS_ROLE_ARN":                        "Specifies the Amazon Resource Name (ARN) of an IAM role",
				"AWS_ROLE_SESSION_NAME":               "Specifies the name to attach to the role session",
				"AWS_SECRET_ACCESS_KEY":               "Specifies the secret key associated with the access key",
				"AWS_SHARED_CREDENTIALS_FILE":         "Specifies the location of the file that the AWS CLI uses to store access keys",
				"AWS_USE_FIPS_ENDPOINT":               "Federal Information Processing Standard (FIPS) endoint",
				"AWS_WEB_IDENTITY_TOKEN_FILE":         "Specifies the path to a file that contains an OAuth 2.0 access token",
			},
			VariableCompletion: map[string]carapace.Action{
				"AWS_ACCESS_KEY_ID": carapace.ActionValues(),
				"AWS_CA_BUNDLE":     carapace.ActionFiles(),
				"AWS_CLI_AUTO_PROMPT": carapace.ActionValuesDescribed(
					"on", "full auto-prompt mode each time you attempt to run an aws command",
					"on-partial", "partial auto-prompt mode",
				).StyleF(style.ForKeyword),
				"AWS_CLI_FILE_ENCODING":     carapace.ActionValues(), // TODO encodings
				"AWS_CONFIG_FILE":           carapace.ActionFiles(),
				"AWS_DATA_PATH":             carapace.ActionDirectories().List(string(os.PathListSeparator)),
				"AWS_DEFAULT_OUTPUT":        aws.ActionOutputFormats(),
				"AWS_DEFAULT_REGION":        aws.ActionRegions(),
				"AWS_EC2_METADATA_DISABLED": _bool,
				"AWS_ENDPOINT_URL":          carapace.ActionValues(),
				// "AWS_ENDPOINT_URL_<SERVICE>": carapace.ActionValues(), // TODO
				"AWS_IGNORE_CONFIGURED_ENDPOINT_URLS": _bool,
				"AWS_MAX_ATTEMPTS":                    carapace.ActionValues(),
				"AWS_METADATA_SERVICE_NUM_ATTEMPTS":   carapace.ActionValues(),
				"AWS_METADATA_SERVICE_TIMEOUT":        carapace.ActionValues(),
				"AWS_PAGER":                           bridge.ActionCarapaceBin().Split(),
				"AWS_PROFILE":                         aws.ActionProfiles(),
				"AWS_REGION":                          aws.ActionRegions(),
				"AWS_RETRY_MODE":                      carapace.ActionValues("legacy", "standard", "adaptive"),
				"AWS_ROLE_ARN":                        carapace.ActionValues(),
				"AWS_ROLE_SESSION_NAME":               carapace.ActionValues(),
				"AWS_SECRET_ACCESS_KEY":               carapace.ActionValues(),
				"AWS_SHARED_CREDENTIALS_FILE":         carapace.ActionFiles(),
				"AWS_USE_FIPS_ENDPOINT":               carapace.ActionValues(),
				"AWS_WEB_IDENTITY_TOKEN_FILE":         carapace.ActionFiles(),
			},
		}
	}
}
