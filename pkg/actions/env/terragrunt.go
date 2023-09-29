package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
	knownVariables["terragrunt"] = variables{
		Condition: conditions.ConditionPath("terragrunt"),
		Variables: map[string]string{
			"TERRAGRUNT_CONFIG":                             "The path to the Terragrunt config file. Default is terragrunt.hcl.",
			"TERRAGRUNT_TFPATH":                             "Path to the Terraform binary. Default is terraform (on PATH).",
			"TERRAGRUNT_NO_AUTO_INIT":                       "Don't automatically run 'terraform init' during other terragrunt commands. You must run 'terragrunt init' manually.",
			"TERRAGRUNT_NO_AUTO_RETRY":                      "Don't automatically re-run command in case of transient errors.",
			"TERRAGRUNT_NO_AUTO_APPROVE":                    "Don't automatically append `-auto-approve` to the underlying Terraform commands run with 'run-all'.",
			"TERRAGRUNT_NON_INTERACTIVE":                    `Assume "yes" for all prompts.`,
			"TERRAGRUNT_WORKING_DIR":                        "The path to the Terraform templates. Default is current directory.",
			"TERRAGRUNT_DOWNLOAD":                           "The path where to download Terraform code. Default is .terragrunt-cache in the working directory.",
			"TERRAGRUNT_SOURCE":                             "Download Terraform configurations from the specified source into a temporary folder, and run Terraform in that temporary folder.",
			"TERRAGRUNT_SOURCE_UPDATE":                      "Delete the contents of the temporary folder to clear out any old, cached source code before downloading new source code into it.",
			"TERRAGRUNT_SOURCE_MAP":                         "Replace any source URL (including the source URL of a config pulled in with dependency blocks) that has root source with dest.",
			"TERRAGRUNT_IAM_ROLE":                           "Assume the specified IAM role before executing Terraform. Can also be set via the TERRAGRUNT_IAM_ROLE environment variable.",
			"TERRAGRUNT_IAM_ASSUME_ROLE_DURATION":           "Session duration for IAM Assume Role session. Can also be set via the TERRAGRUNT_IAM_ASSUME_ROLE_DURATION environment variable.",
			"TERRAGRUNT_IAM_ASSUME_ROLE_SESSION_NAME":       "Name for the IAM Assummed Role session. Can also be set via TERRAGRUNT_IAM_ASSUME_ROLE_SESSION_NAME environment variable.",
			"TERRAGRUNT_INCLUDE_EXTERNAL_DEPENDENCIES":      "*-all commands will include external dependencies",
			"TERRAGRUNT_PARALLELISM":                        "*-all commands parallelism set to at most N modules",
			"TERRAGRUNT_EXCLUDE_DIR":                        "Unix-style glob of directories to exclude when running *-all commands.",
			"TERRAGRUNT_DEBUG":                              "Write terragrunt-debug.tfvars to working folder to help root-cause issues.",
			"TERRAGRUNT_LOG_LEVEL":                          "Sets the logging level for Terragrunt. Supported levels: panic, fatal, error, warn, info, debug, trace.",
			"TERRAGRUNT_NO_COLOR":                           "If specified, Terragrunt output won't contain any color.",
			"TERRAGRUNT_USE_PARTIAL_PARSE_CONFIG_CACHE":     "Enables caching of includes during partial parsing operations. Will also be used for the terragrunt-iam-role option if provided.",
			"TERRAGRUNT_FETCH_DEPENDENCY_OUTPUT_FROM_STATE": "The option fetchs dependency output directly from the state file instead of init dependencies and running terraform on them.",
			"TERRAGRUNT_INCLUDE_MODULE_PREFIX":              "When this flag is set output from Terraform sub-commands is prefixed with module path.",
			"TERRAGRUNT_FAIL_ON_STATE_BUCKET_CREATION":      "When this flag is set Terragrunt will fail if the remote state bucket needs to be created.",
			"TERRAGRUNT_DISABLE_BUCKET_UPDATE":              "When this flag is set Terragrunt will not update the remote state bucket.",
			"TERRAGRUNT_DISABLE_COMMAND_VALIDATION":         "When this flag is set, Terragrunt will not validate the terraform command.",
		},
		VariableCompletion: map[string]carapace.Action{
			"TERRAGRUNT_CONFIG":          carapace.ActionFiles(),
			"TERRAGRUNT_TFPATH":          carapace.ActionFiles(),
			"TERRAGRUNT_NO_AUTO_INIT":    _bool,
			"TERRAGRUNT_NO_AUTO_RETRY":   _bool,
			"TERRAGRUNT_NO_AUTO_APPROVE": _bool,
			"TERRAGRUNT_NON_INTERACTIVE": _bool,
			"TERRAGRUNT_WORKING_DIR":     carapace.ActionDirectories(),
			"TERRAGRUNT_DOWNLOAD":        carapace.ActionDirectories(),
			// "TERRAGRUNT_SOURCE":                             carapace.ActionValues(), // TODO
			"TERRAGRUNT_SOURCE_UPDATE":                      _bool,
			"TERRAGRUNT_SOURCE_MAP":                         carapace.ActionValues(),
			"TERRAGRUNT_IAM_ROLE":                           carapace.ActionValues(),
			"TERRAGRUNT_IAM_ASSUME_ROLE_DURATION":           carapace.ActionValues(),
			"TERRAGRUNT_IAM_ASSUME_ROLE_SESSION_NAME":       carapace.ActionValues(),
			"TERRAGRUNT_INCLUDE_EXTERNAL_DEPENDENCIES":      _bool,
			"TERRAGRUNT_PARALLELISM":                        carapace.ActionValues(),
			"TERRAGRUNT_EXCLUDE_DIR":                        carapace.ActionDirectories(),
			"TERRAGRUNT_DEBUG":                              _bool,
			"TERRAGRUNT_LOG_LEVEL":                          carapace.ActionValues("panic", "fatal", "error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
			"TERRAGRUNT_NO_COLOR":                           _bool,
			"TERRAGRUNT_USE_PARTIAL_PARSE_CONFIG_CACHE":     _bool,
			"TERRAGRUNT_FETCH_DEPENDENCY_OUTPUT_FROM_STATE": _bool,
			"TERRAGRUNT_INCLUDE_MODULE_PREFIX":              _bool,
			"TERRAGRUNT_FAIL_ON_STATE_BUCKET_CREATION":      _bool,
			"TERRAGRUNT_DISABLE_BUCKET_UPDATE":              _bool,
			"TERRAGRUNT_DISABLE_COMMAND_VALIDATION":         _bool,
		},
	}
}
