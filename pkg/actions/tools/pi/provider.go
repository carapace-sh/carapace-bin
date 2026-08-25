package pi

import "github.com/carapace-sh/carapace"

// ActionProviders completes providers
//
//	anthropic
//	openai
func ActionProviders() carapace.Action {
	return carapace.ActionValuesDescribed(
		"anthropic", "Anthropic",
		"openai", "OpenAI",
		"google", "Google",
		"azure", "Azure",
		"bedrock", "Bedrock",
		"mistral", "Mistral",
		"groq", "Groq",
		"cerebras", "Cerebras",
		"xai", "xAI",
		"huggingface", "Hugging Face",
		"kimi", "Kimi for Coding",
		"minimax", "MiniMax",
		"nvidia", "NVIDIA",
		"openrouter", "OpenRouter",
		"ollama", "Ollama",
		"llamacpp", "llama.cpp",
	).Tag("providers")
}
