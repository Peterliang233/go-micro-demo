package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/peterliang/demo/user/model"
)



//将对数据库的操作方法封装在一个对象里面
type Repository interface {
	Find(id int32) (*model.User, error)   //通过用户的id查找用户
	Create(*model.User) error  //数据库中创建用户
	Update(*model.User, int64) (*model.User, error)  //更新用户信息
	FindByField(string, string, string)(*model.User, error)  //查找用户
}

type User struct {
	Db *gorm.DB
}



//方法的具体实现
func(repo *User) Find(id int32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	if err := repo.Db.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func(repo *User) Create(user *model.User) error {
	if err := repo.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *User) Update(user *model.User)(*model.User, error) {
	if err := repo.Db.Model(user).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func (repo *User) FindByField(key string, value string, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}

	user := &model.User{}

	if err := repo.Db.Select(fields).Where(key+"= ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}