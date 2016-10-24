package main

import "fmt"
import "os"
import "io"

var path = "c:/contoh/teks.txt"

func cekerror(err error){
	if err != nil {
	fmt.Println(err.Error())
	os.Exit(0)
	}
}

func buatfile(){
var _,err = os.Stat(path)

	if os.IsNotExist(err) {
	var file, err = os.Create(path)
	cekerror(err)
	defer file.Close()
	}
}

func editfile(){

	var file ,err = os.OpenFile(path,os.O_RDWR, 0644)
	cekerror(err)
	defer file.Close()

	_,err = file.WriteString("hallo reno\n")
	cekerror(err)
	_,err = file.WriteString("kamu baru saja berhasil")
	cekerror(err)

	err = file.Sync()
	cekerror(err)
}

func bacafile(){
	var file ,err = os.OpenFile(path,os.O_RDWR, 0644)
	cekerror(err)
	defer file.Close()

	var text = make([]byte, 1024)
	for{
	n, err := file.Read(text)
	if err != io.EOF{
	cekerror(err)
	}
		if n == 0{
		break
		}
	}

func hapusfile(){
var err = os.Remove(path)
cekerror(err)
}

fmt.Println(string(text))
cekerror(err)
}

func main(){
	buatfile()
	editfile()
	bacafile()
}