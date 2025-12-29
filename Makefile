# round robin
run-rr:
	@echo "[ Runing... ]"
	@go run ./cmd/roundRobin/main.go

# weighted Round Robin
run-wrr:
	@echo "[ Runing... ]"
	@go run ./cmd/wrr/main.go

# least Connections
run-lb:
	@echo "[ Runing... ]"
	@go run ./cmd/leastconn/main.go

# examples servers
run-servers:
	@echo "[ Runing... ]"
	PORT=8081 NAME=ping-1 go run ./cmd/server/main.go & \
	PORT=8082 NAME=ping-2 go run ./cmd/server/main.go & \
	PORT=8083 NAME=ping-3 go run ./cmd/server/main.go 

