package flags

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/lookup"
	"github.com/whosonfirst/go-whosonfirst-iterate/emitter"
	"sort"
	"strings"
)

func AppendIndexingFlags(fs *flag.FlagSet) error {

	modes := emitter.Schemes()
	sort.Strings(modes)

	valid_modes := strings.Join(modes, ", ")
	desc_modes := fmt.Sprintf("A valid whosonfirst/go-whosonfirst-iterate/emitter URI. Supported schemes are: %s.", valid_modes)

	fs.String(ITERATOR_URI, "repo://", desc_modes)

	return nil
}

func ValidateIndexingFlags(fs *flag.FlagSet) error {

	_, err := lookup.StringVar(fs, ITERATOR_URI)

	if err != nil {
		return err
	}

	return nil
}
