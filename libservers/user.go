package libservers

import (
	"fmt"
	"gin/forms"
	"gin/models"
	"log"
)

type User struct {
	Servers
}
//list
func(u User) GetList() []models.User {
	DB := models.GetDb()
	var users []models.User
	u.Name = fmt.Sprintf("%%%s%%",u.Name)
	DB.Where("name like ?",u.Name).Limit(u.Pagesize).Offset(u.Offset).Order("id desc").Find(&users)
	log.Println(users)
	return users
}

//添加
func (u User) Add(form forms.User) (uint64, bool) {
	DB := models.GetDb()
	var user models.User
	/**/
	user.Name = form.Name
	user.Password = form.Password
	user.Email = form.Email
	user.Phone = form.Phone
	DB.Create(&user)
	if user.Id == 0 {
		return 0, false
	}
	return user.Id, true
}

/*
@查询一条数据
*/
func (u User) Show() (models.User, error) {
	DB := models.GetDb()
	var user models.User
	//查主键
	err := DB.First(&user, u.Id).Error
	//其他条件
	//DB.Where("name = ?",u.Name).First(&user)
	if user.Id == 0 {
		return user, err
	}
	return user, nil
}

/*
@删除一条数据
*/
func (u User) Del() bool {
	DB := models.GetDb()
	var user models.User
	//判断是否存在
	user, err := u.Show()
	if err != nil {
		return false
	}
	err = DB.Delete(&user).Error
	if err != nil {
		return false
	}
	return true
}

/*
$修改数据
*/
func (u User) Update(form forms.UpUser) error {
	DB := models.GetDb()
	var user models.User
	user, err := u.Show()
	if err != nil {
		return err
	}
	user.Name = form.Name
	user.Phone = form.Phone
	user.Email = form.Email
	user.Message = form.Message
	err = DB.Model(&user).Update(&user).Error
	if err != nil {
		return err
	}
	return nil
}
