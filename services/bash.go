package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kiliczsh/gondit/model"
	"os/exec"
)

func Execute(db *gorm.DB, scan model.Scan) string{

	res := Clone(scan)
	if res == "" {
		return fmt.Sprintf("Cloning %s failed!", scan.Url)
	}

	cmd := exec.Command("docker", "run", "--rm", "-v", res + ":/code", "opensorcery/bandit", "-f", "json", "-r", "/code")
	stdout, err := cmd.Output()
	if err != nil && err.Error() != "exit status 1" {
		fmt.Println(err.Error())
		return err.Error()
	}
	scan.Response = string(stdout)
	db.Save(&scan)
	db.Find(&scan, scan.ID)

	return scan.Response
}
