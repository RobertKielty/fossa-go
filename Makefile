# FOSSA's published swagger.json - API definition with version info
FOSSA_REF      := https://app.fossa.com/api/api-docs/swagger.json

# where we record the last version
VERSION_FILE  := api-version.txt

# shell-extracted variables
REMOTE_VERSION := $(shell curl -s $(FOSSA_REF) | jq -r .info.version)
LOCAL_VERSION  := $(shell test -f $(VERSION_FILE) && cat $(VERSION_FILE) || echo "none")

FOSSA_GO       := fossa-go

# TODO Make this work, RN the ifeq seems to take priority
#clean:
#	rm $(VERSION_FILE)
#	rm swagger.json
#	rm -rf ./fossa-go

#ifeq ($(REMOTE_VERSION),$(LOCAL_VERSION))
#$(error Local version is $(LOCAL_VERSION) Latest FOSSA version $(REMOTE_VERSION) are the same)
#endif

all: generate-go
.PHONY: all report-fossa-version generate-go

# download swagger.json, record new version
report-fossa-version:
	@echo "⬇️  Downloading swagger.json…"
	curl -s -O $(FOSSA_REF)
	@echo "✏️  Recording version $(REMOTE_VERSION)…"
	@echo $(REMOTE_VERSION) > $(VERSION_FILE)

install-openapi-generator-cli:
	@echo "🔍  Checking for openapitools/openapi-generator-cli…"
	@docker image inspect openapitools/openapi-generator-cli > /dev/null 2>&1 \
		&& echo "✅  openapitools/openapi-generator-cli is already available locally." \
		|| (echo "⬇️  Image not found locally, pulling…" && docker pull openapitools/openapi-generator-cli)

generate-go: install-openapi-generator-cli report-fossa-version
	@echo "🚀  Generating Go client for version $(REMOTE_VERSION)…"
	mkdir -p $(FOSSA_GO)
	chown $$(id -u):$$(id -g) $(FOSSA_GO)
	docker run --rm \
	  -u $$(id -u):$$(id -g) \
	  -v "$${PWD}:/local" \
	  openapitools/openapi-generator-cli generate \
	    -i $(FOSSA_REF) \
	    --skip-validate-spec \
	    -g go \
	    --additional-properties=packageName=fossa,packageVersion=$(REMOTE_VERSION) \
	    --git-user-id=RobertKielty \
	    --git-repo-id=fossa-go \
	    -o /local/
	@echo "✅  Done. Generated go client for FOSSA REST API version $(REMOTE_VERSION) in $(FOSSA_GO)"

.PHONY: prepare-release
prepare-release:
	@echo "🔖 Preparing a release of fossa-go for FOSSA's REST API version $(REMOTE_VERSION)…"
	@git add .
	@git commit -sS -m "adds generated fossa-go for API $(REMOTE_VERSION)"
	@git tag $(REMOTE_VERSION)
	@echo "✅ Tagged commit as ‘$(REMOTE_VERSION)’"
	@echo "Next: push your branch and tags with:"
	@echo "      git push origin main"
	@echo "      git push origin --tags"

.PHONY: publish-release
publish-release:
	@echo "🚀 Publishing GitHub release v$(REMOTE_VERSION)…"
	@gh release create $(REMOTE_VERSION) \
		--title "v$(REMOTE_VERSION)" \
		--notes "Auto-generated FOSSA Go client for API v$(REMOTE_VERSION)"
	@echo "✅ Release v$(REMOTE_VERSION) created on GitHub"
