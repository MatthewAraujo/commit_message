package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/MatthewAraujo/commit_message/config"
	"github.com/MatthewAraujo/commit_message/prompt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	taskFlag := flag.String("task", "", "Descrição da tarefa para o commit")
	apiKeyFlag := flag.String("apikey", "", "API Key para OpenAI (pode ser configurado via variável de ambiente)")

	flag.Parse()

	if *taskFlag == "" {
		fmt.Println("Erro: A flag -task é obrigatória.")
		flag.Usage()
		os.Exit(1)
	}

	if *apiKeyFlag == "" && config.Envs.OpenAi.API_KEY == "" {
		fmt.Println("Erro: A chave da API OpenAI não foi fornecida. Use a flag -apikey ou defina a variável de ambiente.")
		os.Exit(1)
	}

	apiKey := *apiKeyFlag
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}

	if isGitRepositoryClean() {
		fmt.Println("Erro: O repositório tem alterações não commitadas. Por favor, commit ou descarte as mudanças antes de continuar.")
		os.Exit(1)
	}

	branch, err := getGitBranch()
	if err != nil {
		log.Fatalf("Erro ao obter branch do git: %v", err)
	}

	diff, err := getGitDiff()
	if err != nil {
		log.Fatalf("Erro ao obter diferenças do git: %v", err)
	}

	openAIClient := openai.NewClient(option.WithAPIKey(apiKey))
	commitMessage, err := generateCommitMessage(openAIClient, diff, branch, *taskFlag)
	if err != nil {
		log.Fatalf("Erro ao gerar mensagem de commit: %v", err)
	}
	// Exibe a mensagem de commit gerada
	fmt.Println("Aqui está a mensagem de commit gerada:")
	fmt.Println(commitMessage)

	// Pergunta ao usuário se ele quer continuar com o commit
	var response string
	fmt.Print("Você deseja continuar com esse commit? (s/n): ")
	fmt.Scanln(&response)

	if strings.ToLower(response) != "s" {
		fmt.Println("Commit cancelado.")
		return
	}

	err = commitChanges(commitMessage)
	if err != nil {
		log.Fatalf("Erro ao realizar o commit: %v", err)
	}

	fmt.Println("Commit realizado com sucesso!")
}

func isGitRepositoryClean() bool {
	status, err := shellCommand("git", "status", "--porcelain")
	if err != nil {
		log.Printf("Erro ao verificar o status do git: %v", err)
		return false
	}

	return status == ""
}

func getGitBranch() (string, error) {
	return shellCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
}

func getGitDiff() (string, error) {
	return shellCommand("git", "diff", "--cached")
}

func shellCommand(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("comando falhou: %v\n%s", err, output)
	}
	return string(output), nil
}

func generateCommitMessage(client *openai.Client, diff, branch, task string) (string, error) {
	prompt := prompt.GetPrompt()
	userMessage := createUserMessage(diff, branch, task)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.AssistantMessage(prompt),
			openai.UserMessage(userMessage),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", err
	}
	fmt.Printf("chatCompletion.Choices[0].Message.Content: %v\n", chatCompletion.Choices[0].Message.Content)
	return extractCommitMessage(chatCompletion.Choices[0].Message.Content), nil
}

func normalizeMessage(line string) string {
	line = strings.TrimSpace(line)
	line = strings.TrimLeft(line, "0123456789.*- ")
	line = strings.Trim(line, "`\"'")
	line = strings.ReplaceAll(line, "\\n", "")
	line = strings.ReplaceAll(line, ": `", ":")
	line = strings.ReplaceAll(line, "`:", ":")

	return line
}

func extractCommitMessage(response string) string {
	lines := strings.Split(response, "\n")
	var commitMessage string

	for _, line := range lines {
		normalizedLine := normalizeMessage(line)

		if normalizedLine != "" {
			if commitMessage != "" {
				commitMessage += "\n"
			}
			commitMessage += normalizedLine
		}
	}

	return commitMessage
}

// Função para realizar o commit no git
func commitChanges(commitMessage string) error {
	_, err := shellCommand("git", "commit", "-m", commitMessage)
	return err
}

// Função para criar a mensagem do usuário para a OpenAI
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
