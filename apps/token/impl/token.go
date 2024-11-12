package impl

import (
	"context"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/token"
)

// This function is used to issue a token based on the given request and return the token.
func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 1.issuer token
	tk, err := s.issuer_token(ctx, req)
	if err != nil {
		return nil, err
	}
	// 2.login security after loading

	return tk, nil
}

func (s *service) ValicateToken(ctx context.Context, req *token.ValicateTokenRequest) (*token.Token, error) {
	// 1.query wether token exist

	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2.verify wether the token is expired
	time, err := tk.CheckIssue_Token_expried(tk.IssueAt, tk.AccessExpiredAt, tk.AccessToken)
	if err != nil {
		return nil, exception.NewAccessTokenExpired("Token expried,Please login again")
	}
	// 3.print info log
	s.log.Info().Msg(time)
	return tk, err
}

// remove Token
func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	//1.query wether the token exist
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	//2.judge wether the refresh token is consistent
	if req.ACCESS_TOKEN_NAME != "ACCESS_TOKEN_COOKIE_KEY" {
		return nil, exception.NewBadRequest("ACCESS_TOKEN_COOKIE_KEY token name not find")
	}

	//3.delete token
	if err := s.delete(ctx, tk); err != nil {
		return nil, err
	}
	return tk, nil
}
