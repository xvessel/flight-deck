package component

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Component struct {
	Name string
	Dir  string
}

func (c *Component) Run(cmdstr string, env []string, namespace string, id string) (error, []string) {
	cmd := exec.Command("bash", "-c", cmdstr+" "+namespace+" "+id)
	cmd.Env = env
	cmd.Dir = c.Dir

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err, stdout.String(), stderr.String())
	}
	tmp := strings.Split(stdout.String(), "\n")
	ret := make([]string, 0)
	for _, j := range tmp {
		if strings.TrimSpace(j) != "" {
			ret = append(ret, j)
		}
	}
	return err, ret
}

func (c *Component) Input() (ret map[string][2]string, err error) {
	return extractFile(c.Dir + "/" + c.Name + "/INPUT")
}

func (c *Component) Output() (ret map[string][2]string, err error) {
	return extractFile(c.Dir + "/" + c.Name + "/OUTPUT")
}

func extractFile(fileName string) (ret map[string][2]string, err error) {
	ret = make(map[string][2]string)
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		var key, val, comment string
		keys := strings.Split(line, "=")
		if len(keys) == 2 {
			vals := strings.Split(keys[1], "#")
			key = strings.TrimSpace(keys[0])
			val = strings.TrimSpace(vals[0])
			if len(vals) == 2 {
				comment = strings.TrimSpace(vals[1])
			}
			ret[key] = [2]string{val, comment}
		}

		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return nil, err
		}
	}
	return ret, nil
}
