package roles

import "github.com/go-playground/validator"

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}
