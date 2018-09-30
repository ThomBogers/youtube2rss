package youtube2rss

import (
	"log"
	"github.com/pkg/errors"
)

func Check(e error, info string) {
	if e != nil {
		log.Fatal(errors.Wrap(e, info))
	}
}
