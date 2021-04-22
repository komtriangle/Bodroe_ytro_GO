package repositories


import(
	"github.com/jinzhu/gorm"
)



type UserIRepository interface{
	Insert(user *User) (bool, error)
	GetAll() ([]User, error)	
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{GetDB()}
}

func (u UserRepository) Insert(user *User) (bool, error){
	res, err := user.Validate()
	if(!res){
		return false, err
	}
	err = u.db.Create(&user).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
} 

func(u UserRepository) GetAll() ([]User, error){
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}