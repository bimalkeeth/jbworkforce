package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateTableAllowanceLevel(database *gorm.DB) {
	if !database.HasTable(&ent.TableAllowanceLevel{}) {
		database.CreateTable(&ent.TableAllowanceLevel{})
	}
}
