
run1: 
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7778 \
	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7778 \
	CLUSTER_ADDR=127.0.0.1 \
	CLUSTER_PORT=7777 \
	NAME=worker1 \
	DATA_DIR=node_data_1 \
	RAFT_ADDR=127.0.0.1 \
	RAFT_PORT=8888 \
	GRPC_ADDR=127.0.0.1 \
	GRPC_PORT=5001 \
	go run .
run2: 
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7779 \
	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7779 \
	CLUSTER_ADDR=127.0.0.1 \
	CLUSTER_PORT=7777 \
	NAME=worker2 \
	DATA_DIR=node_data_2 \
	RAFT_ADDR=127.0.0.1 \
	RAFT_PORT=8889 \
	GRPC_ADDR=127.0.0.1 \
	GRPC_PORT=5002 \
	go run .

docker:
	docker build -t dinghy-worker .

up:
	kubectl apply -k deploy/ -n default
dn:
	kubectl delete -k deploy/ -n default

.PHONY: gproto
gproto:
	protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--proto_path=.