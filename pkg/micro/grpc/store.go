package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) update(ins *micro.Micro) error {
	ins.UpdateAt = ftime.Now().Timestamp()
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": ins.Id}, bson.M{"$set": ins})
	if err != nil {
		return exception.NewInternalServerError("update service(%s) error, %s", ins.Name, err)
	}

	return nil
}
