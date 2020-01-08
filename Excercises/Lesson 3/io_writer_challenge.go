/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_3

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error)  {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))

	for i, c := range p {
		if c >= 'a' && c<= 'z'{
			c -= diff
		}
		out[i] = c
	}

	return c.wtr.Write(out)
}

func main()  {
	c := &Capper{os.Stdout}
	cmd := exec.Command("docker --version")
	fmt.Println(cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	n, err := fmt.Fprintln(c, "Hello There")
	if err == nil || n != 0 {
		os.Exit(1)
	}
}