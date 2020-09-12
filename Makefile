all:
	go build
test:
	go test -v
	rm -rf testdata/bag_golden
prod:
	go build -ldflags "-w -s"
