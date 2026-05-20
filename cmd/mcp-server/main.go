package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Suggestion represents a single completion suggestion.
// This struct is simplified based on common autocomplete needs.
// A real carapace integration might have more fields, like 'Icon', 'Priority', etc.
type Suggestion struct {
	Value       string `json:"value"`                // The actual string to insert
	Display     string `json:"display,omitempty"`    // Optional: string to show to the user
	Description string `json:"description,omitempty"`// Optional: detailed description
	// Icon      string `json:"icon,omitempty"`     // Optional: icon for the suggestion in UI
	// Priority  int    json:"priority,omitempty"` // Optional: sorting priority for display
}

// AutocompleteRequest defines the structure for an incoming autocomplete request.
type AutocompleteRequest struct {
	Query string `json:"query"` // The current text typed by the user
	// Context map[string]string `json:"context,omitempty"` // Optional: environmental context, e.g., current directory, shell type
}

// AutocompleteResponse defines the structure for the server's response.
type AutocompleteResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
	Error       string       `json:"error,omitempty"` // For internal errors that might be safe to expose
}

// autocompleteHandler handles requests to the /autocomplete endpoint.
func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are accepted", http.StatusMethodNotAllowed)
		return
	}

	// Set a context with a timeout for the LLM request.
	// This prevents a slow LLM from holding up the HTTP handler indefinitely.
	// The timeout duration should be configurable, perhaps via an environment variable.
	llmTimeoutStr := os.Getenv("LLM_TIMEOUT_SECONDS")
	llmTimeout := 5 * time.Second // Default timeout for LLM interaction
	if d, err := time.ParseDuration(llmTimeoutStr + "s"); err == nil && llmTimeoutStr != "" {
		llmTimeout = d
	} else if llmTimeoutStr != "" {
		log.Printf("Warning: Invalid LLM_TIMEOUT_SECONDS '%s', using default %v", llmTimeoutStr, llmTimeout)
	}

	ctx, cancel := context.WithTimeout(r.Context(), llmTimeout)
	defer cancel() // Ensure the context is cancelled to release resources

	defer r.Body.Close() // Ensure the request body is closed

	var req AutocompleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received autocomplete request for query: '%s'", req.Query)

	// --- LLM Integration ---
	// This function would interact with an LLM API (e.g., OpenAI, Anthropic, custom local LLM).
	// It sends `req.Query` and potentially other context, then parses the LLM's response.
	suggestions, err := getLLMCompletions(ctx, req.Query) // Pass the context for cancellation/timeout
	if err != nil {
		// Differentiate between LLM error and context timeout/cancellation
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Error getting LLM completions (timeout after %v): %v", llmTimeout, err)
			http.Error(w, fmt.Sprintf("LLM request timed out after %v", llmTimeout), http.StatusGatewayTimeout)
		} else {
			log.Printf("Error getting LLM completions: %v", err)
			http.Error(w, fmt.Sprintf("Failed to get completions from LLM: %v", err), http.StatusInternalServerError)
		}
		return
	}
	// --- End LLM Integration ---

	resp := AutocompleteResponse{
		Suggestions: suggestions,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		// It's generally too late to send an http.Error if WriteHeader has implicitly
		// been called by Encode. Just log and continue.
	}
}

