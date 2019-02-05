package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	con "jbworkforce/masterdataservice/master_dataaccess/contracts"
	"strings"
)

type Db struct{}

func (db *Db) OpenDatabase(clientInfo *con.ClientInfo) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		clientInfo.HostName,
		clientInfo.Port,
		strings.ToLower(clientInfo.UserName),
		fmt.Sprintf("%s_%s", strings.ToLower(clientInfo.Client), "master"),
		clientInfo.Password)

	gormDb, err := gorm.Open("postgres", connectionString)
	return gormDb, err
}
