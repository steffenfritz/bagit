test:
	go test -v
	rm -rf testdata/bag_golden

.PHONY: test