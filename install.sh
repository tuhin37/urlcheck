#!/bin/bash

# Define the URL of the binary
URL="https://github.com/tuhin37/urlcheck/blob/main/bin/urlcheck-darwin-amd64?raw=true"

# Define the target installation directory
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="urlcheck"

# Download the binary
echo "Downloading the binary from $URL..."
curl -L -o "$BINARY_NAME" "$URL"

# Check if the download was successful
if [ $? -ne 0 ]; then
    echo "Failed to download the binary. Exiting."
    exit 1
fi

# Move the binary to /usr/local/bin
echo "Installing the binary to $INSTALL_DIR..."
sudo mv "$BINARY_NAME" "$INSTALL_DIR/"

# Make the binary executable
echo "Making the binary executable..."
sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Confirm installation
echo "Installation completed. You can now run the binary using the command '$BINARY_NAME'."

# Verify installation
if command -v $BINARY_NAME >/dev/null 2>&1; then
    echo "Installation successful!"
else
    echo "Installation failed. Please check the steps."
fi
