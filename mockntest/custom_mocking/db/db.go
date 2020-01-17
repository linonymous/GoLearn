package db

type DML interface{
	CreateTable(string) bool
	ListTable() []string
	DeleteTable(string) string
}

type DB struct{
	Tables []string
}

func (d *DB) CreateTable(s string) bool{
	if s == ""{
		return false
	}
	if IsExists(d.Tables, s) != -1 {
		return false
	}
	d.Tables = append(d.Tables, s)
	return true
}

func (d *DB) ListTable() []string  {
	if d.Tables == nil{
		return []string{}
	}
	return d.Tables
}

func IsExists(s []string, p string) int  {
	var index int = -1
	for i, t := range s{
		if p == t{
			index = i;
			break
		}
	}
	return index
}

func (d *DB) DeleteTable(s string) string {
	index := IsExists(d.Tables, s)
	if index == -1{
		return ""
	}
	deletedString := d.Tables[index]
	copy(d.Tables[index:], d.Tables[index+1:])
	d.Tables = d.Tables[:len(d.Tables)-1]
	return deletedString
}