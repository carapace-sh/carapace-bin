package ffmpeg

import (
	"github.com/carapace-sh/carapace"
)

// ActionHelpTopics completes help topics
//
//	long (Print advanced tool options in addition to the basic tool opti
//	muxer= (Print detailed information about the muxer)
func ActionHelpTopics() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"long", "Print advanced tool options in addition to the basic tool options.",
				"full", "Print complete list of options",
				"decoder=", "Print detailed information about the decoder",
				"encoder=", "Print detailed information about the encoder",
				"demuxer=", "Print detailed information about the demuxer",
				"muxer=", "Print detailed information about the muxer",
				"filter=", "Print detailed information about the filter",
				"bsf=", "Print detailed information about the bitstream filter",
				"protocol=", "Print detailed information about the protocol",
			)
		default:
			switch c.Parts[0] {
			case "decoder":
				return ActionDecoders(DecoderOpts{}.Default())
			case "encoder":
				return ActionEncoders(EncoderOpts{}.Default())
			case "demuxer":
				return ActionDemuxers()
			case "muxer":
				return ActionMuxers()
			case "filter":
				return ActionFilters()
			case "bsf":
				return ActionBitstreamFilters()
			case "protocol":
				return ActionProtocols()
			default:
				return carapace.ActionValues()
			}
		}
	})
}
