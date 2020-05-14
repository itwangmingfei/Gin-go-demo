package forms

import "github.com/gookit/validate"

type User struct {
	Name     string `json:"name"  validate:"required|minLen:4" `
	Password string `json:"password" validate:"required|minLen:6"`
	Phone    string `json:"phone" validate:"required|cnMobile"`
	Email    string `json:"email" validate:"required|email"`
}
type UpUser struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name"  validate:"required|minLen:4" `
	Password string `json:"password" validate:"required|minLen:6"`
	Phone    string `json:"phone" validate:"cnMobile"`
	Email    string `json:"email" validate:"email"`
	Message  string `json:"message"`
}

// Messages 您可以自定义验证器错误消息
func (f User) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (p User) Translates() map[string]string {
	return validate.MS{
		"Name":     "姓名",
		"Password": "密码",
		"Phone":    "手机号",
		"Email":    "邮箱",
	}
}
