package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/server"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	spatial_app "github.com/whosonfirst/go-whosonfirst-spatial/app"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServerApplication struct {
}

func NewGRPCServerApplication(ctx context.Context) (*GRPCServerApplication, error) {

	server_app := &GRPCServerApplication{}
	return server_app, nil
}

func (server_app *GRPCServerApplication) Run(ctx context.Context) error {

	fs, err := spatial_flags.CommonFlags()

	if err != nil {
		return fmt.Errorf("Failed to instantiate common flags, %v", err)
	}

	err = flags.AppendServerFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to append server flags, %v", err)
	}

	spatial_flags.Parse(fs)

	return server_app.RunWithFlagSet(ctx, fs)
}

func (server_app *GRPCServerApplication) RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	err := spatial_flags.ValidateCommonFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate common flags, %v", err)
	}

	err = flags.ValidateServerFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate server flags, %v", err)
	}

	addr, _ := spatial_flags.StringVar(fs, "server-address")

	sp, err := spatial_app.NewSpatialApplicationWithFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to create new spatial application, %v", err)
	}

	paths := fs.Args()

	err = sp.IndexPaths(ctx, paths...)

	if err != nil {
		return fmt.Errorf("Failed to index paths, %v", err)
	}

	spatial_server, err := server.NewSpatialServer(sp)

	if err != nil {
		return fmt.Errorf("Failed to create spatial server, %v", err)
	}

	grpc_server := grpc.NewServer()

	spatial.RegisterSpatialServer(grpc_server, spatial_server)

	log.Printf("Listening for requests on %s\n", addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return grpc_server.Serve(lis)
}
