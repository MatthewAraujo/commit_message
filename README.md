# Git Commit Message Generator

This application automates the generation of commit messages using OpenAI's GPT models. It allows developers to generate clear, concise, and semantically accurate commit messages based on the changes in their code, including staged files, and the task at hand.

## Features

- Automatically generates commit messages using GPT-4, including semantic commit types (e.g., `feat`, `fix`, `docs`, etc.).
- Supports integration with Git to fetch staged changes using `git diff --cached`.
- Allows custom task description input for more specific commit message generation.
- Automatically normalizes commit message formatting, ensuring consistency in the output.
- Supports API key configuration via command-line flags or environment variables.

## Installation

To get started with the Git Commit Message Generator, follow the steps below:

### Prerequisites

- **Go**: The application is written in Go. Make sure you have Go installed on your machine. You can download and install Go from the official website: https://golang.org/dl/
- **OpenAI API Key**: You will need an OpenAI API key to interact with the GPT-4 model.

### Clone the Repository

```bash
git clone https://github.com/yourusername/git-commit-message-generator.git
cd git-commit-message-generator
```

### Install Dependencies

```bash
go mod tidy
```

### Set Up Your OpenAI API Key

You can set up your OpenAI API key in two ways:
1. By using the `-apikey` flag when running the application.
2. By setting the `OPENAI_API_KEY` environment variable.

To set the environment variable on your machine:

```bash
export OPENAI_API_KEY="your_openai_api_key"
```

### Run the Application

```bash
go run main.go -task "Your commit task description"
```

Optionally, you can pass the `-apikey` flag to specify the API key directly:

```bash
go run main.go -task "Your commit task description" -apikey "your_openai_api_key"
```

## Usage

1. **Task flag (`-task`)**: This is a required argument to describe the task you're working on, which will be included in the commit message.
   
   Example:
   ```bash
   go run main.go -task "Refactor Git commit generation logic"
   ```

2. **API Key flag (`-apikey`)**: Use this flag to pass your OpenAI API key directly. If not provided, the application will attempt to load the key from the environment variable `OPENAI_API_KEY`.

3. **Git Integration**: The application will automatically use `git diff --cached` to check for staged changes and generate a commit message that reflects those changes.

4. **Commit Confirmation**: After generating the commit message, the application will display the message to the user for confirmation. You can then decide whether to proceed with the commit or modify the message.

## Example

### Input
```bash
go run main.go -task "Add new feature for user login"
```

### Output
```
Generated Commit Message:
🎉 feat: Add new user login feature

- Added a new user login functionality using JWT tokens for secure authentication.
- Implemented login page UI and integrated with backend API.
- Fixed minor bugs in the authentication flow.

Do you want to proceed with this commit? (Y/N): Y
```

### After Confirmation:
```
Commit successful!
```

## Contribution

Feel free to fork the repository and submit pull requests! All contributions are welcome.

### Issues
If you encounter any issues, please open an issue in the GitHub repository, and we will get back to you as soon as possible.
