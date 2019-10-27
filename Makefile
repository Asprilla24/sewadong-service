TAG=sewadong-service
SVC=$(TAG)
BUILDER=sewadong/golang-builder:1.12-alpine3.9
MIGRATION=$(SVC)-migration
export REVISION_ID?=unknown
export BUILD_DATE?=unknown

RUN=docker run --rm \
	-v $(CURDIR):/go/src/$(SVC) \
	-w /go/src/$(SVC)

build:
	docker build -f Dockerfile.build -t $(BUILDER) .
ifeq ($(OS),Windows_NT)
# Workaround on Windows for https://github.com/golang/dep/issues/1407
	$(RUN) $(BUILDER) rm -rf vendor vendor.orig
	$(RUN) $(BUILDER) rm -rf vendor vendor.orig
endif
	$(RUN) $(BUILDER) dep ensure -v
	$(RUN) $(BUILDER) gometalinter.v2 --deadline=1200s --enable-gc --cyclo-over=12 \
		./cmd/... ./pkg/...
	$(RUN) -e CGO_ENABLED=0 -e GOOS=linux $(BUILDER) go build -o ./$(TAG) \
		-ldflags "-s -X main.revisionID=$(REVISION_ID) -X main.buildDate=$(BUILD_DATE)" \
		./cmd/$(TAG)/
	docker build --tag="$(TAG):$(REVISION_ID)" -f Dockerfile .

run:
	docker-compose up

clean:
	-$(RUN) $(BUILDER) go clean -i
	-$(RUN) $(BUILDER) rm -rf vendor vendor.* *.log bin test.xml
	-$(RUN) $(BUILDER) find . -type f -name 'coverage.xml' -delete
	-docker rmi -f $(TAG):$(REVISION_ID) $(BUILDER)

rebuild:
	docker-compose stop service
	$(RUN) -e CGO_ENABLED=0 -e GOOS=linux $(BUILDER) go build -o ./$(TAG) \
		-ldflags "-s -X main.revisionID=$(REVISION_ID) -X main.buildDate=$(BUILD_DATE)" \
		./cmd/$(TAG)/
	docker-compose up service

COMPOSE_TEST=docker-compose --no-ansi -f docker-compose-test.yaml
test:
	$(COMPOSE_TEST) up -d
	$(COMPOSE_TEST) run -e CGO_ENABLED=0 -e PWD=$(CURDIR) test bash coverage.sh
	$(COMPOSE_TEST) stop

deploy:
	heroku container:push web --app=$(TAG)
	heroku container:release web --app=$(TAG)