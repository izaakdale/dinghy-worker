
run: 
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7778 \
	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7778 \
	CLUSTER_ADDR=127.0.0.1 \
	CLUSTER_PORT=7777 \
	NAME=worker \
	go run .

docker:
	docker build -t dinghy-worker .