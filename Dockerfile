FROM golang

# shells
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb \
 && dpkg -i packages-microsoft-prod.deb \
 && apt-get update

RUN apt-get install -y fish \
                       elvish \
                       powershell \
                       zsh

RUN apt-get install -y gradle \
                       exa

ENV GOPATH /go
RUN ln -s /carapace-completers/carapace/carapace /usr/local/bin/carapace

# bash
RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source <(carapace _carapace bash)\n\
for c in \$(carapace --list); do\n\
  source <(carapace \$c)\n\
done" \
       > /root/.bashrc

# fish
RUN mkdir -p /root/.config/fish \
 && echo "\n\
function fish_prompt \n\
    set_color cyan \n\
    echo -n 'carapace ' \n\
    set_color normal\n\
end\n\
carapace _carapace fish | source\n\
rm -f /usr/share/fish/completions/chown.fish\n\
for c in (carapace --list)\n\
  complete --erase -c "\$c"\n\
  carapace \$c | source\n\
end" \
       > /root/.config/fish/config.fish

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-v0.14.1.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo "\
carapace _carapace elvish > /root/.elvish/lib/carapace.elv\n\
use carapace\n\
carapace --list | each [c]{ carapace \$c > /tmp/\$c; -source /tmp/\$c }" \
  > /root/.elvish/rc.elv

# powershell
RUN mkdir -p /root/.config/powershell \
 && echo "\n\
function prompt {Write-Host \"carapace\" -NoNewLine -ForegroundColor 3; return \" \"}\n\
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete\n\
carapace _carapace powershell | out-string | Invoke-Expression\n\
carapace --list | foreach { carapace \$_ | Out-String | Invoke-Expression} " \
       > /root/.config/powershell/Microsoft.PowerShell_profile.ps1

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
