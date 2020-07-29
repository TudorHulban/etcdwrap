package etcdwrap

import (
	"log"
	"os/exec"
	"strings"
)

// does not work w etcdl
func shellCmd(pathExec string, params ...string) ([]byte, error) {
	cmd := exec.Command(pathExec, strings.Join(params, " "))
	log.Println("command:", cmd.String())

	result, err := cmd.CombinedOutput()
	log.Println("etcdl result:", string(result))
	return result, err
}
