package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MatthewAraujo/commit_message/prompt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const (
	apiKeyFilePathEnv      = ".open_ai_api_key.json"
	minCommitMessageLength = 20
)

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func isGitRepositoryClean() bool {
	status, err := shellCommand("git", "status", "--porcelain")
	checkError(err, "Error checking Git status")
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

func getOpenAIAPIKeyFilePath() string {
	homeDir, err := os.UserHomeDir()
	checkError(err, "Error getting user directory")
	return filepath.Join(homeDir, apiKeyFilePathEnv)
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

func extractCommitMessage(response string) string {
	lines := strings.Split(response, "\n")
	var commitMessage strings.Builder
	ignoredFirstShortLine := false

	for _, line := range lines {
		normalizedLine := normalizeMessage(line)

		if normalizedLine != "" {
			// Ignore the first short line
			if len(normalizedLine) < minCommitMessageLength && !ignoredFirstShortLine {
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

func getAPIKey() (string, error) {
	filePath := getOpenAIAPIKeyFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("API key file not found")
		}
		return "", fmt.Errorf("error opening API key file: %v", err)
	}
	defer file.Close()

	var key map[string]string
	if err := json.NewDecoder(file).Decode(&key); err != nil {
		return "", fmt.Errorf("error reading API key: %v", err)
	}

	apiKey, exists := key["ApiKey"]
	if !exists {
		return "", fmt.Errorf("API key not found in file")
	}

	return apiKey, nil
}

func setupAPIKey(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("API key argument missing")
	}

	filePath := getOpenAIAPIKeyFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error saving API key: %v", err)
	}
	defer file.Close()

	key := map[string]string{"ApiKey": args[2]}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(key); err != nil {
		return fmt.Errorf("error writing API key: %v", err)
	}

	fmt.Println("API key saved successfully.")
	return nil
}

func main() {
	taskFlag := flag.String("task", "", "Task description for the commit (optional)")
	setAPIKeyFlag := flag.Bool("set_api_key", false, "Set the OpenAI API key (use only once to configure)")
	flag.Parse()

	if *setAPIKeyFlag {
		fmt.Println("Setting up the OpenAI API key...")
		if err := setupAPIKey(os.Args); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("✔ API key setup complete. You can now run the commit process.")
		return
	}

	apiKey, err := getAPIKey()
	checkError(err, "Error retrieving the API key")

	if isGitRepositoryClean() {
		fmt.Println("\n⚠️ Error: The repository has uncommitted changes.")
		fmt.Println("Please commit or discard the changes before continuing.")
		os.Exit(1)
	}

	fmt.Println("\n🔍 Retrieving current Git branch and changes...")
	branch, err := getGitBranch()
	checkError(err, "Error getting Git branch")

	diff, err := getGitDiff()
	checkError(err, "Error getting Git diff")

	fmt.Println("\n🧠 Generating commit message based on the changes...")
	openAIClient := openai.NewClient(option.WithAPIKey(apiKey))
	commitMessage, err := generateCommitMessage(openAIClient, diff, branch, *taskFlag)
	checkError(err, "Error generating commit message")

	fmt.Println("\n💬 Here is the generated commit message:")
	fmt.Println("----------------------------------------------------")
	fmt.Println(commitMessage)
	fmt.Println("----------------------------------------------------")

	var response string
	fmt.Print("\nDo you want to proceed with this commit? (y/n): ")
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("\n❌ Commit canceled. No changes were made.")
		return
	}

	fmt.Println("\n✅ Committing changes to Git...")
	if err := commitChanges(commitMessage); err != nil {
		fmt.Printf("Error committing changes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🎉 Commit successfully completed!")
}
