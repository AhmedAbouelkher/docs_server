clear_tmp:
	rm -rf tmp/*

build:
	go build -o ./exec/server *.go

run: build
	./exec/server

watch:
	clear
	ulimit -n 1000
	reflex -s -r '\.go$$' make run
