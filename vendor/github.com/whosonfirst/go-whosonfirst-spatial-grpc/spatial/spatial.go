package spatial

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-spatial-pip"
)

func NewPointInPolygonRequestFromFlagSet(fs *flag.FlagSet) (*PointInPolygonRequest, error) {

	pip_req, err := pip.NewPointInPolygonRequestFromFlagSet(fs)

	if err != nil {
		return nil, err
	}

	return NewPointInPolygonRequest(pip_req)
}

func NewPointInPolygonRequest(pip_req *pip.PointInPolygonRequest) (*PointInPolygonRequest, error) {

	lat32 := float32(pip_req.Latitude)
	lon32 := float32(pip_req.Longitude)

	req := &PointInPolygonRequest{
		Latitude:  lat32,
		Longitude: lon32,
	}

	return req, nil
}
