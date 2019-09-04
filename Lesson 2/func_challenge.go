/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_2

import (
	"fmt"
	"net/http"
)


func contentType(url string) (string, error)  {
	resp, err:= http.Get(url)
	if err != nil{
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {
		return "", fmt.Errorf("Cant find content type header")
	}
	return ctype, err
}

func main()  {
	ctype, err := contentType("https://golang.org")
	fmt.Println(ctype, " : ", err)
}
