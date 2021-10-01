package services

import (
	"github.com/gofiber/fiber"
	"github.com/kiliczsh/gondit/model"
	"github.com/otiai10/copy"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
)

func Clone(scan model.Scan)  fiber.Handler{

	dir, err := ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: scan.Url + ".git",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll("/tmp/src")

	err = copy.Copy(dir, "/tmp/src")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}