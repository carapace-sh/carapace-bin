package lazyinit

import (
	"fmt"
	"strings"
)

func CmdClink(completers []string) string {
	snippet := `
local function carapace_completion(command)
    return function(word, word_index, line_state, match_builder)
        match_builder:setnosort()
        match_builder:setvolatile()
        os.setenv('CARAPACE_COMPLINE', line_state:getline():sub(1, line_state:getcursor()))

        local file, pclose = io.popenyield(string.format('carapace %%s cmd-clink ""', command))

        if not file then
            return false
        end

        for line in file:lines() do
            local matches = string.explode(line, '\t')

            if matches[1] then
                match_builder:addmatch({
                    match       = matches[1],
                    display     = matches[2],
                    description = matches[3],
                    type        = 'word',
                    appendchar  = matches[4] or ''
                })
            end
        end

        if pclose then
            pclose()
        else
            file:close()
        end

        return not match_builder:isempty()
    end
end

%s
`
	argmatchers := make([]string, 0, len(completers))
	for _, completer := range completers {
		argmatchers = append(argmatchers,
			fmt.Sprintf(`clink.argmatcher(50, '%[1]s', '%[1]s.exe'):addarg({nowordbreakchars="'`+"`"+`=+;,", carapace_completion('%[1]s')}):loop(1)`, completer),
		)
	}
	return fmt.Sprintf(snippet, strings.Join(argmatchers, "\n"))
}
