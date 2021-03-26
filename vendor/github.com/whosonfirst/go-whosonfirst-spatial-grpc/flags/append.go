package flags

import (
	"flag"
)

func AppendGRPCServerFlags(fs *flag.FlagSet) error {

	fs.String(HOST, "localhost", "The host to listen for requests on")
	fs.Int(PORT, 8082, "The port to listen for requests on")

	return nil
}

func AppendGRPCClientFlags(fs *flag.FlagSet) error {

	fs.String(HOST, "localhost", "The host to listen for requests on")
	fs.Int(PORT, 8082, "The port to listen for requests on")

	fs.Bool(TO_STDOUT, true, "Emit results to STDOUT")
	fs.Bool(TO_NULL, false, "Emit results to /dev/null")

	return nil
}
