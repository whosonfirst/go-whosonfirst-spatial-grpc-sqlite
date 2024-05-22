# go-whosonfirst-grpc-sqlite

SQLite support for the gRPC implementation of the `go-whosonfirst-spatial` interfaces.

## Documentation

Documentation is incomplete at this time.

## Tools

### server

```
> ./bin/server -h
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -host string
    	... (default "localhost")
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v2 URI. Supported schemes are: directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://. (default "repo://")
  -port int
    	... (default 8082)
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// repo:// sqlite:// stdin://]. If the value is {spatial-database-uri} then the value of the '-spatial-database-uri' implements the reader.Reader interface and will be used.
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [rtree:// sqlite://] (default "rtree://")
```

For example:

```
$> ./bin/server -spatial-database-uri 'sqlite://?dsn=modernc:///usr/local/data/yugoslavia.db'
2021/03/26 09:03:05 Listening on localhost:8082
2021/03/26 09:03:05 time to index paths (0) 18.661µs
09:03:05.895116 [server] STATUS finished indexing in 209.866µs
```

And then in [whosonfirst/go-whosonfirst-spatial-grpc](https://github.com/whosonfirst/go-whosonfirst-spatial-grpc):

```
$> go build -mod vendor -o bin/client cmd/client/main.go

$> ./bin/client \
	-latitude 37.615761 \
	-longitude -122.389154 \
	-sort-uri placetype:// \
	-sort-uri name:// \
	-sort-uri inception:// \

	| jq '.places[] | "\(."name") \(."inception_date") (\(."placetype"))"'
	
"Earth null (planet)"
"North America null (continent)"
"United States null (country)"
"California null (region)"
"San Mateo null (county)"
"San Francisco International Airport 1948~ (campus)"
"SFO Terminal Complex 2000~ (building)"
"SFO Terminal Complex 2006~ (building)"
"SFO Terminal Complex 2011~ (building)"
"SFO Terminal Complex 2014~ (building)"
"SFO Terminal Complex 2017~ (building)"
"SFO Terminal Complex 2019-07-23 (building)"
"SFO Terminal Complex 2020-~05 (building)"
"SFO Terminal Complex 2021-05-25 (building)"
"SFO Terminal Complex 2021-11-09 (building)"
"International Terminal 2006~ (wing)"
"International Terminal 2017~ (wing)"
"International Terminal 2020-~05 (wing)"
"94128 null (postalcode)"
"International Terminal 2000~ (wing)"
"International Terminal 2006~ (wing)"
"International Terminal 2011~ (wing)"
"International Terminal 2014~ (wing)"
"International Terminal 2017~ (wing)"
"International Terminal 2019-07-23 (wing)"
"International Terminal 2020-~05 (wing)"
"International Terminal 2021-05-25 (wing)"
"International Terminal 2021-11-09 (wing)"
"International Terminal Main Hall 2017~ (concourse)"
"International Terminal Main Hall 2019-07-23 (concourse)"
"International Terminal Main Hall 2020-~05 (concourse)"
"International Terminal Main Hall 2021-05-25 (concourse)"
"International Terminal Main Hall 2021-11-09 (concourse)"
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial
* https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite
* https://github.com/whosonfirst/go-whosonfirst-spatial-grpc
* https://github.com/grpc/grpc-go