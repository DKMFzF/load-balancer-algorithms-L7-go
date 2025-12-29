run:
	@echo "[ Runing... ]"
	@go run ./cmd/balancer/main.go

run-round-robin:
	@echo "[ Runing... ]"
	@go run ./cmd/roundRobin/main.go

run-wrr:
	@echo "[ Runing... ]"
	@go run ./cmd/wrr/main.go

run-lb:
	@echo "[ Runing... ]"
	@go run ./cmd/leastconn/main.go

run-servers:
	@echo "[ Runing... ]"
	PORT=8081 NAME=ping-1 go run ./cmd/server/main.go & \
	PORT=8082 NAME=ping-2 go run ./cmd/server/main.go & \
	PORT=8083 NAME=ping-3 go run ./cmd/server/main.go 

