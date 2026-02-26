package util

import "os"

func GetPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic("get pwd error")
	}
	return pwd
}
