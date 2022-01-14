PKG_DIR       := "./pkg"
TEST_DIR      := "./internal"
OUTPUT_CLIENT := "${PKG_DIR}/client"
OUTPUT_MODELS := "${PKG_DIR}/models"

# ====================================================================================
# Colors

RED          := $(shell printf "\033[31m")
GREEN        := $(shell printf "\033[32m")
BLUE         := $(shell printf "\033[34m")
CNone        := $(shell printf "\033[0m")

# ====================================================================================
# Logger

TIME_SHORT      = `date +%H:%M:%S`
TIME            = $(TIME_SHORT)

INFO    = echo ${TIME} ${BLUE}[ .. ]${CNone}
OK              = echo ${TIME} ${GREEN}[ OK ]${CNone}
FAIL    = (echo ${TIME} ${RED}[FAIL]${CNone} && false)

# ====================================================================================
# Targets

clean:
	@$(INFO) Removing generated client
	-@rm -rf pkg/models pkg/client
	@$(OK) Removed generated client

generate:
	@$(INFO) Generating go resources
	@go generate ${PKG_DIR} || $(FAIL)
	@$(OK) Successfully generated go resources

test:
	@$(INFO) go test unit-tests
	@CGO_ENABLED=0 go test ${TEST_DIR} || $(FAIL)
	@$(OK) go test unit-tests

# ensure generate target doesn't create a diff
check-diff: clean generate
	@$(INFO) checking that branch is clean
	@if git status --porcelain | grep . ; then $(FAIL); else $(OK) branch is clean; fi
