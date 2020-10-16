FROM golang

# shells
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb \
 && dpkg -i packages-microsoft-prod.deb \
 && apt-get update

RUN apt-get install -y bash-completion \ 
                       fish \
                       elvish \
                       powershell \
                       python3-pip \
                       zsh

ENV GOPATH /go
RUN ln -s /carapace-bin/cmd/carapace/carapace /usr/local/bin/carapace

# bash
RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source /usr/share/bash-completion/bash_completion \n\
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
carapace --list | each [c]{\n\
    edit:completion:arg-completer[\$c] = [@arg]{\n\
        carapace \$c > /tmp/carapace_\$c; -source /tmp/carapace_\$c\n\
        \$edit:completion:arg-completer[\$c] \$@arg\n\
    }\n\
}" \
  > /root/.elvish/rc.elv

# oil
RUN curl https://www.oilshell.org/download/oil-0.8.0.tar.gz | tar -xvz \
 && cd oil-*/ \
 && ./configure \
 && make \
 && ./install

RUN mkdir -p ~/.config/oil \
 && echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source <(sed 's/let \"OPTIND += 1\"/(( OPTIND += 1 ))/' /usr/share/bash-completion/bash_completion)\n\
source <(carapace _carapace)\n\
for c in \$(carapace --list); do\n\
  source <(carapace \$c)\n\
done" \
       > ~/.config/oil/oshrc

# powershell
RUN mkdir -p /root/.config/powershell \
 && echo "\n\
function prompt {Write-Host \"carapace\" -NoNewLine -ForegroundColor 3; return \" \"}\n\
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete\n\
carapace _carapace powershell | out-string | Invoke-Expression\n\
carapace --list | foreach { carapace \$_ | Out-String | Invoke-Expression} " \
       > /root/.config/powershell/Microsoft.PowerShell_profile.ps1

# xonsh
RUN pip3 install --no-cache-dir --disable-pip-version-check xonsh \
 && ln -s $(which xonsh) /usr/bin/xonsh

RUN mkdir -p ~/.config/xonsh \
 && echo "\n\
\$COMPLETIONS_CONFIRM=True\n\
exec(\$(carapace _carapace xonsh)) \n\
for \$c in \$(carapace --list).split('\\\n'):\n\
    exec(\$(carapace \$c))"\
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
