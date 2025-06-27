package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shayansm2/shell-genie/llm"
)

const EXIT = "q"

func main() {
	reader := bufio.NewReader(os.Stdin)
	ollama := llm.GetLLM()
	for {
		fmt.Print("> ")
		userPrompt, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		userPrompt = strings.TrimSpace(userPrompt)
		if userPrompt == EXIT {
			break
		}
		response, err := ollama.Generate(userPrompt)
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	}
}
