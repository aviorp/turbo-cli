#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo -e "${BLUE}🚀 Installing turbo-cli...${NC}"
echo -e "${BLUE}📂 Scripts directory: ${SCRIPT_DIR}${NC}"

# Create bin directory if it doesn't exist
mkdir -p ~/bin || {
    echo -e "${RED}❌ Failed to create bin directory${NC}"
    exit 1
}

# Copy the binary to bin directory
cp "${SCRIPT_DIR}/turbo-cli" ~/bin/ || {
    echo -e "${RED}❌ Failed to copy binary${NC}"
    exit 1
}

# Make it executable
chmod 755 ~/bin/turbo-cli || {
    echo -e "${RED}❌ Failed to set permissions${NC}"
    exit 1
}

# Add to PATH if not already added
if ! grep -q "export PATH=\$HOME/bin:\$PATH" ~/.zshrc; then
    echo 'export PATH=$HOME/bin:$PATH' >> ~/.zshrc
    echo -e "${BLUE}✨ Added ~/bin to PATH in ~/.zshrc${NC}"
fi

# Create alias if it doesn't exist
if ! grep -q "alias tc=" ~/.zshrc; then
    echo 'alias tc="turbo-cli"' >> ~/.zshrc
    echo -e "${BLUE}✨ Added 'tc' alias to ~/.zshrc${NC}"
fi

# Create additional helpful aliases
if ! grep -q "# Turbo CLI aliases" ~/.zshrc; then
    echo -e "\n# Turbo CLI aliases" >> ~/.zshrc
    echo 'alias tca="tc create app"' >> ~/.zshrc
    echo 'alias tcs="tc create service"' >> ~/.zshrc
    echo 'alias tcc="tc create controller"' >> ~/.zshrc
    echo 'alias tcm="tc create middleware"' >> ~/.zshrc
    echo -e "${BLUE}✨ Added helpful aliases to ~/.zshrc${NC}"
fi

echo -e "${GREEN}✨ turbo-cli has been installed successfully!${NC}"
echo -e "${BLUE}📝 Please run: source ~/.zshrc${NC}"
echo -e "${BLUE}🚀 You can now use:${NC}"
echo -e "${GREEN}  turbo-cli${NC} - Full command"
echo -e "${GREEN}  tc${NC}         - Short alias"
echo -e "${GREEN}  tca${NC}        - Create app"
echo -e "${GREEN}  tcs${NC}        - Create service"
echo -e "${GREEN}  tcc${NC}        - Create controller"
echo -e "${GREEN}  tcm${NC}        - Create middleware"
echo -e "${BLUE}🎯 Try the interactive mode with:${NC} ${GREEN}tc new${NC}"
