FROM golang

RUN apt-get update && apt-get install -y fish

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

RUN ln -s /carapace-completers/carapace/carapace /usr/local/bin/carapace

ENTRYPOINT [ "/usr/bin/fish" ]

