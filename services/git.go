package services

import (
	"fmt"
	"github.com/kiliczsh/gondit/model"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Clone(scan model.Scan)  string {

	dir, err := ioutil.TempDir("", "gondit-clone")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer os.RemoveAll(dir)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: scan.Url + ".git",
	})

	if err != nil {
		log.Fatal(err)
		return ""
	}

	cmd := exec.Command("cp", "-r", dir, "/tmp/src")
	_, e := cmd.Output()
	if e != nil && e.Error() != "exit status 1" {
		fmt.Println(e)
		return ""
	}
	return "/tmp/src"
}
