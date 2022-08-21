package jvdata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func raRead(filename string) (ra *JvRaRace, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf(filename, err)
	}
	defer f.Close()

	ra = NewJvRaRace()
	stream := kaitai.NewStream(f)
	err = ra.Read(stream, nil, ra)
	if err != nil {
		records, _ := json.Marshal(ra)
		return nil, fmt.Errorf(filename, string(records), err)
	}

	return ra, err
}
