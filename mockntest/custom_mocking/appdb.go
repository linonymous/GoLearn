package main

import (
	"fmt"
	"github.com/linonymous/GoLearn/mockntest/custom_mocking/db"
)

func PerformAddition(d db.DML, name string) bool{
	a := d.CreateTable(name)
	return a
}

func GetTables(d db.DML) []string {
	return d.ListTable()
}
func main()  {
	d := db.DB{Tables:make([]string, 5)}
	fmt.Println(PerformAddition(&d, "abc"))
	fmt.Println(PerformAddition(&d, "dba"))
	a := GetTables(&d)
	fmt.Println(a)
}
