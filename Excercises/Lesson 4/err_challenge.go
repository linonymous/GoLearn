/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_4

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil{
		return errors.Wrap(err, "Cant open pid file!")
	}

	if err := os.Remove(pidFile); err != nil{
		log.Printf("Warning: can't remove pid file - %s", err)
	}

	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid)
	if err != nil {
		return errors.Wrap(err, "Bad process ID")
	}

	fmt.Printf("Killing server with pid=%d\n", pid)
	return nil
}

func main()  {
	if err := killServer("server.pid"); err != nil{
		fmt.Println("Deleted the pid successfully!")
	}
}