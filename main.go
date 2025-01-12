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
	taskFlag := flag.String("task", "", "Task description for the commit")
	apiKeyFlag := flag.String("apikey", "", "API Key for OpenAI (can be configured via environment variable)")

	flag.Parse()

	if *taskFlag == "" {
		fmt.Println("Error: The -task flag is required.")
		flag.Usage()
		os.Exit(1)
	}

	if *apiKeyFlag == "" && config.Envs.OpenAi.API_KEY == "" {
		fmt.Println("Error: The OpenAI API key was not provided. Use the -apikey flag or set the environment variable.")
		os.Exit(1)
	}

	apiKey := *apiKeyFlag
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}

	if isGitRepositoryClean() {
		fmt.Println("Error: The repository has uncommitted changes. Please commit or discard the changes before continuing.")
		os.Exit(1)
	}

	branch, err := getGitBranch()
	if err != nil {
		log.Fatalf("Error getting the Git branch: %v", err)
	}

	diff, err := getGitDiff()
	if err != nil {
		log.Fatalf("Error getting the Git diff: %v", err)
	}

	openAIClient := openai.NewClient(option.WithAPIKey(apiKey))
	commitMessage, err := generateCommitMessage(openAIClient, diff, branch, *taskFlag)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	fmt.Println("Here is the generated commit message:")
	fmt.Println(commitMessage)

	var response string
	fmt.Print("Do you want to proceed with this commit? (y/n): ")
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("Commit canceled.")
		return
	}

	err = commitChanges(commitMessage)
	if err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}

	fmt.Println("Commit successfully completed!")
}

func isGitRepositoryClean() bool {
	status, err := shellCommand("git", "status", "--porcelain")
	if err != nil {
		log.Printf("Error checking Git status: %v", err)
		return false
	}

	return status == ""
}

func getGitBranch() (string, error) {
	return shellCommand("git", "branch")
}

func getGitDiff() (string, error) {
	return shellCommand("git", "diff", "--cached")
}

func shellCommand(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %v\n%s", err, output)
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
	return extractCommitMessage(chatCompletion.Choices[0].Message.Content), nil
}

func normalizeMessage(line string) string {
	line = strings.TrimSpace(line)

	line = strings.TrimLeft(line, "0123456789.*- ")

	line = strings.Trim(line, "`\"'")

	line = strings.ReplaceAll(line, "\\n", "")

	line = strings.ReplaceAll(line, ": `", ":")
	line = strings.ReplaceAll(line, "`:", ":")

	line = strings.ReplaceAll(line, "*", "")

	line = strings.ReplaceAll(line, ". ", ".\n")

	return line
}

func extractCommitMessage(response string) string {
	lines := strings.Split(response, "\n")
	var commitMessage strings.Builder

	ignoredFirstShortLine := false

	for _, line := range lines {
		normalizedLine := normalizeMessage(line)

		if normalizedLine != "" {
			if len(normalizedLine) < 20 && !ignoredFirstShortLine {
				ignoredFirstShortLine = true
				continue
			}

			if commitMessage.Len() > 0 {
				commitMessage.WriteString("\n - ")
			}

			commitMessage.WriteString(normalizedLine)
		}
	}

	return commitMessage.String()
}

func commitChanges(commitMessage string) error {
	_, err := shellCommand("git", "commit", "-m", commitMessage)
	return err
}

func createUserMessage(diff, branch, task string) string {
	return fmt.Sprintf(`
### Provided Details:
- **Task**: %s
- **Branch**: %s
- **Changes**:
%s

Based on this, generate a commit message.`, task, branch, diff)
}
