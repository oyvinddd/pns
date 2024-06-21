build_pns:
	go build -o bin/pns.out ./cmd/pns

run_pns:
	go run ./cmd/pns

build_apns:
	go build -o bin/apns.out ./cmd/apns

run_apns:
	go run ./cmd/apns

build_all: build_pns build_apns