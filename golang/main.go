package main

import (
	"fmt"

	"github.com/DataDog/go-python3"

	"shorten/routers"
)

func main() {
	r := routers.SetupRouter()
	defer python3.Py_Finalize()

	//defer python3.Py_Finalize()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
