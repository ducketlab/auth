package grpc

import (
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newDescribeRequest(req *user.DescribeAccountRequest) (*describeUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeUserRequest{req}, nil
}

type describeUserRequest struct {
	*user.DescribeAccountRequest
}

func (r *describeUserRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Account != "" {
		filter["_id"] = r.Account
	}

	return filter
}

func newQueryUserRequest(req *user.QueryAccountRequest) (*queryUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &queryUserRequest{
		QueryAccountRequest: req,
	}, nil
}

type queryUserRequest struct {
	*user.QueryAccountRequest
}

func (r *queryUserRequest) FindOptions() *options.FindOptions {

	opt := &options.FindOptions{
		Sort: bson.D{{Key: "create_at", Value: -1}},
	}

	return opt
}

func (r *queryUserRequest) FindFilter() bson.M {
	filter := bson.M{
		"type":   r.UserType,
		"domain": r.Domain,
	}

	if len(r.Accounts) > 0 {
		filter["_id"] = bson.M{"$in": r.Accounts}
	}
	if r.DepartmentId != "" {
		if r.WithAllSub {
			filter["$or"] = bson.A{
				bson.M{"department_id": bson.M{"$regex": r.DepartmentId, "$options": "im"}},
			}
		} else {
			filter["department_id"] = r.DepartmentId
		}
	}

	if r.Keywords != "" {
		filter["$or"] = bson.A{
			bson.M{"_id": bson.M{"$regex": r.Keywords, "$options": "im"}},
			bson.M{"profile.mobile": bson.M{"$regex": r.Keywords, "$options": "im"}},
			bson.M{"profile.email": bson.M{"$regex": r.Keywords, "$options": "im"}},
		}
	}

	return filter
}
