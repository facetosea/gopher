package _case

import (
	"encoding/json"
	"fmt"
)

func T1(){
	type AAA struct  {
		Aa int `json:"aa"`
	}
	a := AAA{
		Aa: 3,
	}
	bodyBytes, err := json.Marshal(a)
	if err != nil {
		type BB struct  {
			Bb string `json:"bb"`
		}
		b := BB{
			Bb: string(bodyBytes),
		}
		str, _ := json.Marshal(b)
		fmt.Println(str)
	}
}
