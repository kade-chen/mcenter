package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/namespace"
	"github.com/kade-chen/mcenter/apps/token"
	"github.com/kade-chen/mcenter/apps/token/provider"
	"go.mongodb.org/mongo-driver/bson"
)

// 1.isseuer token
func (s *service) issuer_token(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 0.get issuer
	issuer := provider.GetTokenIssuer(req.GrantType)

	// 确保有provider
	if issuer == nil {
		return nil, exception.NewBadRequest("grant type %s not support", req.GrantType)
	}

	// 1.issue token
	tk, err := issuer.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}
	// 2.默认放到domain default
	tk.Namespace = namespace.DEFAULT_NAMESPACE

	if !req.DryRun {
		// 入库保存
		if err := s.save(ctx, tk); err != nil {
			return nil, err
		}
		// 关闭之前的web登录
		if err := s.blockOtherWebToken(ctx, tk); err != nil {
			return nil, err
		}
	}

	// 这里应该是返回生成的token和用户信息
	return tk, nil
}

// 1.1 close web logins
func (s *service) blockOtherWebToken(ctx context.Context, tk *token.Token) error {
	// 共享账号不关闭之前的Token
	if tk.SharedUser {
		return nil
	}

	// 如果不是web登陆, 不需要关闭之前的登录令牌
	if !tk.Platform.Equal(token.PLATFORM_WEB) {
		return nil
	}

	now := time.Now()
	status := token.NewStatus()
	status.IsBlock = true
	status.BlockAt = now.UnixMilli()
	status.BlockReason = fmt.Sprintf("你于 %s 从其他地方通过 %s 登录", now.Format(time.RFC3339), tk.GrantType)
	status.BlockType = token.BLOCK_TYPE_OTHER_PLACE_LOGGED_IN

	rs, err := s.col.UpdateMany(
		ctx,
		bson.M{
			"platform":        token.PLATFORM_WEB,
			"user_id":         tk.UserId,
			"issue_at":        bson.M{"$lt": tk.IssueAt},
			"status.is_block": false,
		},
		bson.M{"$set": bson.M{"status": status}},
	)
	if err != nil {
		return err
	}
	s.log.Debug().Msgf("block %d tokens", rs.ModifiedCount)
	return nil
}

// check security after loading
func (s *service) checkSecurity(ctx context.Context, tk *token.Token) error {
	return nil
}

// 检查是否在黑名单
