module github.com/carapace-sh/carapace-bin

go 1.26.2

require (
	github.com/carapace-sh/carapace v1.13.1
	github.com/carapace-sh/carapace-bridge v1.6.1
	github.com/carapace-sh/carapace-jjlex v0.1.8
	github.com/carapace-sh/carapace-jq v0.0.1
	github.com/carapace-sh/carapace-selfupdate v0.0.10
	github.com/carapace-sh/carapace-shlex v1.1.1
	github.com/carapace-sh/carapace-spec v1.8.0
	github.com/micromdm/plist v0.2.2
	github.com/pelletier/go-toml v1.9.5
	github.com/spf13/cobra v1.10.2
	github.com/spf13/pflag v1.0.10
	golang.org/x/mod v0.38.0
	gopkg.in/ini.v1 v1.67.3
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/pelletier/go-toml/v2 v2.4.2 // indirect

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kevinburke/ssh_config v1.4.0
)

replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag v1.1.0

replace github.com/kevinburke/ssh_config => github.com/carapace-sh/ssh_config v1.4.1-0.20260319075335-4f04016b8b4b
