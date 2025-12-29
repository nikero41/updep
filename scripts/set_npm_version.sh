#!/usr/bin/env bash

set -e

# RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

VERSION=$1

echo -e "${GREEN}Version: ${VERSION}${NC}\n"

echo -e "${YELLOW}📤 Setting version on platform packages...${NC}"
find ./npm/platforms -name "package.json" -execdir sh -c "npm pkg set version=$VERSION" ";"

pushd npm

pkgs=$(npm pkg get optionalDependencies | jq -r "keys | .[]")

echo -e "\n${YELLOW}📤 Updating platform versions on main package...${NC}\n"
for pkg in $pkgs; do
  echo -e "${GREEN}Setting version for $pkg...${NC}"
  npm pkg set "optionalDependencies.$pkg"="$VERSION"
done

echo -e "\n${YELLOW}📤 Setting version on main package...${NC}"
npm pkg set version="$VERSION"

popd

echo -e "\n${GREEN}🎉 All packages version set successfully!${NC}"
