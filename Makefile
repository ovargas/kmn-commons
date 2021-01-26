.EXPORT_ALL_VARIABLES:
GO111MODULE = on

test-deps:
	@echo "installing test dependencies..."
	@go get github.com/axw/gocov/...
	@go get github.com/AlekSi/gocov-xml
	@go get github.com/matm/gocov-html

test: test-deps
	@mkdir -p test-artifacts/coverage
	@gocov test ./... -covermode=atomic -coverpkg=./... > test-artifacts/gocov.json
	@cat test-artifacts/gocov.json | gocov report
	@cat test-artifacts/gocov.json | gocov-xml > test-artifacts/coverage/coverage.xml
	@cat test-artifacts/gocov.json | gocov-html > test-artifacts/coverage/coverage.html
	@cat test-artifacts/gocov.json | gocov-html > test-artifacts/coverage/coverage.html