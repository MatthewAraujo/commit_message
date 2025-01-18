# Git Commit Message Generator

This project uses OpenAI's GPT-4 model to generate commit messages based on the changes in a Git repository. It analyzes the difference between your current Git changes and generates a commit message that you can review before committing your changes.

## Features

- **Automatic Commit Message Generation**: Generate a commit message based on the changes in your Git repository.
- **Git Integration**: The tool checks if your repository is clean and retrieves the current Git branch and changes.
- **Customizable Task Description**: You can provide a task description for context in the commit message.
- **OpenAI Integration**: Uses OpenAI's API to generate a meaningful commit message based on the provided diff and task details.

## Requirements

- Go 1.18 or higher.
- OpenAI API key (Required to interact with OpenAI GPT-4).
- Git repository with changes.

## Installation

You can install the tool in two different ways: using Go or downloading the executable.

### Option 1: Using Go

1. Run the following command:
   ```bash
   go install github.com/MatthewAraujo/commit_message@latest
   ```
This installs a go binary that will automatically bind to your $GOPATH

> if you’re using Zsh, you’ll need to add it manually to `~/.zshrc`.
```bash
GOPATH=$HOME/go  PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

don't forget to update

```bash
source ~/.zshrc
```

2. You are now ready to use the tool.

### Option 2: Downloading the Executable

1. Download the binary for your operating system (Windows, macOS, or Linux) from the [releases page](https://github.com/MatthewAraujo/commit_message/releases).
2. Follow the steps below to set up and execute the tool for your operating system:

#### For Linux

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/MatthewAraujo/commit_message.git && \
   cd commit_message
   ```
2. Make the setup script executable and run it:
   ```bash
   chmod +x setup-linux.sh
   ./setup-linux.sh
   ```

#### For Windows

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/MatthewAraujo/commit_message.git && \
   cd commit_message
   ```
2. Run the setup script:
   ```cmd
   setup-windows.cmd
   ```

#### For macOS

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/MatthewAraujo/commit_message.git && \
   cd commit_message
   ```
2. Make the setup script executable and run it:
   ```bash
   chmod +x setup-macos.sh
   ./setup-macos.sh
   ```

## Setup API Key

Before using the tool, you need to set up your OpenAI API key:

1. Run the following command to set the API key:
   ```bash
   go run main.go --set_api_key --api_key YOUR_OPENAI_API_KEY
   ```
2. The API key will be stored in a `.open_ai_api_key.json` file in your home directory. If you don't have an OpenAI API key, sign up at [OpenAI](https://platform.openai.com/signup).

## Usage

1. To generate a commit message, run the following command:
   ```bash
   go run main.go --task "TASK_DESCRIPTION"
   ```
   - Replace `TASK_DESCRIPTION` with a short description of the task you're working on (optional).
2. The program will:

   - Check if your repository has uncommitted changes.
   - Retrieve the current Git branch and changes.
   - Generate a commit message based on the Git diff and task description.

3. Review the generated commit message. If you're happy with it, type `y` to commit the changes. Otherwise, type `n` to cancel.

## Example

```bash
$ go run main.go --task "Fix bug in user authentication"
🔍 Retrieving current Git branch and changes...
🧠 Generating commit message based on the changes...
💬 Here is the generated commit message:
----------------------------------------------------
Fix bug in user authentication
 - Fixed issue where users were unable to log in after password reset
----------------------------------------------------
Do you want to proceed with this commit? (y/n): y

✅ Committing changes to Git...
🎉 Commit successfully completed!
```

## Error Handling

- If there are uncommitted changes, the tool will notify you and exit without making any commits.
- If the API key is missing or invalid, an error message will be displayed.
