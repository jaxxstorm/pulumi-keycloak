PROJECT_NAME := Keycloak Package
include build/common.mk

PACK             := keycloak
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-${PACK}
NODE_MODULE_NAME := @pulumi/${PACK}
TF_NAME          := ${PACK}

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell scripts/get-version)
PYPI_VERSION    := $(shell cd scripts && ./get-py-version)
LATEST_RESOURCE_PROVIDER_VERSION := $(shell curl --silent "https://api.github.com/repos/pulumi/pulumi-${PACK}/tags" | jq ".[0]".name -r)
PROVIDER_VERSION := ${LATEST_RESOURCE_PROVIDER_VERSION:v%=%}

DOTNET_PREFIX  := $(firstword $(subst -, ,${VERSION:v%=%})) # e.g. 1.5.0
DOTNET_SUFFIX  := $(word 2,$(subst -, ,${VERSION:v%=%}))    # e.g. alpha.1

ifeq ($(strip ${DOTNET_SUFFIX}),)
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})
else
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})-$(strip ${DOTNET_SUFFIX})
endif

TESTPARALLELISM := 4

OS := $(shell uname)
EMPTY_TO_AVOID_SED := ""

# NOTE: Since the plugin is published using the nodejs style semver version
# We set the PLUGIN_VERSION to be the same as the version we use when building
# the provider (e.g. x.y.z-dev-... instead of x.y.zdev...)
build:: install_plugins provider
	cd provider && for LANGUAGE in "nodejs" "python" "go" "dotnet" ; do \
		$(TFGEN) $$LANGUAGE --overlays overlays/$$LANGUAGE/ --out ../${PACKDIR}/$$LANGUAGE/ || exit 3 ; \
	done
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json
	cd ${PACKDIR}/python/ && \
		cp ../../README.md . && \
		$(PYTHON) setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && $(PYTHON) setup.py build sdist
	cd ${PACKDIR}/dotnet/ && \
		echo "${VERSION:v%=%}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

tfgen::
	cd provider && go install -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/v2/pkg/version.Version=${VERSION}" ${PROJECT}/provider/v2/cmd/${TFGEN}

generate_schema:: tfgen
	$(TFGEN) schema --out ./provider/cmd/${PROVIDER}

provider:: generate_schema
	cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go
	cd provider && go install -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/v2/pkg/version.Version=${VERSION}" ${PROJECT}/provider/v2/cmd/${PROVIDER}

lint::
	#golangci-lint run

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource $(PACK) $(PROVIDER_VERSION)

install:: tfgen provider
	[ ! -e "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" ] || rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	mkdir -p "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	cp -r ${PACKDIR}/nodejs/bin/. "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/node_modules"
	cd "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" && \
		yarn install --offline --production && \
		(yarn unlink > /dev/null 2>&1 || true) && \
		yarn link
	cd ${PACKDIR}/python/bin && $(PIP) install --user -e .
	echo "Copying NuGet packages to ${PULUMI_NUGET}"
	[ ! -e "$(PULUMI_NUGET)" ] || rm -rf "$(PULUMI_NUGET)/*"
	find . -name '*.nupkg' -exec cp -p {} ${PULUMI_NUGET} \;

test_fast::
	cd examples && $(GO_TEST_FAST) .

test_all::
	cd examples && $(GO_TEST) .

.PHONY: publish_tgz
publish_tgz:
	$(call STEP_MESSAGE)
	./scripts/publish_tgz.sh

.PHONY: publish_packages
publish_packages:
	$(call STEP_MESSAGE)
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/publish-tfgen-package .
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/build-package-docs.sh ${PACK}

.PHONY: check_clean_worktree
check_clean_worktree:
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/check-worktree-is-clean.sh

# The travis_* targets are entrypoints for CI.
.PHONY: travis_cron travis_push travis_pull_request travis_api
travis_cron: all
travis_push: only_build check_clean_worktree publish_tgz only_test publish_packages
travis_pull_request: all check_clean_worktree
travis_api: all
