ROX_PROJECT=apollo
TESTFLAGS=-race -p 4
BASE_DIR=$(CURDIR)
TAG=$(shell git describe --tags --abbrev=10 --dirty --long)

RELEASE_GOTAGS := release
ifdef CI
ifneq ($(CIRCLE_TAG),)
GOTAGS := $(RELEASE_GOTAGS)
TAG := $(CIRCLE_TAG)
endif
endif

null :=
space := $(null) $(null)
comma := ,

FORMATTING_FILES=$(shell git grep -L '^// Code generated by .* DO NOT EDIT\.' -- '*.go')

.PHONY: all
all: deps style test image

###########
## Style ##
###########
.PHONY: style
style: fmt imports lint vet roxvet blanks validateimports no-large-files storage-protos-compatible ui-lint qa-tests-style

# staticcheck is useful, but extremely computationally intensive on some people's machines.
# Therefore, to allow people to continue running `make style`, staticcheck is not run along with
# the other style targets by default, when running locally.
# It is always run in CI.
# To run it locally along with the other style targets, you can `export RUN_STATIC_CHECK=true`.
# If you want to run just staticcheck, you can, of course, just `make staticcheck`.
ifdef CI
style: staticcheck
endif

ifdef RUN_STATIC_CHECK
style: staticcheck
endif

.PHONY: qa-tests-style
qa-tests-style:
	@echo "+ $@"
	make -C qa-tests-backend/ style

.PHONY: ui-lint
ui-lint:
	@echo "+ $@"
	make -C ui lint

STATICCHECK_BIN := $(GOPATH)/bin/staticcheck
$(STATICCHECK_BIN):
	@echo "+ $@"
	@go get honnef.co/go/tools/cmd/staticcheck

.PHONY: staticcheck
staticcheck: $(STATICCHECK_BIN)
	@echo "+ $@"
	@$(BASE_DIR)/tools/staticcheck-wrap.sh ./...

.PHONY: fmt
fmt:
	@echo "+ $@"
ifdef CI
		@echo "The environment indicates we are in CI; checking gofmt."
		@echo 'If this fails, run `make style`.'
		@$(eval FMT=`echo $(FORMATTING_FILES) | xargs gofmt -s -l`)
		@echo "gofmt problems in the following files, if any:"
		@echo $(FMT)
		@test -z "$(FMT)"
endif
	@echo $(FORMATTING_FILES) | xargs gofmt -s -l -w

.PHONY: imports
imports: deps volatile-generated-srcs
	@echo "+ $@"
ifdef CI
		@echo "The environment indicates we are in CI; checking goimports."
		@echo 'If this fails, run `make style`.'
		@$(eval IMPORTS=`echo $(FORMATTING_FILES) | xargs goimports -l`)
		@echo "goimports problems in the following files, if any:"
		@echo $(IMPORTS)
		@test -z "$(IMPORTS)"
endif
	@echo $(FORMATTING_FILES) | xargs goimports -w

.PHONY: validateimports
validateimports:
	@echo "+ $@"
	@go run $(BASE_DIR)/tools/validateimports/verify.go $(shell go list -e ./...)

.PHONY: no-large-files
no-large-files:
	@echo "+ $@"
	@$(BASE_DIR)/tools/large-git-files/find.sh

.PHONY: roxvet
roxvet:
	@echo "+ $@"
	@go install $(BASE_DIR)/tools/roxvet
	@go vet -vettool "$$(go env GOPATH)/bin/roxvet" $(shell go list -e ./... | grep -v -e 'stackrox/rox/image')

.PHONY: keys
keys:
	@echo "+ $@"
	go generate github.com/stackrox/rox/central/ed

PROTOLOCK_BIN := $(GOPATH)/bin/protolock
$(PROTOLOCK_BIN):
	@echo "+ $@"
	$(BASE_PATH)/scripts/go-get-version.sh github.com/viswajithiii/protolock 43bb8a9ba4e8de043a5ffacc64b1c38d95419e1d --skip-install
	mkdir -p $(GOPATH)/src/github.com/nilslice
	mv $(GOPATH)/src/github.com/viswajithiii/protolock $(GOPATH)/src/github.com/nilslice/protolock
	go install github.com/nilslice/protolock/...

.PHONY: storage-protos-compatible
storage-protos-compatible: $(PROTOLOCK_BIN)
	@echo "+ $@"
	@protolock status -lockdir=$(BASE_DIR)/proto/storage -protoroot=$(BASE_DIR)/proto/storage

.PHONY: lint
lint:
	@echo "+ $@"
	@$(BASE_DIR)/tools/go-lint.sh $(FORMATTING_FILES)

