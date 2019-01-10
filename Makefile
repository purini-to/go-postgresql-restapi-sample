COVER_PROFILE_PATH=
COVER_PROFILE=coverage.out
COVER_MODE=atomic
TEST_FILE_PATH=
TEST_FILE=test.out
REPORT_FILE_PATH=reports/go/
REPORT_FILE=results.xml
BUILD_FILE_PATH=bin/
BUILD_FILE=stack_code_cov
TESTS_PATH=.tests/

.PHONY: default run test report clean build

default: run

clean:
	rm -rf $(TESTS_PATH) $(BUILD_FILE_PATH)

run:
	go run main.go

test: clean
	mkdir -p $(TESTS_PATH)$(COVER_PROFILE_PATH)
	go test -v -cover -race `go list ./... | grep -v /vendor/` -coverprofile=$(TESTS_PATH)$(COVER_PROFILE_PATH)$(COVER_PROFILE) -covermode=$(COVER_MODE) | tee $(TESTS_PATH)$(TEST_FILE_PATH)$(TEST_FILE)

report: test
	mkdir -p $(TESTS_PATH)$(REPORT_FILE_PATH)
	go-junit-report --set-exit-code < $(TESTS_PATH)$(TEST_FILE_PATH)$(TEST_FILE) > $(TESTS_PATH)$(REPORT_FILE_PATH)$(REPORT_FILE)

build: clean
	go build -race -ldflags "-extldflags '-static'" -o $(BUILD_FILE_PATH)$(BUILD_FILE)
