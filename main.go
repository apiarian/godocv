package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := make([]string, len(os.Args)-1)
	copy(args, os.Args[1:])

	isVendor, err := maybeFixVendorPackage(args)
	if err != nil {
		fmt.Printf("failed to fix vendor package: %v", err)
		os.Exit(1)
	}

	cmd := exec.Command("godoc", args...)

	out, err := cmd.CombinedOutput()

	if isVendor {
		out = bytes.Replace(out, []byte(`import "./vendor/`), []byte(""), -1)
	}

	if err != nil {
		os.Stdout.Write(out)
	}

	pager := exec.Command("/usr/bin/less")
	pager.Stdin = bytes.NewBuffer(out)
	pager.Stdout = os.Stdout
	if err := pager.Run(); err != nil {
		fmt.Printf("failed to run pager: %v", err)
	}
}

func maybeFixVendorPackage(args []string) (bool, error) {
	if len(args) == 0 {
		return false, nil
	}

	for i := len(args) - 1; i >= 0; i-- {
		if strings.HasPrefix(args[i], "-") {
			break
		}

		maybeVendorPackage := fmt.Sprintf("./vendor/%s", args[i])
		s, err := os.Stat(maybeVendorPackage)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return false, err
		}
		if !s.IsDir() {
			continue
		}

		args[i] = maybeVendorPackage
		return true, nil
	}

	return false, nil
}
