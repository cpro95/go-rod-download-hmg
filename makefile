BINARY_NAME=go-rod-download-hmg

run:
	go run *.go

build:
	go build -o ${BINARY_NAME} *.go

clean:
	go clean
	rm -f ${BINARY_NAME}

