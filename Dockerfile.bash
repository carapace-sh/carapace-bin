FROM golang

RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source <(carapace-completers _carapace bash)\n\
for c in \$(carapace-completers --list); do\n\
  source <(carapace-completers \$c)\n\
done" \
       > /root/.bashrc

RUN ln -s /carapace-completers/carapace-completers /usr/local/bin/carapace-completers


ENTRYPOINT [ "bash" ]

