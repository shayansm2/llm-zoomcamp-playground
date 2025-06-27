package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ollama struct {
	apiAddr string
	model   string
}

type OllamagenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
}

func (llm Ollama) Generate(prompt string) (string, error) {
	body := OllamagenerateRequest{
		Model:  llm.model,
		Prompt: prompt,
		Stream: false,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%v/%v", llm.apiAddr, "generate")
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status:%v", resp.Status)
	}
	var response OllamaResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", err
	}
	return string(response.Response), nil
}

func (llm Ollama) Chat(userPrompt string, systemPrompt string) string {
	return ""
}

func GetLLM() Ollama {
	return Ollama{
		apiAddr: "http://localhost:11434/api",
		model:   "llama3.2",
	}
}
