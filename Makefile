APP_NAME = jpr
APP_VERSION = 0.1.0

# build with version infos
VERSION_DIR = github.com/jiuhuche120/${APP_NAME}
BUILD_DATE = $(shell date +%FT%T)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

GO_LDFLAGS += -X "${VERSION_DIR}.BuildDate=${BUILD_DATE}"
GO_LDFLAGS += -X "${VERSION_DIR}.CurrentCommit=${GIT_COMMIT}"
GO_LDFLAGS += -X "${VERSION_DIR}.CurrentBranch=${GIT_BRANCH}"
GO_LDFLAGS += -X "${VERSION_DIR}.CurrentVersion=${APP_VERSION}"

RED=\033[0;31m
GREEN=\033[0;32m
BLUE=\033[0;34m
NC=\033[0m

install:
	go install -ldflags '${GO_LDFLAGS}' ./cmd/${APP_NAME}
	@printf "${GREEN}Build ${APP_NAME} successfully!${NC}\n"