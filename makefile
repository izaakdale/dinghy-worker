
run: 
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7778 \
	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7778 \
	CLUSTER_ADDR=127.0.0.1 \
	CLUSTER_PORT=7777 \
	NAME=worker \
	DATA_DIR=node_data \
	RAFT_ADDR=127.0.0.1 \
	RAFT_PORT=8888 \
	go run .

docker:
	docker build -t dinghy-worker .

up:
	kubectl apply -k deploy/
dn:
	kubectl delete -k deploy/