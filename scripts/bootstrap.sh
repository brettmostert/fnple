#!/usr/bin/env bash

source "$(dirname "$0")/common.sh"
getProjectDir

echo -e "==> ${CYAN}Bootstrap...${NC}"

if [[ -d ${DIR}/bin ]]; then
	rm -rf ./tmp
fi

echo -e "==> Building & Installing 'bob' the builder"
cd "${DIR}/build" || exit
go install ./cmd/bob
cd "${DIR}" || exit

echo -e "==> Installing 'tern'"
go install github.com/jackc/tern@latest
echo -e "==> Installing 'gofumpt'"
go install mvdan.cc/gofumpt@latest
echo -e "==> Installing 'commitlint', 'commitizen', 'config-conventional'"
npm install -g @commitlint/cli @commitlint/config-conventional @commitlint/cz-commitlint commitizen --silent
echo -e "==> Installing 'markdownlint-cli2'"
npm install -g markdownlint-cli2 --silent

echo -e "==> Setup Git Hooks"
git config core.hooksPath .githooks
cd "${DIR}" || exit

echo -e "==> ${YELLOW}Done... ${GREEN}ALL GOOD!!!${NC}"
echo
exit 0
