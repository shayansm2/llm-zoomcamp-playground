package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shayansm2/shell-genie/llm"
	"github.com/shayansm2/shell-genie/prompt"
)

const EXIT = "q"

func main() {
	reader := bufio.NewReader(os.Stdin)
	ollama := llm.GetLLM()
	for {
		fmt.Print("> ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == EXIT {
			break
		}
		llmPrompt, err := prompt.Generate("template", map[string]string{"Question": userInput})
		if err != nil {
			panic(err)
		}
		response, err := ollama.Generate(llmPrompt)
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	}
}
