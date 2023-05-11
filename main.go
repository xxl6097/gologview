package main

import (
	"github.com/xxl6097/gologview/go/logview"
)

func main() {

	//aa := http.Dir("../")
	//fmt.Println(aa)
	//
	//dirs, err := os.ReadDir("./assets")
	//if err != nil {
	//	return
	//}
	//for _, dir := range dirs {
	//	if !dir.IsDir() {
	//		arr, _ := os.ReadFile(filepath.Join("./assets", dir.Name()))
	//		fmt.Println(string(arr))
	//	}
	//}

	//this.subRouter.Handle("/file", http.FileServer()).Methods("GET")
	logview.New().Start(8886)
}