// getLLMCompletions simulates calling an LLM to generate suggestions.
// In a real application, this would involve:
// - Making HTTP calls to an LLM provider (e.g., OpenAI, Anthropic) using http.Client with context.
// - Securely managing API keys (e.g., from environment variables, secret manager).
// - Structuring prompts to elicit desired autocomplete format (e.g., JSON).
// - Implementing retry logic with exponential backoff for transient errors.
// - Potentially caching responses for common queries to reduce latency and cost.
// - Robust error handling specific to the LLM API.
func getLLMCompletions(ctx context.Context, query string) ([]Suggestion, error) {
	// In a real scenario, this would be entirely replaced by actual LLM API calls.
	// For example:
	// client := &http.Client{} // Configure with Transport, timeout if not using context timeout directly for client
	// reqBody := map[string]interface{}{
	//     "model": "gpt-3.5-turbo", // Or other LLM model
	//     "messages": []map[string]string{
	//         {"role": "system", "content": "You are a CLI autocomplete assistant. Provide suggestions as a JSON array of objects with 'Value', 'Display', and 'Description' fields."},
	//         {"role": "user", "content": "Provide CLI autocomplete suggestions for: " + query},
	//     },
	//     "max_tokens": 100,
	// }
	// jsonBody, _ := json.Marshal(reqBody)
	//
	// llmReq, err := http.NewRequestWithContext(ctx, "POST", os.Getenv("LLM_API_ENDPOINT"), bytes.NewBuffer(jsonBody))
	// if err != nil {
	//     return nil, fmt.Errorf("failed to create LLM request: %w", err)
	// }
	// llmReq.Header.Set("Content-Type", "application/json")
	// llmReq.Header.Set("Authorization", "Bearer "+os.Getenv("LLM_API_KEY")) // Securely manage this key
	//
	// llmResp, err := client.Do(llmReq)
	// if err != nil {
	//     return nil, fmt.Errorf("LLM API call failed: %w", err)
	// }
	// defer llmResp.Body.Close()
	//
	// if llmResp.StatusCode != http.StatusOK {
	//     bodyBytes, _ := io.ReadAll(llmResp.Body)
	//     return nil, fmt.Errorf("LLM API returned non-200 status: %d, body: %s", llmResp.StatusCode, string(bodyBytes))
	// }
	//
	// // Parse LLM response, typically JSON, into []Suggestion
	// var llmSuggestions []Suggestion
	// if err := json.NewDecoder(llmResp.Body).Decode(&llmSuggestions); err != nil {
	//     return nil, fmt.Errorf("failed to decode LLM response: %w", err)
	// }
	// return llmSuggestions, nil


	// --- DEMO LOGIC for LLM-like suggestions (retained for illustrative purposes) ---
	// In a real system, this demo logic would be entirely replaced by the actual LLM API call.
	// This "simulation" adds a brief artificial delay to mimic some processing,
	// but the `context` passed to the function handles the overall timeout.
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Context cancelled or timed out before demo logic completes
	case <-time.After(50 * time.Millisecond): // Simulate a very quick LLM internal processing
		// Proceed with demo logic
	}

	switch {
	case len(query) < 2: // Provide broad suggestions for very short queries
		return []Suggestion{
			{Value: "git", Display: "git - the stupid content tracker", Description: "A version control system."},
			{Value: "docker", Display: "docker - A self-sufficient runtime for containers", Description: "Build, share, and run applications with containers."},
			{Value: "npm", Display: "npm - Node Package Manager", Description: "Manage JavaScript packages."},
			{Value: "ls", Display: "ls - list directory contents", Description: "List information about files in current directory."},
		}, nil
	case query == "git c":
		return []Suggestion{
			{Value: "commit", Display: "git commit - Record changes to the repository", Description: "Save changes to the repository."},
			{Value: "checkout", Display: "git checkout - Switch branches or restore working tree files", Description: "Navigate between branches."},
			{Value: "clone", Display: "git clone - Clone a repository into a new directory", Description: "Create a copy of a remote repository."},
		}, nil
	case query == "docker r":
		return []Suggestion{
			{Value: "run", Display: "docker run - Run a command in a new container", Description: "Execute a command inside a new Docker container."},
			{Value: "rmi", Display: "docker rmi - Remove one or more images", Description: "Delete Docker images."},
			{Value: "restart", Display: "docker restart - Restart one or more containers", Description: "Restart running Docker containers."},
		}, nil
	case query == "npm i":
		return []Suggestion{
			{Value: "install", Display: "npm install - Install dependencies in the current project", Description: "Install packages specified in package.json."},
			{Value: "init", Display: "npm init - Create a package.json file", Description: "Initialize a new Node.js project."},
		}, nil
	case query == "ls -":
		return []Suggestion{
			{Value: "-l", Display: "ls -l - Use a long listing format", Description: "Display file details including permissions, size, owner, etc."},
			{Value: "-a", Display: "ls -a - Do not ignore entries starting with .", Description: "Show hidden files."},
			{Value: "-h", Display: "ls -h - With -l, print human readable sizes", Description: "Show file sizes in human-readable format (e.g., 1K, 234M, 2G)."},
		}, nil
	default:
		// Generic LLM-like response for other queries, demonstrating dynamic suggestions
		return []Suggestion{
			{Value: fmt.Sprintf("%s_command", query), Display: fmt.Sprintf("LLM suggestion for '%s'", query), Description: "A command inferred by the LLM."},
			{Value: fmt.Sprintf("%s_option", query), Display: fmt.Sprintf("Another LLM idea for '%s'", query), Description: "An option related to the query, suggested by LLM."},
			{Value: "context_aware_suggestion", Description: "This suggestion is context-aware based on LLM analysis and current environment."},
		}, nil
	}
	// --- END DEMO LOGIC ---
}

func main() {
	// Read the port from environment variable, default to 8080 if not set.
	// This allows flexible deployment (e.g., Kubernetes, Docker).
	port := os.Getenv("MCP_SERVER_PORT")
	if port == "" {
		port = "8080" // Default port for the server
	}

	// Register the autocomplete handler for the /autocomplete endpoint.
	// This is the primary endpoint carapace-bin would call.
	http.HandleFunc("/autocomplete", autocompleteHandler)

	addr := ":" + port
	log.Printf("Starting MCP Autocomplete Server on %s...", addr)

	// Start the HTTP server.
	// log.Fatalf ensures the program exits if the server cannot start.
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("MCP Server failed to start: %v", err)
	}
}