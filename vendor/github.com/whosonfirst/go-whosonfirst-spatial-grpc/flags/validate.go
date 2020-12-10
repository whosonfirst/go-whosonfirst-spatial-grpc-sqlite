package flags

import (
	"errors"
	"flag"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

func ValidateServerFlags(fs *flag.FlagSet) error {

	addr, err := spatial_flags.StringVar(fs, "server-address")

	if err != nil {
		return err
	}

	if addr == "" {
		return errors.New("Invalid server address")
	}

	return nil
}
