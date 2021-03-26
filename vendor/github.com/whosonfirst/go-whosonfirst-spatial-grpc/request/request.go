package request

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial-pip"
)

func NewPointInPolygonRequestFromFlagSet(fs *flag.FlagSet) (*spatial.PointInPolygonRequest, error) {

	pip_req, err := pip.NewPointInPolygonRequestFromFlagSet(fs)

	if err != nil {
		return nil, err
	}

	return NewPointInPolygonRequest(pip_req)
}

// https://github.com/whosonfirst/go-whosonfirst-spatial-pip/blob/main/pip.go

func NewPointInPolygonRequest(pip_req *pip.PointInPolygonRequest) (*spatial.PointInPolygonRequest, error) {

	lat32 := float32(pip_req.Latitude)
	lon32 := float32(pip_req.Longitude)

	is_current := existentialIntFlagsToProtobufExistentialFlags(pip_req.IsCurrent)
	is_ceased := existentialIntFlagsToProtobufExistentialFlags(pip_req.IsCeased)
	is_deprecated := existentialIntFlagsToProtobufExistentialFlags(pip_req.IsDeprecated)
	is_superseded := existentialIntFlagsToProtobufExistentialFlags(pip_req.IsSuperseded)
	is_superseding := existentialIntFlagsToProtobufExistentialFlags(pip_req.IsSuperseding)

	req := &spatial.PointInPolygonRequest{
		Latitude:            lat32,
		Longitude:           lon32,
		Placetypes:          pip_req.Placetypes,
		Geometries:          pip_req.Geometries,
		AlternateGeometries: pip_req.AlternateGeometries,
		InceptionDate:       pip_req.InceptionDate,
		CessationDate:       pip_req.CessationDate,
		IsCurrent:           is_current,
		IsCeased:            is_ceased,
		IsDeprecated:        is_deprecated,
		IsSuperseded:        is_superseded,
		IsSuperseding:       is_superseding,
	}

	return req, nil
}
