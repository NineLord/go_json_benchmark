#region Parameters
GO_COMPILER = go
COMMAND_DIRECTORY = ./cmd
COMMAND_FILES_EXTENTIONS = go
PACKAGE_DIRECTORY = ./pkg
BINARY_DIRECTORY = ./bin
PREFIX_TO_PRINTS = === 
#endregion

#region Colors
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)
#endregion

# Searches for all the executable files that can be compiled in the binary directory
COMMAND_FILES = $(shell ls -1 $(COMMAND_DIRECTORY) | sed -e 's/\.$(COMMAND_FILES_EXTENTIONS)$$//')

.PHONY: all $(COMMAND_FILES) run clean
default: help

#region Help menu
help:
	@echo 'Targets to compile:'
	@for TARGET_NAME in $(COMMAND_FILES); do \
		echo "${CYAN}$$TARGET_NAME${RESET}"; \
	done

	@echo ''
	@echo 'To compile all executables use the target: ${CYAN}all${RESET}'
#endregion

#region Target to compile all executables at once
all: $(COMMAND_FILES)
#endregion

#region Target for each executable files to be compiled
$(COMMAND_FILES):
	@# Saves into variable the actual file name
	$(eval COMMAND_FILE = $@.$(COMMAND_FILES_EXTENTIONS))

	@# Making sure there is binary directory
	@mkdir -p $(BINARY_DIRECTORY)

	@echo "$(PREFIX_TO_PRINTS)${GREEN}Compiling $(COMMAND_FILE)${RESET}"
	$(GO_COMPILER) build -o $(BINARY_DIRECTORY)/$@ $(COMMAND_DIRECTORY)/$(COMMAND_FILE)
#endregion

#region Target for running specific build
#region Fixing run target argument passing
# If the first argument is "run"...
ifeq (run, $(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif
#endregion
run: $(TARGET)
	@echo "$(PREFIX_TO_PRINTS)${GREEN}Running the target ${CYAN}$(TARGET)${RESET}"
	$(BINARY_DIRECTORY)/$(TARGET) $(RUN_ARGS)
#endregion

# Target for removing all previous compilations
clean:
	@echo "$(PREFIX_TO_PRINTS)${GREEN}Removing all compiled files${RESET}"
	rm -rf $(BINARY_DIRECTORY)