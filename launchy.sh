#!/bin/bash

# Launchy launcher script for Wayland compatibility
# This script sets the necessary environment variables for Wails to work properly on Wayland

export WEBKIT_DISABLE_COMPOSITING_MODE=1
export WEBKIT_FORCE_COMPLEX_TEXT=1

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Run the Launchy application
"$SCRIPT_DIR/bin/Launchy" "$@"
