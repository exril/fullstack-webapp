package controllers

import (
	"fmt"
	"time"
)

func CreateTimer() {
	createTime := time.NewTimer(2*time.Second)
	fmt.Println("Starting Timer")
	<-createTime.C
	fmt.Println("TImer EXPIRED")
}