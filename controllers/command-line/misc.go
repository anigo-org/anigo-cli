package commandline

import "os/exec"

func Play(url string) error {
	return exec.Command("mpv", url).Run()
}
