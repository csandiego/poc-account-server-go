package grpc

import (
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"os"
)

func fromMessage(src proto.Message, dst interface{}) error {
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	defer r.Close()
	defer w.Close()
	m := jsonpb.Marshaler{}
	d := json.NewDecoder(r)
	if err = m.Marshal(w, src); err != nil {
		return err
	}
	return d.Decode(dst)
}
