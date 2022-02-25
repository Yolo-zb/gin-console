package model

const (
	// Dictionary
)

type UserAuth struct {
 
	Uid int  
	AuthRule string  
	UpdatedAt int 
}

func (*UserAuth) TableName() string {
	return "user_auth"
}
