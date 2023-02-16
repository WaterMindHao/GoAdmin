package controllers

import "fmt"

type BaseControllers struct {
}

func (c *BaseControllers) Index() {
	fmt.Println("这是首页")
}
