package service

import (
	"time"

	"fmgo.io/microv2/go-micro/v2/config/source"
	proto "fmgo.io/microv2/go-micro/v2/config/source/service/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      []byte(c.Data),
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
