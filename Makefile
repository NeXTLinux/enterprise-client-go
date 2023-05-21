# The version of this client (should be in line with the highest supported enterprise version
CLIENT_VERSION = 5.0.0

# where all generated code will be located
PROJECT_ROOT = pkg
CLONE_DIR = local

# OpenAPI generator version to use
OPENAPI_GENERATOR_VERSION = v5.2.1

# --- nextlinux enterprise references
# a git tag/branch/commit within nextlinux/enterprise repo
ENTERPRISE_REF = 894ae07df975363212899193b6e1fdf6c3f4ab6e
ENTERPRISE_ROOT = $(PROJECT_ROOT)/enterprise
ENTERPRISE_OPENAPI_DOC = $(PROJECT_ROOT)/nextlinux-api-swagger-$(ENTERPRISE_REF).yaml

define generate_openapi_client
	# remove previous API clients
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo rm -rf ${3}; else rm -rf ${3}; fi

	# generate the API client
	docker run \
		--rm \
		-v $${PWD}:/local \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
			generate \
				-c /local/generator.yaml \
				--additional-properties packageName=${2} \
				--additional-properties packageVersion=$(CLIENT_VERSION) \
				--http-user-agent "nextlinux-client/$(CLIENT_VERSION)/go" \
				-i /local/${1} \
				-o /local/${3} \
				-t /local/templates \
				-g go

	# keep permissions the same as the user
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo chown -R $(shell id -u):$(shell id -g) ${3}; fi

	# remove unused files (go.mod is curated manually)
	rm ${3}/{.travis.yml,git_push.sh,go.*}
endef

.PHONY :=
.DEFAULT_GOAL :=
update: clean generate ## pull all swagger definitions and generate client code

.PHONY :=
generate: generate-clients patch ## generate all client code from all swagger documents

define clone
	if [ ! -d "./${CLONE_DIR}" ]; then git clone git@github.com:nextlinux/enterprise.git $(CLONE_DIR); fi
	if [ -d "./${CLONE_DIR}" ]; then cd ${CLONE_DIR} && git fetch; fi
	cd ${CLONE_DIR} && git checkout ${ENTERPRISE_REF}
endef

$(PROJECT_ROOT):
	mkdir -p $(PROJECT_ROOT)

$(ENTERPRISE_OPENAPI_DOC): $(PROJECT_ROOT) ## pull the enterprise external API swagger document
	$(call clone)
	# note: the existing upstream swagger document needs to be corrected, otherwise invalid code will be generated.
	# the tr/sed cmds are a workaround for now.
	cp $(CLONE_DIR)/nextlinux_enterprise/services/api/swagger/nextlinux_api_swagger.yaml $(ENTERPRISE_OPENAPI_DOC)

.PHONY :=
generate-clients: $(ENTERPRISE_OPENAPI_DOC) ## generate client code for nextlinux external API
	$(call generate_openapi_client,$(ENTERPRISE_OPENAPI_DOC),enterprise,$(ENTERPRISE_ROOT))
	# add any tailored code via go generate
	go generate .

.PHONY :=
clean: ## remove all swagger documents and generated client code
	rm -rf $(PROJECT_ROOT)

.PHONY :=
patch: ## applies any post-generation patches
	git apply patches/compliance_binary_properties.patch

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
