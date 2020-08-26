FROM golang

RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source <(carapace _carapace bash)\n\
for c in \$(carapace --list); do\n\
  source <(carapace \$c)\n\
done" \
       > /root/.bashrc

RUN ln -s /carapace-completers/carapace/carapace /usr/local/bin/carapace


ENTRYPOINT [ "bash" ]

