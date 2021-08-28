// +build linux

package key

import "os/exec"

func genKeys() error {
	scriptPath := projectDir() + "/tools/gen-https-keys.sh"
	if err := exec.Command("chmod", "+x", scriptPath).Run(); err != nil {
		return err
	}
	return exec.Command("sh", scriptPath).Run()
}