.PHONY: vet-active-tags
vet-active-tags: deps volatile-generated-srcs
	@echo "+ $@"
	@$(BASE_DIR)/tools/go-vet.sh -tags "$(subst $(comma),$(space),$(GOTAGS))" $(shell go list -e ./... | grep -v generated | grep -v vendor)

.PHONY: vet
vet: vet-active-tags
ifdef CI
	@echo "+ $@ ($(RELEASE_GOTAGS))"
	@$(BASE_DIR)/tools/go-vet.sh -tags "$(subst $(comma),$(space),$(RELEASE_GOTAGS))" $(shell go list -e ./... | grep -v generated | grep -v vendor)
endif

.PHONY: blanks
blanks:
	@echo "+ $@"
ifdef CI
	@find . \( \( -name vendor -o -name generated \) -type d -prune \) -o \( -name \*.go -print0 \) | xargs -0 $(BASE_PATH)/tools/import_validate.py
else
	@find . \( \( -name vendor -o -name generated \) -type d -prune \) -o \( -name \*.go -print0 \) | xargs -0 $(BASE_PATH)/tools/fix-blanks.sh
endif

.PHONY: dev
dev:
	@echo "+ $@"
	@go get -u golang.org/x/lint/golint
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/jstemmer/go-junit-report
	@go get -u github.com/golang/dep/cmd/dep
	@go install ./tools/roxvet


#####################################
## Generated Code and Dependencies ##
#####################################

PROTO_GENERATED_SRCS = $(GENERATED_PB_SRCS) $(GENERATED_API_GW_SRCS)

include make/protogen.mk

STRINGER_BIN := $(GOPATH)/bin/stringer
$(STRINGER_BIN):
	@echo "+ $@"
	@go get golang.org/x/tools/cmd/stringer

MOCKGEN_BIN := $(GOPATH)/bin/mockgen
$(MOCKGEN_BIN):
	@echo "+ $@"
	@$(BASE_PATH)/scripts/go-get-version.sh golang.org/x/tools e21233ffa6c386fc230b4328493f77af54ff9372 --skip-install
	@$(BASE_PATH)/scripts/go-get-version.sh github.com/golang/mock/mockgen 8a44ef6e8be577e050008c7886f24fc705d709fb

GENNY_BIN := $(GOPATH)/bin/genny
$(GENNY_BIN):
	@echo "+ $@"
	@$(BASE_PATH)/scripts/go-get-version.sh github.com/mauricelam/genny e937528877485c089aa62cfa9f60968749d650f1

PACKR_BIN := $(GOPATH)/bin/packr
$(PACKR_BIN):
	@echo "+ $@"
	@$(BASE_PATH)/scripts/go-get-version.sh github.com/gobuffalo/packr/packr 899fe0e4176fca9bca81763c810d74af82548c78

.PHONY: go-packr-srcs
go-packr-srcs: $(PACKR_BIN)
	@echo "+ $@"
	@packr

# For some reasons, a `packr clean` is much slower than the `find`. It also does not work.
.PHONY: clean-packr-srcs
clean-packr-srcs:
	@echo "+ $@"
	@find . -name '*-packr.go' -exec rm {} \;

EASYJSON_BIN := $(GOPATH)/bin/easyjson
$(EASYJSON_BIN):
	@echo "+ $@"
	@$(BASE_PATH)/scripts/go-get-version.sh github.com/mailru/easyjson/easyjson 60711f1a8329503b04e1c88535f419d0bb440bff

.PHONY: go-easyjson-srcs
go-easyjson-srcs: $(EASYJSON_BIN)
	@echo "+ $@"
	@easyjson -pkg pkg/docker/types/types.go
	@echo "//lint:file-ignore SA4006 This is a generated file" >> pkg/docker/types/types_easyjson.go
	@easyjson -pkg pkg/docker/types/container.go
	@echo "//lint:file-ignore SA4006 This is a generated file" >> pkg/docker/types/container_easyjson.go
	@easyjson -pkg pkg/docker/types/image.go
	@echo "//lint:file-ignore SA4006 This is a generated file" >> pkg/docker/types/image_easyjson.go

.PHONY: clean-easyjson-srcs
clean-easyjson-srcs:
	@echo "+ $@"
	@find . -name '*_easyjson.go' -exec rm {} \;

.PHONY: go-generated-srcs
go-generated-srcs: deps go-easyjson-srcs $(MOCKGEN_BIN) $(STRINGER_BIN) $(GENNY_BIN)
	@echo "+ $@"
	PATH=$(PATH):$(BASE_DIR)/tools/generate-helpers go generate -v ./...

