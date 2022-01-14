SWAGGER_FILE  := "doc/styra-swagger.json"
SCHEME        := "https"
PKG_DIR       := "./pkg"
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
	@$(INFO) Generating swagger client
	@go generate ${PKG_DIR} || $(FAIL)
	@$(OK) Generated swagger client
