package password

import (
	"context"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/domain"
	"github.com/kade-chen/mcenter/apps/token"
	"github.com/kade-chen/mcenter/apps/token/provider"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/rs/zerolog"
)

var (
	AUTH_FALIED = exception.NewUnauthorized("user or password not empty")
)

// var _ provider.Issuer = &password{}

var _ provider.Issuer = (*password)(nil)

func init() {
	provider.Registe(&password{})
}

type password struct {
	user   user.Service
	log    *zerolog.Logger
	domain domain.Service
}

func (p *password) Init() error {
	p.user = ioc.Controller().Get(user.AppName).(user.Service)
	p.domain = ioc.Controller().Get(domain.AppName).(domain.Service)
	p.log = log.Sub("issuer_password_token")
	return nil
}

func (*password) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PASSWORD // 密码令牌
}

func (p *password) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := p.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// fmt.Println("=-----", u)
	p.log.Info().Msg("password token")
	// 3. 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.SharedUser = u.Spec.Shared
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id

	return tk, err
}

func (p *password) validate(ctx context.Context, username, password string) (*user.User, error) {

	if username == "" || password == "" {
		return nil, AUTH_FALIED
	}

	//1.query user whether in the document
	u, err := p.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(username))
	if err != nil {
		return nil, err
	}
	//2.verify password
	if err := u.Password.CheckPassword(password); err != nil {
		return nil, exception.NewBadRequest("user or password not error: %s", err.Error())
	}

	// 2.check whether the password has expried
	var expiredRemain, expiredDays uint
	switch u.Spec.Type {
	case user.TYPE_SUB:
		// 子账号过期策略
		d, err := p.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(u.Spec.Domain))
		if err != nil {
			return nil, err
		}
		ps := d.Spec.PasswordConfig
		expiredRemain, expiredDays = uint(ps.BeforeExpiredRemindDays), uint(ps.PasswordExpiredDays)
	default:
		// 主账号和管理员密码过期策略
		expiredRemain, expiredDays = uint(u.Password.ExpiredRemind), uint(u.Password.ExpiredDays)
	}

	err = u.Password.CheckPasswordExpired(expiredRemain, expiredDays)
	if err != nil {
		return nil, err
	}

	return u, err
}
