// +build linux dragonfly freebsd netbsd openbsd
package browser

import "os/exec"

func openURL(u string) error {
	return exec.Command("xdg-open", u).Start()
}
