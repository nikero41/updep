#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

cd npm

APP_NAME=$(jq -r ".name" ./package.json)

echo -e "${YELLOW}📦 Publishing ${APP_NAME} to npm${NC}\n"

if ! npm whoami &>/dev/null; then
    echo -e "${RED}❌ Not logged in to npm. Please run 'npm login' first.${NC}"
    exit 1
fi

VERSION=$(jq -r ".version" ./package.json)
echo -e "${GREEN}Version: ${VERSION}${NC}"
echo ""

echo -e "${YELLOW}📤 Publishing platform-specific packages...${NC}"

PLATFORMS=(./platforms/*)

for platform in "${PLATFORMS[@]}"; do
    echo -e "${GREEN}Publishing ${platform}...${NC}"
    pushd "${platform}"
    npm publish --access public
    popd
    echo -e "${GREEN}✓ ${platform} published${NC}\n"
done

echo -e "${YELLOW}⏳ Waiting 5 seconds for npm to propagate packages...${NC}"
sleep 5

echo -e "${YELLOW}📤 Publishing main package...${NC}"
npm publish --access public
echo -e "${GREEN}✓ updep published${NC}\n"

echo -e "${GREEN}🎉 All packages published successfully!${NC}"
echo -e "${GREEN}You can now install with: npm install -g ${APP_NAME}${NC}"
echo -e "${GREEN}Or run with: npx ${APP_NAME}${NC}"
