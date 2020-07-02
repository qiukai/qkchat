package loginController

import "encoding/json"

type LoginParam struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func (self *LoginParam) String() string {
	bytes, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
