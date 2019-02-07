### Steps for starting project
1) git https://github.com/smirnoffV/482.solution-storage-client.git
2) install package manager "dep" to machine
3) set GOPATH
4) in the project folder run a command "make install && make build"
5) to start the client in the project folder run a command "./bin/run -h=127.0.0.1 -p=8080" (-h storage host, -p storage port)

### Available commands
1) SET key value
2) GET key
3) exit
