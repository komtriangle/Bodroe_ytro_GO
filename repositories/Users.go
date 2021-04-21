package repositories

type UserIRepository interface{
	Insert(user *User) (bool, error)
	GetAll() ([]User, error)	
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{GetDb()}
}

func (u UserRepository) Insert(user *User) (bool, error){
	err := u.db.Create(&user).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
} 

func(u UserRepository) GetAll() ([]User, error){
	var users []User
	err := h.db.Find(&users).Error
	return training, err
}


