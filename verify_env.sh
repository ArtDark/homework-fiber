#!/bin/bash

# Verify Go environment for project
set -e

echo "🔍 Verifying Go environment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓${NC} $1"
    else
        echo -e "${RED}✗${NC} $1"
        exit 1
    fi
}

# Function to install tool
install_tool() {
    local tool_name=$1
    local install_cmd=$2
    
    echo -e "${YELLOW}⚠ $tool_name not found, attempting to install...${NC}"
    if eval $install_cmd; then
        echo -e "${GREEN}✓${NC} $tool_name installed successfully"
    else
        echo -e "${RED}✗ Failed to install $tool_name${NC}"
        exit 1
    fi
}

# Check Go installation
echo "Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Go is not installed${NC}"
    exit 1
fi
echo -e "${GREEN}✓${NC} Go is installed"

# Check Go version
MIN_GO_VERSION="1.19"
GO_VERSION=$(go version | grep -o 'go[0-9]\+\.[0-9]\+' | cut -d'o' -f2)
if [[ "$(printf '%s\n' "$MIN_GO_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$MIN_GO_VERSION" ]]; then
    echo -e "${RED}✗ Go version $GO_VERSION is less than required $MIN_GO_VERSION${NC}"
    exit 1
fi
echo -e "${GREEN}✓${NC} Go version $GO_VERSION (required: $MIN_GO_VERSION)"

# Check GOPATH
if [[ -z "$GOPATH" ]]; then
    echo -e "${YELLOW}⚠ GOPATH is not set, using default${NC}"
    GOPATH=$(go env GOPATH)
fi
echo -e "${GREEN}✓${NC} GOPATH: $GOPATH"

# Check GOROOT
GOROOT=$(go env GOROOT)
echo -e "${GREEN}✓${NC} GOROOT: $GOROOT"

# Check required tools
echo "Checking required tools..."
REQUIRED_TOOLS=("git" "gcc")
for tool in "${REQUIRED_TOOLS[@]}"; do
    if command -v $tool &> /dev/null; then
        echo -e "${GREEN}✓${NC} $tool found"
    else
        echo -e "${RED}✗ $tool not found${NC}"
        exit 1
    fi
done

# Check templ tool
echo "Checking templ tool..."
if ! command -v templ &> /dev/null; then
    install_tool "templ" "go install github.com/a-h/templ/cmd/templ@latest"
fi
echo -e "${GREEN}✓${NC} templ found"

# Check air tool
echo "Checking air tool..."
if ! command -v air &> /dev/null; then
    install_tool "air" "go install github.com/air-verse/air@v1.62.0"
fi
echo -e "${GREEN}✓${NC} air found"

# Check Go modules
echo "Checking Go modules..."
if [ ! -f "go.mod" ]; then
    echo -e "${YELLOW}⚠ go.mod not found, initializing...${NC}"
    go mod init 2>/dev/null || true
fi

if [ -f "go.mod" ]; then
    echo -e "${GREEN}✓${NC} go.mod found"
    echo "Checking dependencies..."
    go mod tidy >/dev/null 2>&1
    print_status "Dependencies are up to date"
else
    echo -e "${YELLOW}⚠ Could not initialize or find go.mod${NC}"
fi

# Verify build
echo "Verifying build..."
go build -v ./... >/dev/null 2>&1
print_status "Project builds successfully"

echo -e "\n${GREEN}✅ All checks passed! Your environment is ready.${NC}"