
all:
	go build -v -o main .

build:
	GOOS=linux go build -o main .

zip: build
	zip deployment.zip main
