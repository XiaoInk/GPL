/*
	Created by XiaoInk at 2021/04/23 16:07:24
	GitHub: https://github.com/XiaoInk
*/

package main

import (
	"log"

	"github.com/XiaoInk/GPL/router"

	_ "github.com/XiaoInk/GPL/model"
)

func main() {
	err := router.App.Run()
	if err != nil {
		log.Fatalln("Main.App.Run.err", err.Error())
	}
}
