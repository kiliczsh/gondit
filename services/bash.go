package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kiliczsh/gondit/model"
	"os/exec"
)

func Execute(db *gorm.DB, scan model.Scan) string {

	cmd := exec.Command("docker", "run", "--rm", "-v", "/tmp/src:/code",
		"opensorcery/bandit", "-f", "json", "-r", "/code")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	scan.Response = string(stdout)
	db.Save(&scan)
	db.Find(&scan, scan.ID)
	return scan.Response
}
