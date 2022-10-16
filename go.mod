module github.com/rsteube/carapace-bin

go 1.18

require (
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/pelletier/go-toml v1.9.5
	github.com/rsteube/carapace v0.25.1
	github.com/rsteube/carapace-spec v0.1.0
	github.com/spf13/cobra v1.6.0
	github.com/spf13/pflag v1.0.5
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)

replace github.com/spf13/pflag => github.com/rsteube/carapace-pflag v0.0.4
