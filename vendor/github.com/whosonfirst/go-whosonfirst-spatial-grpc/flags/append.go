package flags

import (
	"flag"
)

func AppendServerFlags(fs *flag.FlagSet) error {

	fs.String("server-address", "localhost:8282", "A valid gRPC server address")

	return nil
}
