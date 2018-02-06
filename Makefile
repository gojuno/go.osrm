all: test bench

test:
	go test -v -cover ./...

bench:
	go test -run=- -bench=Benchmark* -benchmem ./...

lint:
	gometalinter --deadline=1m ./...
