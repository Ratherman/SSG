package middlewares

import (
	"fmt"
	"os/exec"
	"shorten/inits"
	"strings"

	"github.com/DataDog/go-python3"
)

var I int

func Calc(c string) string {
	abspath := "..\\Golang_AI\\cat_dog_classifier.py"
	//abspath2 := "..\\Golang_AI\\model\\CNN_model_weights.pth"
	fmt.Println(abspath)

	cmd := exec.Command("python", ".\\cat_dog_classifier.py", ".\\model\\CNN_model_weights.pth", c)
	cmd.Dir = "..\\Golang_AI\\"
	fmt.Println("python " + "cat_dog_classifier.py" + " " + c + " " + ".\\model\\CNN_model_weights.pth")
	//	cmd.Dir = "\\Users\\USER\\Desktop\\go\\SSG\\golang"
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println("Python Execution Error :", e)
	}
	fmt.Println("output is :" + string(output))
	result := string(output)
	strs := strings.Split(result, "\n")

	println(len(strs))
	//	sI, e = strconv.Atoi(strs[0])
	//	dI, e = strconv.Atoi(strs[1])
	//	fmt.Println(sI)
	//	fmt.Println(dI)
	return result
}

func Dump_to_python(s string) string {
	python3.Py_Initialize()
	defer python3.Py_Finalize()

	/*defer func() {
		fmt.Println("python3.Py_Finalize()")
		python3.Py_Finalize()

	}()
	python3.Py_Initialize()

	sysModule := python3.PyImport_ImportModule("sys")
	defer sysModule.DecRef()
	path := sysModule.GetAttrString("path")
	defer path.DecRef()
	//pathStr, _ := pythonRepr2(path)
	//log.Println("before add path is " + pathStr)
	//python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("//middlewares//"))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(" "))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("/.py"))
	dir, _ := os.Getwd()
	//dir = dir + "\\middlewares\\"
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(base.Python_dir))
	//fmt.Println(dir)
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	//pathStr, _ := PythonRepr(path)
	//fmt.Println(pathStr)
	python3.PyRun_SimpleString("import sys")

	python3.PyRun_SimpleString("import torch")
	python3.PyRun_SimpleString("import cv2 as cv")
	fmt.Println("module in python is")
	concurrencyFile := python3.PyImport_ImportModule("cat_dog_classifier")

	if concurrencyFile == nil {
		panic("concurrencyFile == nil")
	}

	defer concurrencyFile.DecRef()
	oDict := python3.PyModule_GetDict(concurrencyFile)
	defer oDict.DecRef()
	genTestdata := python3.PyDict_GetItemString(oDict, "Train")
	defer genTestdata.DecRef()

	//a := python3.PyFloat_FromDouble(1)
	*/
	O := inits.GetInstance()

	fmt.Println(O.Method)
	fmt.Println("123")
	fmt.Println(O.Name)
	if O.Method == nil {
		panic("wwww")
	}

	a := python3.PyUnicode_FromString(s)

	//fmt.Println(a)
	//genTestdata := O.Method
	fmt.Println("456")

	//python3.PyEval_InitThreads()

	//State := python3.PyGILState_Ensure()

	testdataPy := O.Method.CallFunctionObjArgs(a) //retval: New reference
	if testdataPy == nil {
		panic("123123")
	}

	//l := genTestdata.CallFunctionObjArgs(a)
	//...

	str, _ := PythonRepr(testdataPy)

	fmt.Println("780")
	//err, _ := PythonRepr(l)
	fmt.Printf("the result :" + str)

	//fmt.Printf("the result :" + err)

	//python3.PyGILState_Release(State)

	return str
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
