package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/MatthewAraujo/commit_message/config"
	"github.com/MatthewAraujo/commit_message/prompt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	if len(os.Args) > 1 {
		task := os.Args[1]
		branch, err := shellCommand("git", "branch")
		if err != nil {
			panic(err)
		}
		diff, err := shellCommand("git", "diff")
		if err != nil {
			panic(err)
		}
		prompt := prompt.GetPrompt()
		openAIClient := openai.NewClient(
			option.WithAPIKey(config.Envs.OpenAi.API_KEY),
		)

		chatCompletion, err := openAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.AssistantMessage(prompt),
				openai.UserMessage(createUserMessage(diff, branch, task)),
			}),
			Model: openai.F(openai.ChatModelGPT4o),
		})
		if err != nil {
			panic(err.Error())
		}
		commitMessage := extractCommitMessage(chatCompletion.Choices[0].Message.Content)
		fmt.Println(commitMessage)

	}
}

func shellCommand(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func extractCommitMessage(response string) string {
	re := regexp.MustCompile(`(?m)^:.*?: .*?:.*?$`)
	match := re.FindString(response)

	if match != "" {
		body := regexp.MustCompile(`(?m)^\s*-.*?$`).FindAllString(response, -1)
		if len(body) > 0 {
			match += "\n\n" + fmt.Sprintf("%s", body)
		}
	}

	return match
}

func createUserMessage(diff, branch, task string) string {
	return fmt.Sprintf(`
Você é um assistente especializado em escrever mensagens de commit perfeitas. Siga estas diretrizes:
 
1. Crie uma mensagem de commit clara e concisa. 
2. Use o formato de commits semânticos, incluindo um tipo (como feat, fix, docs) e um emoji correspondente. 
3. Mantenha a mensagem dentro de 72 caracteres na linha principal e inclua uma descrição mais detalhada no corpo, se necessário. 
4. Inclua informações relevantes da tarefa, quando fornecidas. 
 
Aqui estão os detalhes: 
- **Tarefa**: %s
- **Branch atual**: %s 
- **Alterações**: 
%s

Baseado nisso, escreva a mensagem de commit.`, task, branch, diff)
}
