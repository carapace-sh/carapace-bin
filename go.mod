module github.com/carapace-sh/carapace-bin

go 1.23.1

require (
	github.com/carapace-sh/carapace v1.10.0
	github.com/carapace-sh/carapace-bridge v1.4.8
	github.com/carapace-sh/carapace-selfupdate v0.0.10
	github.com/carapace-sh/carapace-shlex v1.1.0
	github.com/carapace-sh/carapace-spec v1.3.4
	github.com/pelletier/go-toml v1.9.5
	github.com/spf13/cobra v1.9.1
	github.com/spf13/pflag v1.0.9
	golang.org/x/mod v0.27.0
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kevinburke/ssh_config v1.4.0
	github.com/stretchr/testify v1.8.1 // indirect
)

replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag v1.0.0

replace github.com/kevinburke/ssh_config => github.com/carapace-sh/ssh_config v1.4.1-0.20251016142129-177d27a2d08a
