package inits

import (
	"fmt"
	"os"
	"shorten/base"
	"sync"

	"github.com/DataDog/go-python3"
)

var lock = &sync.Mutex{}

type Single struct {
	Name   string
	File   *python3.PyObject
	Method *python3.PyObject
}

func Init(single Single) Single {

	fmt.Println("######0")

	sysModule := python3.PyImport_ImportModule("sys")

	path := sysModule.GetAttrString("path")

	//fmt.Println("###################1")
	//pathStr, _ := pythonRepr2(path)
	//log.Println("before add path is " + pathStr)
	//python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("//middlewares//"))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(" "))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("/.py"))
	fmt.Println("######5")

	dir, _ := os.Getwd()
	//dir = dir + "\\middlewares\\"
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(base.Python_dir))
	//fmt.Println(dir)
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	//pathStr, _ := PythonRepr(path)
	//fmt.Println(pathStr)
	fmt.Println("######1")

	concurrencyFile := python3.PyImport_ImportModule("cat_dog_classifier")

	if concurrencyFile == nil {
		panic("concurrencyFile == nil")
	}
	fmt.Println("######2")

	oDict := python3.PyModule_GetDict(concurrencyFile)

	genTestdata := python3.PyDict_GetItemString(oDict, "train")

	fmt.Println("######2")

	single.File = oDict

	single.Method = genTestdata
	fmt.Println(single.Method)
	fmt.Println(genTestdata)
	fmt.Println("######3")
	//a := python3.PyFloat_FromDouble(1)

	single.Name = "a"
	return single
}

var once sync.Once
var I int

func GetInstance() Single {
	I++
	//	once.Do(

	fmt.Println("Creating single instance now.")
	var singleInstance Single
	singleInstance = Single{}
	singleInstance = Init(singleInstance)

	//fmt.Println("###################2")

	//fmt.Println("###################3")
	fmt.Println(I)
	fmt.Println("singleInstance")
	fmt.Println(&singleInstance.Method)

	return singleInstance
}

func PythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}

	s := o.Repr()
	if s == nil {
		python3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}

	return python3.PyUnicode_AsUTF8(s), nil
}
