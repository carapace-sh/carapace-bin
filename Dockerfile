FROM rsteube/carapace

RUN ln -s /carapace-bin/cmd/carapace/carapace /usr/local/bin/carapace

# bash
RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source /usr/share/bash-completion/bash_completion \n\
source <(carapace _carapace)" \
       > /root/.bashrc

# fish
RUN mkdir -p /root/.config/fish \
 && echo "\n\
function fish_prompt \n\
    set_color cyan \n\
    echo -n 'carapace ' \n\
    set_color normal\n\
end\n\
mkdir -p ~/.config/fish/completions\n\
carapace --list | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)\n\
carapace _carapace fish | source" \
       > /root/.config/fish/config.fish

# elvish
RUN mkdir -p /root/.elvish/lib \
 && echo "\
eval (carapace _carapace|slurp)" \
  > /root/.elvish/rc.elv

# oil
RUN mkdir -p ~/.config/oil \
 && echo "source <(carapace _carapace)" \
    > ~/.config/oil/oshrc

# powershell
RUN mkdir -p /root/.config/powershell \
 && echo "\n\
function prompt {Write-Host \"carapace\" -NoNewLine -ForegroundColor 3; return \" \"}\n\
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete\n\
carapace _carapace | out-string | Invoke-Expression" \
       > /root/.config/powershell/Microsoft.PowerShell_profile.ps1

# xonsh
RUN mkdir -p ~/.config/xonsh \
 && echo "\n\
\$COMPLETIONS_CONFIRM=True\n\
exec(\$(carapace _carapace xonsh))"\
  > ~/.config/xonsh/rc.xsh

# zsh
RUN echo "\n\
PS1=$'%{\e[0;36m%}carapace %{\e[0m%}'\n\
\n\
zstyle ':completion:*' menu select \n\
zstyle ':completion:*' matcher-list 'm:{a-zA-Z}={A-Za-z}' 'r:|=*' 'l:|=* r:|=*' \n\
\n\
autoload -U compinit && compinit \n\
source <(carapace _carapace zsh) \n\
for c in \$(carapace --list); do\n\
  source <(carapace \$c)\n\
done"  > /root/.zshrc
