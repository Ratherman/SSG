package main

import (
	"fmt"

	"shorten/routers"

	"github.com/DataDog/go-python3"
)

func main() {
	r := routers.SetupRouter()
	defer python3.Py_Finalize()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
