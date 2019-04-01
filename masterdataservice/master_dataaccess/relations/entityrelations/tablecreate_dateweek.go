package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateDateWeek(database *gorm.DB) {
	if !database.HasTable(&ent.TableDateWeek{}) {
		database.CreateTable(&ent.TableDateWeek{})
		database.Model(&ent.TableDateWeek{}).AddForeignKey("financialyearid", "table_financialyear(id)", "CASCADE", "RESTRICT")
	}
}
