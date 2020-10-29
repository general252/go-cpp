
all:
	go build -o ./bin/go-cpp.exe

lin:
	go env -w GOPROXY=https://goproxy.io,direct
	go build -o ./bin/go-cpp.exe
	export LD_LIBRARY_PATH=./bin
	./bin/go-cpp.exe

run:
	./bin/go-cpp.exe
