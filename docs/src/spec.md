# Spec

Custom completions can be defined using [yaml specs](https://github.com/rsteube/carapace-spec).

```sh
# bash
source <(carapace --spec example.yaml) 

# elvish
eval (carapace --spec example.yaml|slurp)

# fish
carapace --spec example.yaml | source 

# oil
source <(carapace --spec example.yaml) 

# nushell
carapace --spec example.yaml | save example.nu ; nu -c 'source example.nu' 

# powershell
carapace --spec example.yaml | Out-String | Invoke-Expression 

# tcsh
eval `carapace --spec example.yaml` 

# xonsh
exec($(carapace --spec example.yaml)) 

# zsh
source <(carapace --spec example.yaml) 
```

## Custom Macros

Carapace provides a range of [custom macros](./spec/macros.md):

```sh
carapace --macros                       # list macros
carapace --macros color.HexColors       # show macro details
carapace --macros color.HexColors <TAB> # test macro
```
