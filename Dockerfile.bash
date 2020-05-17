FROM golang

RUN echo "\n\
PS1=$'\e[0;36mcarapace \e[0m'\n\
source <(example _carapace bash)" \
       > /root/.bashrc

RUN ln -s /carapace-completers/carapace-completers /usr/local/bin/carapace-completers


ENTRYPOINT [ "bash" ]

