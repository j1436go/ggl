package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const base = "https://www.google.de/search"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing search parameters\n")
		os.Exit(1)
	}
	u, _ := url.Parse(base)
	q := u.Query()
	q.Add("q", strings.Join(os.Args[1:], " "))
	u.RawQuery = q.Encode()
	open(u.String())
}

func open(u string) {
	var cmd string
	switch {
	case runtime.GOOS == "linux" || runtime.GOOS == "openbsd" || runtime.GOOS == "freebsd":
		cmd = "xdg-open"
	case runtime.GOOS == "darwin":
		cmd = "open"
	case runtime.GOOS == "windows":
		cmd = "start"
	default:
		fmt.Fprintln(os.Stderr, "os not supported")
		os.Exit(1)
	}
	_, err := exec.Command(cmd, u).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute google search with command %s: %s\n", cmd, err.Error())
		os.Exit(1)
	}
}
