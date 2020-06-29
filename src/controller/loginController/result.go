package loginController

import "encoding/json"

type LoginResult struct {
	Name string `json:"name"`
}

func (self *LoginResult) String() string {
	bytes, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
