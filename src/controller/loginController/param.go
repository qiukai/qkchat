package loginController

import "encoding/json"

type LoginParam struct {
	Name string `json:"name"`
}

func (self *LoginParam) String() string {
	bytes, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
