package main

import "Demo/app/configs"

func main() {
	r := configs.SetupRoute()
	r.Run(":3030")
}
