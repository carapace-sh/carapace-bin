package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// BlockOpts provides scope filters for block completion
type BlockOpts struct {
	Server  string // host label or server address
	Session string // session target
	Window  string // window target
}

func (o BlockOpts) cacheKey() key.Key {
	return key.String(o.Server, o.Session, o.Window)
}

// ActionBlocks completes block targets
//
//	block:01a14a6c-c0d6-7a48-92f9-86f965f29d47 (shell)
func ActionBlocks(opts BlockOpts) carapace.Action {
	args := []string{"block", "ls", "--json"}
	if opts.Session != "" {
		args = append(args, "--session", opts.Session)
	}
	if opts.Window != "" {
		args = append(args, "--window", opts.Window)
	}
	if opts.Server != "" {
		args = append(args, "--server", opts.Server)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		var blocks struct {
			Blocks []struct {
				BlockId string `json:"block_id"`
				Label   string `json:"label"`
			}
		}
		if err := json.Unmarshal(output, &blocks); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0, len(blocks.Blocks)*2)
		for _, block := range blocks.Blocks {
			if block.Label != "" {
				vals = append(vals, block.BlockId, block.Label)
			} else {
				vals = append(vals, block.BlockId)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(time.Minute, opts.cacheKey()).Tag("blocks")
}
