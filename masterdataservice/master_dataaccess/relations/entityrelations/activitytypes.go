package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

func CreateActivityTable(database *gorm.DB) {

	if !database.HasTable(&ent.TableActivityTypes{}) {

		database.CreateTable(&ent.TableActivityTypes{})
		database.Model(&ent.TableActivityTypes{}).AddUniqueIndex("idx_table_activitytype_name", "activitytypename")
		database.Model(&ent.TableActivityTypes{}).AddUniqueIndex("idx_table_activitytype_abbr", "activityabbr")
		database.Model(&ent.TableAddressType{}).Association("clinicalstreams")
	}
}
