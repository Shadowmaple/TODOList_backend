package model

type UserModel struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
	QQ       string `json:"qq" gorm:"column:qq"` // QQ 的用户唯一标识符，使用的是 QQ 的 openID
}

func (*UserModel) TableName() string {
	return "user"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(u).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint32) error {
	user := UserModel{ID: id}
	return DB.Self.Delete(&user).Error
}

// GetUser gets an user by the user identifier.
func GetUserById(id uint32) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

// GetUser gets an user by the user identifier.
func GetUserByQQ(qq string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("qq = ?", qq).First(&u)
	return u, d.Error
}

// // Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
// func (u *UserModel) Compare(pwd string) (err error) {
// 	err = auth.Compare(u.Password, pwd)
// 	return
// }

// // Encrypt the user password.
// func (u *UserModel) Encrypt() (err error) {
// 	u.Password, err = auth.Encrypt(u.Password)
// 	return
// }

// // Validate the fields.
// func (u *UserModel) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(u)
// }
