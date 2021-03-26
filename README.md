# go-whosonfirst-grpc-sqlite

SQLite support for the gRPC implementation of the `go-whosonfirst-spatial` interfaces.

## Important

This is work in progress. Documentation to follow.

## Tools

### server

```
$> ./bin/server -h
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -host string
    	The host to listen for requests on (default "localhost")
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/emitter URI. Supported schemes are: directory://, featurecollection://, file://, filelist://, geojsonl://, repo://. (default "repo://")
  -port int
    	The port to listen for requests on (default 8082)
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [file:// fs:// null://]
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [sqlite://]
  -verbose
    	Be chatty.
```

For example:

```
$> ./bin/server -spatial-database-uri 'sqlite://?dsn=/usr/local/data/yugoslavia.db'
2021/03/26 09:03:05 Listening on localhost:8082
2021/03/26 09:03:05 time to index paths (0) 18.661µs
09:03:05.895116 [server] STATUS finished indexing in 209.866µs
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial
* https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite
* https://github.com/whosonfirst/go-whosonfirst-spatial-grpc
* https://github.com/grpc/grpc-go