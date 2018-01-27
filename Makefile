authDocker: authserver
	sudo docker run --rm -p 8081:8081 authserver

authserver:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o authserver  

