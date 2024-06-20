unit-tests:
	@go test ./tests/unit/... --tags=unit -v

unit-tests-cover:
	@go test ./tests/unit/... -coverpkg ./internal/... --tags=unit -v

unit-tests-report:
	mkdir -p "coverage" \
	&& go test ./tests/unit/... -v -coverprofile=coverage/cover.out -coverpkg ./internal/... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

.PHONY: run-app,
		unit-tests,
		unit-tests-cover,
		unit-tests-report,