package models

//结构定义
type User struct {
	Id       uint64 `gorm:"primary_key;commit:'用户ID'"`
	Name     string `gorm:"size:100;unique,not null;commit:'用户名'"`
	Password string `gorm:"size:100;unique,not null;commit:'密码'"`
	Phone    string `gorm:"size:11;commit:'手机号'"`
	Email    string `gorm:"size:100;index;commit:'Email'"`
	State    int    `gorm:"default:1;commit:'状态1正常0关闭'" `
	Message  string `gorm:"commit:'状态1正常0关闭'" `
}

//设置表名
func (u User) TableName() string {
	return "user"
}


