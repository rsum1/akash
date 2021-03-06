$(CACHE):
	@echo "creating .cache dir structure..."
	mkdir -p $@
	mkdir -p $(CACHE_BIN)
	mkdir -p $(CACHE_INCLUDE)
	mkdir -p $(CACHE_VERSIONS)

$(GOLANGCI_LINT_VERSION_FILE): $(CACHE)
	@echo "installing golangci-lint..."
	rm -f $(GOLANGCI_LINT)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(CACHE_BIN) $(GOLANGCI_LINT_VERSION)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(GOLANGCI_LINT): $(GOLANGCI_LINT_VERSION_FILE)

.PHONY:lintdeps-install
lintdeps-install: $(GOLANGCI_LINT)
	@echo "lintdeps-install is deprecated and will be removed once Github Actions migrated to use .cache/bin/golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION)

$(BUF_VERSION_FILE): $(CACHE)
	@echo "installing protoc buf cli..."
	rm -f $(BUF)
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(BUF): $(BUF_VERSION_FILE)

$(PROTOC_VERSION_FILE): $(CACHE)
	@echo "installing protoc compiler..."
	rm -f $(PROTOC)
	(cd /tmp; \
	curl -sOL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) bin/protoc; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) 'include/*'; \
	rm -f ${PROTOC_ZIP})
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC): $(PROTOC_VERSION_FILE)

$(MODVENDOR): $(CACHE)
	@echo "installing modvendor..."
	GOBIN=$(CACHE_BIN) GO111MODULE=off go get github.com/goware/modvendor

clang-format-install:
ifeq (, $(shell which ${CLANG_FORMAT_BIN}))
	@echo "Installing ${CLANG_FORMAT_BIN}..."
ifeq ($(UNAME_OS),Darwin)
	curl https://gist.githubusercontent.com/bvigueras/daf11aee6876fb9ba4c925c2c31bc04b/raw/ \
	526ff0eebbc0476f568c852a8cc5d4cc48281475/clang-format@6.rb -o \
	$(brew --repo)/Library/Taps/homebrew/homebrew-core/Formula/clang-format@6.rb
	brew install clang-format@6
endif
ifeq ($(UNAME_OS),Linux)
	if [ -e /etc/debian_version ]; then \
		sudo apt-get install -y ${CLANG_FORMAT_BIN} ; \
	elif [ -e /etc/fedora-release ]; then \
		sudo dnf install clang; \
	else \
		echo -e "\tRun (as root): subscription-manager repos --enable rhel-7-server-devtools-rpms ; \
		yum install llvm-toolset-7" >&2; \
	fi;
endif
else
	@echo "${CLANG_FORMAT_BIN} already installed; skipping..."
endif

kubetypes-deps-install:
	if [ -d "$(shell go env GOPATH)/src/k8s.io/code-generator" ]; then    \
		cd "$(shell go env GOPATH)/src/k8s.io/code-generator" && git pull;  \
		exit 0;                                                             \
	fi;                                                                   \
	mkdir -p "$(shell go env GOPATH)/src/k8s.io" && \
	git clone  git@github.com:kubernetes/code-generator.git \
		"$(shell go env GOPATH)/src/k8s.io/code-generator"

devdeps-install: $(GOLANGCI_LINT) kubetypes-deps-install
	$(GO) install github.com/vektra/mockery/.../
	$(GO) install k8s.io/code-generator/...
	$(GO) install sigs.k8s.io/kind
	$(GO) install golang.org/x/tools/cmd/stringer

cache-clean:
	rm -rf $(CACHE)
