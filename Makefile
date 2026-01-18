.PHONY: test coverage coverage-html clean

test:
	go test -v ./...

coverage:
	go test -v -coverprofile=coverage.out ./...

coverage-html: coverage
	go tool cover -html=coverage.out -o coverage.html

coverage-report: coverage
	go tool cover -func=coverage.out

clean:
	rm -f coverage.out coverage.html
