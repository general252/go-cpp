
all:
	go build -ldflags "-w -s" -o ./bin/go-cpp.exe

lin:
	go env -w GOPROXY=https://goproxy.io,direct
	go build -ldflags "-w -s" -o ./bin/go-cpp.app
	export LD_LIBRARY_PATH=./bin
	./bin/go-cpp.app

run:
	./bin/go-cpp.exe
