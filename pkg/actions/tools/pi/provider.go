package pi

import "github.com/carapace-sh/carapace"

// ActionProviders completes providers
//
//	anthropic
//	openai
func ActionProviders() carapace.Action {
	return carapace.ActionValuesDescribed(
		"anthropic", "Anthropic",
		"azure", "Azure",
		"bedrock", "Bedrock",
		"cerebras", "Cerebras",
		"google", "Google",
		"groq", "Groq",
		"huggingface", "Hugging Face",
		"kimi", "Kimi for Coding",
		"llamacpp", "llama.cpp",
		"minimax", "MiniMax",
		"mistral", "Mistral",
		"nvidia", "NVIDIA",
		"ollama", "Ollama",
		"openai", "OpenAI",
		"openrouter", "OpenRouter",
		"xai", "xAI",
	).Tag("providers")
}
