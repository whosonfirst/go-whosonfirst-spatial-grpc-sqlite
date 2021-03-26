package request

import (
	"github.com/skelterjohn/geom"
	geojson_utils "github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
)

func CoordsFromPointInPolygonRequest(req *spatial.PointInPolygonRequest) (geom.Coord, error) {

	lat := float64(req.Latitude)
	lon := float64(req.Longitude)

	return geojson_utils.NewCoordinateFromLatLons(lat, lon)
}
