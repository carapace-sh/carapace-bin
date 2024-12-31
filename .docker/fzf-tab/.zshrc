zstyle ':completion:*' menu select
zstyle ':completion:*' matcher-list 'm:{a-zA-Z}={A-Za-z}' 'r:|=*' 'l:|=* r:|=*'
zstyle ':fzf-tab:*' query-string ''

autoload -U compinit && compinit
source <(carapace _carapace zsh)
source ~/fzf-tab/fzf-tab.plugin.zsh
