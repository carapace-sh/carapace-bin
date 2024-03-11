# Bridge

Bridging completions from another engine for shells not natively supported by it is possible with the corresponding [Macro](./macros.md).

> Even when the command supports your current shell it is still beneficial to bridge it as this enables [embedding](./embed.md) like `sudo [spec.name] <TAB>`.
>
> It also avoids the issue of [shell startup delay] when sourcing the completion in init scripts otherwise circumvented with [lazycomplete].
>
> However, bridging is limited to supported commands/frameworks and how well it actually works.


## Frameworks

### Argcomplete

[kislyuk/argcomplete] based commands can be bridged with the [`bridge.Argcomplete`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: az
description: Azure Command-Line Interface
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Argcomplete([az])"]
```

### Carapace

[carapace-sh/carapace] based commands can be bridged with the [`bridge.Carapace`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: freckles
description: simple dotfile manager
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Carapace([freckles])"]
```

### CarapaceBin

[Completers](../completers.md) and [Specs](../spec.md) registered in [carapace-sh/carapace-bin] can be bridged with the [`bridge.CarapaceBin`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: github-cli
description: Work seamlessly with GitHub from the command line
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.CarapaceBin([gh])"]
```

### Clap

[clap-rs/clap] based commands can be bridged with the [`bridge.Clap`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: dyamic
description: clap dynamic example
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Clap([dynamic])"]
```

### Click

[pallets/click] based commands can be bridged with the [`bridge.Click`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: watson
description: Watson is a tool aimed at helping you monitoring your time
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Click([watson])"]
```

### Cobra

[spf13/cobra] based commands can be bridged with the [`bridge.Cobra`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: kubectl
description: kubectl controls the Kubernetes cluster manager
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Cobra([kubectl])"]
```

### Complete
[posener/complete] based commands can be bridged with the [`bridge.Complete`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: vault
description: Manage Secrets & Protect Sensitive Data
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Complete([vault])"]
```

### Inshellisense
Commands provided by [microsoft/inshellisense] can be bridged with the [`bridge.Inshellisense`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: node
description: Run the node interpreter
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Inshellisense([node])"]
```

### Kingpin

[alecthomas/kingpin] based commands can be bridged with the [`bridge.Kingpin`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: tsh
description: Teleport Command Line Client
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Kingpin([tsh])"]
```

### Urfavecli

[urfave/cli] based commands can be bridged with the [`bridge.Urfavecli`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: tea
description: command line tool to interact with Gitea
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Urfavecli([tea])"]
```

### Yargs
[yargs/yargs] based commands can be bridged with the [`bridge.Yargs`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: ng
description: CLI tool for Angular
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Yargs([ng])"]
```

## Shells

> For shells custom configurations are loaded from [`${UserConfigDir}/carapace/bridge`].
> Invoking completion in shells is quite tricky though and edge cases are likely to fail.

### Bash

Commands registered in [bash] can be bridged with the [`bridge.Bash`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name:  tail
description: output the last part of files
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Bash([tail])"]
```

### Fish

Commands registered in [fish-shell/fish-shell] can be bridged with the [`bridge.Fish`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name:  git
description: the stupid content tracker
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Fish([git])"]
```

### Powershell

Commands registered in [powershell] can be bridged with the [`bridge.Powershell`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name:  ConvertTo-Json
description: convert to json
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Powershell([ConvertTo-Json])"]
```

### Zsh

Commands registered in [zsh] can be bridged with the [`bridge.Zsh`] macro:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name:  git
description: the stupid content tracker
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Zsh([git])"]
```

[lazycomplete]:https://github.com/rsteube/lazycomplete
[shell startup delay]:https://jzelinskie.com/posts/dont-recommend-sourcing-shell-completion/

[bash]:https://www.gnu.org/software/bash/
[`bridge.Bash`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionBash

[kislyuk/argcomplete]:https://github.com/kislyuk/argcomplete
[`bridge.Argcomplete`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionArgcomplete

[carapace-sh/carapace]:https://github.com/carapace-sh/carapace
[`bridge.Carapace`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCarapace

[carapace-sh/carapace-bin]:https://github.com/carapace-sh/carapace-bin
[`bridge.CarapaceBin`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCarapaceBin

[clap-rs/clap]:https://github.com/clap-rs/clap
[`bridge.Clap`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionClap

[pallets/click]:https://github.com/pallets/click
[`bridge.Click`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionClick

[spf13/cobra]:https://github.com/spf13/cobra
[`bridge.Cobra`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCobra

[posener/complete]:https://github.com/posener/complete
[`bridge.Complete`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionComplete

[microsoft/inshellisense]:https://github.com/microsoft/inshellisense
[`bridge.Inshellisense`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionInshellisense

[alecthomas/kingpin]:https://github.com/alecthomas/kingpin
[`bridge.Kingpin`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionKingpin

[powershell]:https://microsoft.com/powershell
[`bridge.Powershell`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionPowershell

[fish-shell/fish-shell]:https://github.com/fish-shell/fish-shell
[`bridge.Fish`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionFish

[urfave/cli]:https://github.com/urfave/cli
[`bridge.Urfavecli`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionUrfavecli

[yargs/yargs]:https://github.com/yargs/yargs
[`bridge.Yargs`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionYargs

[zsh]:https://www.zsh.org/
[`bridge.Zsh`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionZsh

[`${UserConfigDir}/carapace/bridge`]:https://pkg.go.dev/os#UserConfigDir