proto-generated-srcs: $(PROTO_GENERATED_SRCS)
	@echo "+ $@"
	@touch proto-generated-srcs

clean-proto-generated-srcs:
	@echo "+ $@"
	git clean -xdf generated

# volatile-generated-srcs are all generated sources that are NOT committed
.PHONY: volatile-generated-srcs
volatile-generated-srcs: proto-generated-srcs go-packr-srcs keys

.PHONY: generated-srcs
generated-srcs: volatile-generated-srcs go-generated-srcs

# clean-generated-srcs cleans ONLY volatile-generated-srcs.
.PHONY: clean-generated-srcs
clean-generated-srcs: clean-packr-srcs clean-proto-generated-srcs
	@echo "+ $@"

deps: Gopkg.toml Gopkg.lock proto-generated-srcs
	@echo "+ $@"
ifdef CI
	@# `dep check` exits with a nonzero code if there is a toml->lock mismatch.
	dep check -skip-vendor
endif
	@$(eval GOMOCK_REFLECT_DIRS=`find . -type d -name 'gomock_reflect_*'`)
	@test -z $(GOMOCK_REFLECT_DIRS) || { echo "Found leftover gomock directories. Please remove them and rerun make deps!"; echo $(GOMOCK_REFLECT_DIRS); exit 1; }
	@# `dep ensure` can be flaky sometimes, so try rerunning it if it fails.
	dep ensure || (rm -rf vendor .vendor-new && dep ensure)
	@touch deps

.PHONY: clean-deps
clean-deps:
	@echo "+ $@"
	@rm -f deps

###########
## Build ##
###########
PURE := --features=pure
RACE := --features=race
LINUX_AMD64 := --cpu=k8
VARIABLE_STAMPS := --workspace_status_command=$(BASE_DIR)/status.sh
BAZEL_OS=linux
ifeq ($(UNAME_S),Darwin)
    BAZEL_OS=darwin
endif
PLATFORMS := --platforms=@io_bazel_rules_go//go/toolchain:$(BAZEL_OS)_amd64

BAZEL_BASE_FLAGS := $(PURE) $(VARIABLE_STAMPS) --define gotags=$(GOTAGS)
BAZEL_FLAGS := $(BAZEL_BASE_FLAGS) $(LINUX_AMD64)

cleanup:
	@echo "Total BUILD.bazel files deleted: "
	@git status --ignored --untracked-files=all --porcelain | grep '^\(!!\|??\) ' | cut -d' ' -f 2- | grep '\(/\|^\)BUILD\.bazel$$' | xargs rm -v | wc -l

.PHONY: gazelle
gazelle: deps volatile-generated-srcs cleanup
	bazel run //:gazelle -- -build_tags=$(GOTAGS)

cli: gazelle
ifdef CI
	bazel build $(BAZEL_FLAGS) --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 -- //roxctl
	bazel build $(BAZEL_FLAGS) --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -- //roxctl
	bazel build $(BAZEL_FLAGS) --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 -- //roxctl
else
	bazel build $(BAZEL_FLAGS) --platforms=@io_bazel_rules_go//go/toolchain:$(BAZEL_OS)_amd64 -- //roxctl
endif
	# Copy the user's specific OS into gopath
	cp bazel-bin/roxctl/$(BAZEL_OS)_amd64_pure_stripped/roxctl $(GOPATH)/bin/roxctl
	chmod u+w $(GOPATH)/bin/roxctl

.PHONY: main-build
main-build: gazelle
	@echo "+ $@"
	@# PLEASE KEEP THE TWO LISTS BELOW IN SYNC.
	@# The only exception is that `roxctl` should not be built in CI here, since it's built separately when in CI.
	@# This isn't pretty, but it saves 30 seconds on every build, which seems worth it.
ifdef CI
	bazel build $(BAZEL_FLAGS) //central //migrator //sensor/kubernetes //compliance/collection
else
	bazel build $(BAZEL_FLAGS) //central //migrator //sensor/kubernetes //compliance/collection //roxctl
endif

.PHONY: scale-build
scale-build: gazelle
	@echo "+ $@"
	bazel build $(BAZEL_FLAGS) \
		//scale/mocksensor \
		//scale/mockcollector \
		//scale/profiler

.PHONY: webhookserver-build
webhookserver-build: gazelle
	@echo "+ $@"
	bazel build $(BAZEL_FLAGS) \
		//webhookserver

.PHONY: mock-grpc-server-build
mock-grpc-server-build: gazelle
	@echo "+ $@"
	bazel build $(BAZEL_FLAGS) \
		//integration-tests/mock-grpc-server

