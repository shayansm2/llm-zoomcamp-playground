package prompt

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func Generate(promptName string, values map[string]string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("prompt/%v.txt", promptName))
	if err != nil {
		return "", err
	}

	promptTemplate, err := template.New(promptName).Parse(string(content))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = promptTemplate.Execute(&buf, values)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
