#!/bin/bash

# Variables
BINARY_NAME="commit_message"
CUSTOM_BIN_DIR="$HOME/bin"
BINARY_URL="https://github.com/MatthewAraujo/commit_message/releases/download/binary/commit_message-linux-amd64"

if ! command -v curl &>/dev/null; then
    echo "Error: curl is required but not installed. Please install curl and try again."
    exit 1
fi

if [[ ! -d "$CUSTOM_BIN_DIR" ]]; then
    echo "Creating directory: $CUSTOM_BIN_DIR"
    mkdir -p "$CUSTOM_BIN_DIR"
fi

echo "Downloading $BINARY_NAME from $BINARY_URL"
curl -L -o "$CUSTOM_BIN_DIR/$BINARY_NAME" "$BINARY_URL"
if [[ $? -ne 0 ]]; then
    echo "Error: Failed to download the binary from $BINARY_URL"
    exit 1
fi

echo "Making $CUSTOM_BIN_DIR/$BINARY_NAME executable"
chmod +x "$CUSTOM_BIN_DIR/$BINARY_NAME"

cat <<EOF

Installation complete!

To use '$BINARY_NAME' globally, you need to add '$CUSTOM_BIN_DIR' to your PATH.

Steps:
1. Open your shell configuration file:
   - ~/.bashrc for Bash users
   - ~/.zshrc for Zsh users
   - Other shell config files as appropriate

2. Add the following line at the end of the file:
   export PATH="\$PATH:$CUSTOM_BIN_DIR"

3. Reload your shell configuration:
   source <your-shell-config-file>

4. Verify the installation by running:
   $BINARY_NAME

EOF

