# go-whosonfirst-grpc-sqlite

## Important

This is work in progress. Documentation to follow.

## Tools

### server

```
$> ./bin/server -h
  -custom-placetypes string
    	...
  -custom-placetypes-source string
    	...
  -enable-custom-placetypes
    	...
  -enable-properties
    	Enable support for 'properties' parameters in queries.
  -exclude value
    	Exclude (WOF) records based on their existential flags. Valid options are: ceased, deprecated, not-current, superseded.
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -mode string
    	Valid modes are: directory, featurecollection, file, filelist, geojsonl, repo. (default "repo://")
  -properties-reader-uri string
    	Valid options are: [sqlite://]
  -server-address string
    	A valid gRPC server address (default "localhost:8282")
  -setenv
    	Set flags from environment variables.
  -spatial-database-uri string
    	Valid options are: [sqlite://] (default "sqlite://")
  -verbose
    	Be chatty.
```

For example:

```
$> ./bin/server go run \
	-spatial-database-uri 'sqlite:///?dsn=/usr/local/data/sfomuseum-data-architecture.db' \
	-properties-reader-uri 'sqlite:///?dsn=/usr/local/data/sfomuseum-data-architecture.db' \
	-mode directory://
	
2020/12/10 12:40:21 Listening for requests on localhost:8282
12:40:21.258506 [main] STATUS finished indexing in 68.655µs
2020/12/10 12:40:25 COUNT 10 37.61701965332031 -122.38666534423828
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial
* https://github.com/whosonfirst/go-whosonfirst-spatial-grpc
* https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite