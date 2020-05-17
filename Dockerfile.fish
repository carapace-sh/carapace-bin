FROM golang

RUN apt-get update && apt-get install -y fish

RUN mkdir -p /root/.config/fish \
 && echo "\n\
function fish_prompt \n\
    set_color cyan \n\
    echo -n 'carapace ' \n\
    set_color normal\n\
end\n\
carapace-completers _carapace fish | source" \
       > /root/.config/fish/config.fish

ENV PS1=$'%{\e[0;36m%}carapace %{\e[0m%}'

RUN ln -s /carapace-completers/carapace-completers /usr/local/bin/carapace-completers


ENTRYPOINT [ "fish" ]

