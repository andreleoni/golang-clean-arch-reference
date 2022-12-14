mocks:
	docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage
	chown -R $USER:$USER *

test:
	go test ./...
