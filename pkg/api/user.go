package api

type (
	Gender string
	GetUserResponse struct {
		Code int `json:"code" example:"200"`
		Msg string `json:"msg" example:"ok"`
		User User `json:"user"`
	}
	User struct {
		//this is the id of user.
		ID int `json:"id" example:"123"`
		Name string `json:"name"`
		Age int `json:"age" minimum:"10" maximum:"99" default:"20"`
		Gender Gender `json:"gender" enums:"Male, Female"`
	}
)

const (
	Male Gender = "Male"
	Female Gender = "Female"
)
