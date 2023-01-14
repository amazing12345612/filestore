package main

import (
	"filestore/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(":8088")
}
