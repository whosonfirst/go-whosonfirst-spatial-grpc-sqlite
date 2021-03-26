package flags

import (
	"flag"
	"github.com/sfomuseum/go-flags/lookup"
)

func ValidateGRPCServerFlags(fs *flag.FlagSet) error {

	var err error

	_, err = lookup.StringVar(fs, HOST)

	if err != nil {
		return err
	}

	_, err = lookup.IntVar(fs, PORT)

	if err != nil {
		return err
	}

	return nil
}

func ValidateGRPCClientFlags(fs *flag.FlagSet) error {

	var err error

	_, err = lookup.StringVar(fs, HOST)

	if err != nil {
		return err
	}

	_, err = lookup.IntVar(fs, PORT)

	if err != nil {
		return err
	}

	_, err = lookup.BoolVar(fs, TO_STDOUT)

	if err != nil {
		return err
	}

	_, err = lookup.BoolVar(fs, TO_NULL)

	if err != nil {
		return err
	}

	return nil
}
