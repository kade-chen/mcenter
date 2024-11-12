package impl

import (
	"context"
	"fmt"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

// query wether  the token  exists for mongdb
func (s *service) get(ctx context.Context, id string) (*token.Token, error) {
	filter := bson.M{"_id": id}

	tk := token.NewToken(token.NewIssueTokenRequest())
	// find
	if err := s.col.FindOne(ctx, filter).Decode(&tk); err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error().Msgf("%s is not found for token error: %s", tk.AccessToken, err.Error())
			return nil, exception.NewNotFound("%s is not found for token error: %s", tk.AccessToken, err.Error())
		}
		return nil, exception.NewInternalServerError("Query token %s error:%s", tk.AccessToken, err.Error())
	}
	return tk, nil
}

func (s *service) save(ctx context.Context, tk *token.Token) error {
	if _, err := s.col.InsertOne(ctx, tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, ins *token.Token) error {
	if ins == nil || ins.AccessToken == "" {
		return fmt.Errorf("access tpken is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.AccessToken})
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", ins.AccessToken, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("book %s not found", ins.AccessToken)
	}

	return nil
}
