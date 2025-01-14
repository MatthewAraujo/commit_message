#!/bin/bash

# Variables
BINARY_NAME="commit_message"
CUSTOM_BIN_DIR="$HOME/bin"

# Check if the binary exists
if [[ ! -f "$BINARY_NAME" ]]; then
    echo "Error: Binary '$BINARY_NAME' not found in the current directory."
    echo "Please place the binary in this directory and run the script again."
    exit 1
fi

# Create the custom directory if it doesn't exist
if [[ ! -d "$CUSTOM_BIN_DIR" ]]; then
    echo "Creating directory: $CUSTOM_BIN_DIR"
    mkdir -p "$CUSTOM_BIN_DIR"
fi

# Copy the binary to the custom directory
echo "Copying $BINARY_NAME to $CUSTOM_BIN_DIR"
cp "$BINARY_NAME" "$CUSTOM_BIN_DIR/"

# Ensure the binary is executable
echo "Making $CUSTOM_BIN_DIR/$BINARY_NAME executable"
chmod +x "$CUSTOM_BIN_DIR/$BINARY_NAME"

# Display instructions for the user
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

