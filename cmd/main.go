package main

import (
	"github.com/l1rn/order-handler/internal"

)

func main() {
	r := router.InitializeRouter()
	r.Run(":8081")
}
