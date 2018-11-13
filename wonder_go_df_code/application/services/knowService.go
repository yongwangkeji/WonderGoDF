package services

import (
	"fmt"
	"wonder_go_df/application/model"
)

// 服务层测试方法
func HoHello(appName string) {
	fmt.Println("HoHello", appName)
	model.GetHoInfo()
}
