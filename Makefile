cli:
	go build -mod vendor -o bin/server cmd/server/main.go

debug:
	go run -mod vendor cmd/server/main.go -enable-properties -spatial-database-uri 'sqlite:///?dsn=$(DSN)' -properties-reader-uri 'sqlite:///?dsn=$(DSN)' -mode directory:// 
