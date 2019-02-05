package relations

import (
	"github.com/jinzhu/gorm"
	db "jbworkforce/masterdataservice/master_dataaccess/common"
	con "jbworkforce/masterdataservice/master_dataaccess/contracts"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
	"log"
)

type RelationBuilder struct{}

func (*RelationBuilder) BuildDatabase(connInfo *con.ClientInfo) (err error) {
	var ss db.Idb = &db.Db{}
	database, err := ss.OpenDatabase(connInfo)
	if err != nil {
		log.Fatal("database connection was not successful.")
		return err
	}
	database.SingularTable(true)

	createActivityTable(database)
	return nil
}

func createActivityTable(database *gorm.DB) {
	if !database.HasTable(&ent.TableActivityTypes{}) {
		database.CreateTable(&ent.TableActivityTypes{})
		database.Model(&ent.TableActivityTypes{}).AddUniqueIndex("idx_table_activitytype_name", "activitytypename")
		database.Model(&ent.TableActivityTypes{}).AddUniqueIndex("idx_table_activitytype_abbr", "activityabbr")
	}
}
