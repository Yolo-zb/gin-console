package model

const (
	// Dictionary
)

type User struct {
 
	Id int  
	Phone string  
	Username string  
	Avatar string  
	AccessToken string  
	AuthKey string  
	PasswordHash string  
	Status int8  
	Role string  
	ServerId int  
	NotifyNum int32  
	MsgNum int  
	BurseNum float64  
	PointNum int  
	CreatedAt int  
	UpdatedAt int 
}

func (*User) TableName() string {
	return "user"
}
