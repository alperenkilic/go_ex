package main

import (
	"os"
)

func main(){

/*	oldPath:="testttt.txt"
	newPath:="./yeniYol/alperen.txt"

	err:=os.Rename(oldPath,newPath)
	if err != nil{
	log.Fatal(err)
	} */
// işin özü os.Rename... err ile dönen eroru yakalıyoruz.
	alp:="a.txt"
	blp:="aaa.txt"

	os.Rename(alp,blp)
}