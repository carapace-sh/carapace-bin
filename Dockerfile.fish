FROM golang

RUN apt-get update && apt-get install -y fish

RUN mkdir -p /root/.config/fish \
 && echo "\n\
function fish_prompt \n\
    set_color cyan \n\
    echo -n 'carapace ' \n\
    set_color normal\n\
end\n\
carapace-completers _carapace fish | source\n\
rm -f /usr/share/fish/completions/chown.fish\n\
for c in (carapace-completers --list)\n\
  complete --erase -c "\$c"\n\
  carapace-completers \$c | source\n\
end" \
       > /root/.config/fish/config.fish

RUN ln -s /carapace-completers/carapace-completers /usr/local/bin/carapace-completers

ENTRYPOINT [ "/usr/bin/fish" ]

