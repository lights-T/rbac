GIT_USER_NAME=lights-T
PERSONAL_ACCESS_TOKEN=$(shell echo ${GITHUB_ACCESS_TOKEN})

proto:
	protoc --proto_path=${GOPATH}/src --proto_path=. --gogofaster_out=. --micro_out=. --govalidators_out=gogoimport=true:.  proto/*.proto

mod:
	git config --global url."https://$(GIT_USER_NAME):$(PERSONAL_ACCESS_TOKEN)@github.com".insteadOf "https://github.com"
	go env -w GO111MODULE=on GOPRIVATE=github.com/$(GIT_USER_NAME) GOPROXY=https://goproxy.cn,direct
	go mod tidy

test: mod
	go test -v ./...

clean:
	rm -rf ./$(NAME)


.PHONY: install proto build clean test docker
