all: test bench

test:
	go test -v -race -cover ./...

bench:
	go test -run=- -bench=Benchmark* -benchmem ./...

lint:
	gometalinter --deadline=1m ./...
