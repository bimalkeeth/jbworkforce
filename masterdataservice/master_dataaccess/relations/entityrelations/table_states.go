package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// State
//----------------------------------------------------
func CreateState(database *gorm.DB) {

	if !database.HasTable(&ent.TableState{}) {

		database.CreateTable(&ent.TableState{})
		database.Model(&ent.TableState{}).AddUniqueIndex("idx_table_state_name", "statename")
		database.Model(&ent.TableState{}).AddUniqueIndex("idx_table_state_abbr", "stateabbr")
	}
}
