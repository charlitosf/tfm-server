run: server.exe
	.\server.exe

server.exe: cmd/server/main.go pkg/*/*.go internal/*/*.go
	clear
	go build -o server.exe cmd\server\main.go