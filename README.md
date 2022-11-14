# go
- install go using https://go.dev/doc/install
- install vs code
- install go extension
- using a command line install packages
  - go get github.com/go-kit/kit/endpoint
  - go get github.com/go-kit/kit/transport/http
  
##Database
 - github.com/lib/pq
 - added "host  all  all 0.0.0.0/0 md5" to "C:\Program Files\PostgreSQL\15\data\pg_hba.conf" for testing purposes as appplication connects to external IP address server

 ##Testing
 - "go test" (in the current directory)

##Docker (all commands in the root dir)
 - "docker build . -t users"
 - "docker run -p 12345:12345 users"