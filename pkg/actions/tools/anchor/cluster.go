package anchor

import "github.com/carapace-sh/carapace"

// ActionClusters completes clusters
//
//	localnet (http://127.0.0.1:8899)
//	devnet (https://api.devnet.solana.com)
func ActionClusters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"localnet", "http://127.0.0.1:8899",
		"devnet", "https://api.devnet.solana.com",
		"testnet", "https://api.testnet.solana.com",
		"mainnet", "https://api.mainnet-beta.solana.com",
		"debug", "http://34.90.18.145:8899",
	)
}

// ActionCommitments completes commitment levels
//
//	processed (the node will query its most recent block)
//	finalized (the node will query the most recent block confirmed by supermajority as finalized)
func ActionCommitments() carapace.Action {
	return carapace.ActionValuesDescribed(
		"processed", "the node will query its most recent block",
		"confirmed", "the node will query the most recent block that has been voted on by supermajority",
		"finalized", "the node will query the most recent block confirmed by supermajority as finalized",
	)
}
