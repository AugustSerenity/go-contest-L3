run:
	docker-compose up -d
	
vet:
	go vet ./...

lint:
	golangci-lint run ./...

check: vet lint