package user

import (
	"fmt"
	"time"

	"github.com/kade-chen/library/exception"
	"golang.org/x/crypto/bcrypt"
)

// hash password
func NewHashPassword(password string) (*Password, error) {
	//1.generate hash password
	bycrptpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, exception.NewBadRequest("hash password error: %s", err.Error())
	}
	//2.create password object
	p := &Password{
		Password:      string(bycrptpassword),
		CreateAt:      time.Now().Unix(),
		UpdateAt:      time.Now().Unix(),
		ExpiredDays:   180,
		ExpiredRemind: 30,
	}
	return p, nil
}

// CheckPassword 判断password 是否正确
func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewBadRequest("用户名或者密码不正确")
	}

	return nil
}

// CheckPasswordExpired 检测password是否已经过期
// remindDays 提前多少天提醒用户修改密码
// expiredDays 多少天后密码过期
func (p *Password) CheckPasswordExpired(remindDays, expiredDays uint) error {
	// 永不过期
	if expiredDays == 0 {
		return nil
	}

	now := time.Now()
	expiredAt := time.Unix(p.UpdateAt, 0).Add(time.Duration(expiredDays) * time.Hour * 24)

	ex := now.Sub(expiredAt).Hours() / 24
	if ex > 0 {
		return exception.NewPasswordExired("password expired %f days", ex)
	} else if ex >= -float64(remindDays) {
		p.SetNeedReset("密码%f天后过期, 请重置密码", -ex)
	}

	return nil
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}