.PHONY: gendocs
gendocs: $(GENERATED_API_DOCS)
	@echo "+ $@"

# We don't need to do anything here, because the $(MERGED_API_SWAGGER_SPEC) target already performs validation.
.PHONY: swagger-docs
swagger-docs: $(MERGED_API_SWAGGER_SPEC)
	@echo "+ $@"

BAZEL_TEST_PATTERNS ?= //...

.PHONY: bazel-test
bazel-test: gazelle
	-rm vendor/github.com/coreos/pkg/BUILD
	-rm vendor/github.com/cloudflare/cfssl/script/BUILD
	-rm vendor/github.com/grpc-ecosystem/grpc-gateway/BUILD
	@# Be careful if you add action_env arguments; their values can invalidate cached
	@# test results. See https://github.com/bazelbuild/bazel/issues/2574#issuecomment-320006871.
	bazel coverage $(BAZEL_BASE_FLAGS) $(RACE) \
	    --instrumentation_filter=//... \
	    --test_output=errors \
	    -- \
	    $(BAZEL_TEST_PATTERNS) -proto/... -tests/... -vendor/...


.PHONY: ui-build
ui-build:
ifdef SKIP_UI_BUILD
	test -d ui/build || make -C ui build
else
	make -C ui build
endif

.PHONY: ui-test
ui-test:
	@# UI tests don't work in Bazel yet.
	make -C ui test

.PHONY: test
test: bazel-test ui-test

.PHONY: integration-unit-tests
integration-unit-tests: gazelle
	 go test -tags=integration $(shell go list ./... | grep  "registries\|scanners\|notifiers")

upload-coverage:
	@# 'mode: set' is repeated in each coverage file, but Coveralls only wants it
	@# exactly once at the head of the file.
	@# We might be able to use Coveralls parallel builds to resolve this:
	@#     https://docs.coveralls.io/parallel-build-webhook

	@echo 'mode: set' > combined_coverage.dat
	@find ./bazel-testlogs/ -name 'coverage.dat' | xargs -I {} cat "{}" | grep -v 'mode: set' | grep -v vendor >> combined_coverage.dat
	goveralls -coverprofile="combined_coverage.dat" -ignore 'central/graphql/resolvers/generated.go,generated/storage/*,generated/*/*/*' -service=circle-ci -repotoken="$$COVERALLS_REPO_TOKEN"

.PHONY: coverage
coverage:
	@echo "+ $@"
	@go test -cover -coverprofile coverage.out $(shell go list -e ./... | grep -v /tests)
	@go tool cover -html=coverage.out -o coverage.html

###########
## Image ##
###########

# Exists for compatibility reasons. Please consider migrating to using `make main-image`.
.PHONY: image
image: main-image monitoring-image deployer-image

.PHONY: monitoring-image
monitoring-image:
	docker build -t stackrox/monitoring:$(TAG) monitoring

.PHONY: all-builds
all-builds: cli main-build clean-image $(MERGED_API_SWAGGER_SPEC) ui-build

.PHONY: main-image
main-image: all-builds
	make docker-build-main-image

.PHONY: main-image-rhel
main-image-rhel: all-builds
	make docker-build-main-image-rhel

.PHONY: deployer-image
deployer-image: gazelle
	bazel build $(BAZEL_FLAGS) //roxctl
	make docker-build-deployer-image

# The following targets copy compiled artifacts into the expected locations and
# runs the docker build.
# Please DO NOT invoke this target directly unless you know what you're doing;
# you probably want to run `make main-image`. This target is only in Make for convenience;
# it assumes the caller has taken care of the dependencies, and does not
# declare its dependencies explicitly.
.PHONY: docker-build-main-image
docker-build-main-image: copy-binaries-to-image-dir docker-build-data-image
	docker build -t stackrox/main:$(TAG) --build-arg DATA_IMAGE_TAG=$(TAG) image/
	@echo "Built main image with tag: $(TAG)"
	@echo "You may wish to:       export MAIN_IMAGE_TAG=$(TAG)"

.PHONY: docker-build-main-image-rhel
docker-build-main-image-rhel: copy-binaries-to-image-dir docker-build-data-image
	docker build -t stackrox/main-rhel:$(TAG) --file image/Dockerfile_rhel --label version=$(TAG) --label release=$(TAG) --build-arg DATA_IMAGE_TAG=$(TAG) image/
	@echo "Built main image for RHEL with tag: $(TAG)"
	@echo "You may wish to:       export MAIN_IMAGE_TAG=$(TAG)"

