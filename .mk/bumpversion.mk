# Makefile Configuration for bumpversion

__BUMPVERSION__ := .bumpversion

export BUMPVERSION := $(shell cat $(__BUMPVERSION__))

# semver schemata functionality
## function "set" assign new values to variables
set = $(eval $1 := $2)

## semver functions to set the version consistently
define set-major
  	$(call set,major,$(shell expr $(call major, $(1)) + 1))
  	$(call set,minor,0)
  	$(call set,patch,0)
  	$(call major).$(call minor).$(call patch)
endef

define set-minor
  	$(call set,minor,$(shell expr $(call minor, $(1)) + 1))
  	$(call set,patch,0)
  	$(call major,$(1)).$(call minor).$(call patch)
endef

define set-patch
  	$(call set,patch,$(shell expr $(call patch, $(1)) + 1))
  	$(call major,$(1)).$(call minor,$(1)).$(call patch)
endef

define set-bumpversion
	$(eval current_version = $(shell cat $(__BUMPVERSION__)))
	$(eval new_version := $(shell echo $(call set-$(1),$(current_version))))
	@echo $(new_version) > $(__BUMPVERSION__)
	@echo "Updating version: $(current_version) -> $(new_version)"
endef

semver = $(word $(1), $(subst +, ,$(2)))
version = $(call semver, 1,$(1))

version_core = $(word 1, $(subst -, ,$(call version,$(1))))
version_core_items = $(subst ., ,$(call version_core,$(1)))
version_core_item = $(word $(2), $(call version_core_items,$(1)))

major = $(call version_core_item, $(1),1)
minor = $(call version_core_item, $(1),2)
patch = $(call version_core_item, $(1),3)

all: help
.PHONY: all

##@ Version Management
bumpversion: ## Display the current version
	@cat $(__BUMPVERSION__)

bumpversion/major: ## Increase the major version by 1 (e.g., 1.2.3 -> 2.0.0)
	@$(call set-bumpversion,major)

bumpversion/minor: ## Increase the minor version by 1 (e.g., 1.2.3 -> 1.3.0)
	@$(call set-bumpversion,minor)

bumpversion/patch: ## Increase the patch version by 1 (e.g., 1.2.3 -> 1.2.4)
	@$(call set-bumpversion,patch)