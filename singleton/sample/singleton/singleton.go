package singleton

import "fmt"

type singleton struct{}

// may replace bool
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		fmt.Println("Create Instance")
		instance = &singleton{}
	}
	return instance
}