.PHONY: docker-build-data-image
docker-build-data-image:
	test -f $(CURDIR)/image/keys/data-key
	test -f $(CURDIR)/image/keys/data-iv
	docker build -t stackrox-data:$(TAG) image/ --file image/stackrox-data.Dockerfile

.PHONY: docker-build-deployer-image
docker-build-deployer-image:
	cp -f bazel-bin/roxctl/linux_amd64_pure_stripped/roxctl image/bin/roxctl-linux
	docker build -t stackrox/deployer:$(TAG) --build-arg MAIN_IMAGE_TAG=$(TAG) --build-arg SCANNER_IMAGE_TAG=$(shell cat SCANNER_VERSION) image/ --file image/Dockerfile_gcp

.PHONY: copy-binaries-to-image-dir
copy-binaries-to-image-dir:
	cp -r ui/build image/ui/
	cp bazel-bin/central/linux_amd64_pure_stripped/central image/bin/central
ifdef CI
	cp bazel-bin/roxctl/linux_amd64_pure_stripped/roxctl image/bin/roxctl-linux
	cp bazel-bin/roxctl/darwin_amd64_pure_stripped/roxctl image/bin/roxctl-darwin
	cp bazel-bin/roxctl/windows_amd64_pure_stripped/roxctl.exe image/bin/roxctl-windows.exe
else
ifneq ($(BAZEL_OS),linux)
	cp bazel-bin/roxctl/linux_amd64_pure_stripped/roxctl image/bin/roxctl-linux
endif
	cp bazel-bin/roxctl/$(BAZEL_OS)_amd64_pure_stripped/roxctl image/bin/roxctl-$(BAZEL_OS)
endif
	cp bazel-bin/migrator/linux_amd64_pure_stripped/migrator image/bin/migrator
	cp bazel-bin/sensor/kubernetes/linux_amd64_pure_stripped/kubernetes image/bin/kubernetes-sensor
	cp bazel-bin/compliance/collection/linux_amd64_pure_stripped/collection image/bin/compliance

ifdef CI
	@[ -d image/THIRD_PARTY_NOTICES ] || { echo "image/THIRD_PARTY_NOTICES dir not found! It is required for CI-built images."; exit 1; }
else
	@[ -f image/THIRD_PARTY_NOTICES ] || mkdir -p image/THIRD_PARTY_NOTICES
endif
	@[ -d image/docs ] || { echo "Generated docs not found in image/docs. They are required for build."; exit 1; }

.PHONY: scale-image
scale-image: scale-build clean-image
	cp bazel-bin/scale/mocksensor/linux_amd64_pure_stripped/mocksensor scale/image/bin/mocksensor
	cp bazel-bin/scale/mockcollector/linux_amd64_pure_stripped/mockcollector scale/image/bin/mockcollector
	cp bazel-bin/scale/profiler/linux_amd64_pure_stripped/profiler scale/image/bin/profiler
	chmod +w scale/image/bin/*
	docker build -t stackrox/scale:$(TAG) -f scale/image/Dockerfile scale

webhookserver-image: webhookserver-build
	-mkdir webhookserver/bin
	cp bazel-bin/webhookserver/linux_amd64_pure_stripped/webhookserver webhookserver/bin/webhookserver
	chmod +w webhookserver/bin/webhookserver
	docker build -t stackrox/webhookserver:1.1 -f webhookserver/Dockerfile webhookserver

.PHONY: mock-grpc-server-image
mock-grpc-server-image: mock-grpc-server-build clean-image
	cp bazel-bin/integration-tests/mock-grpc-server/linux_amd64_pure_stripped/mock-grpc-server integration-tests/mock-grpc-server/image/bin/mock-grpc-server
	docker build -t stackrox/grpc-server:$(TAG) integration-tests/mock-grpc-server/image

###########
## Clean ##
###########
.PHONY: clean
clean: clean-image clean-generated-srcs
	@echo "+ $@"

.PHONY: clean-image
clean-image:
	@echo "+ $@"
	git clean -xf image/bin
	git clean -xdf image/ui image/docs
	git clean -xf integration-tests/mock-grpc-server/image/bin/mock-grpc-server

.PHONY: tag
tag:
ifdef COMMIT
	@git describe $(COMMIT) --tags --abbrev=10 --long
else
	@echo $(TAG)
endif

ossls-audit:
	ossls version
	ossls audit

ossls-notice:
	ossls version
	ossls audit --export image/THIRD_PARTY_NOTICES

.PHONY: collector-tag
collector-tag:
	@cat COLLECTOR_VERSION
