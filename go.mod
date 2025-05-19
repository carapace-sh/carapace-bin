module github.com/carapace-sh/carapace-bin

go 1.23.1

require (
	github.com/carapace-sh/carapace v1.8.2
	github.com/carapace-sh/carapace-bridge v1.2.9
	github.com/carapace-sh/carapace-selfupdate v0.0.10
	github.com/carapace-sh/carapace-shlex v1.0.1
	github.com/carapace-sh/carapace-spec v1.1.2
	github.com/pelletier/go-toml v1.9.5
	github.com/spf13/cobra v1.9.1
	github.com/spf13/pflag v1.0.6
	golang.org/x/mod v0.24.0
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
)

replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag v1.0.0
