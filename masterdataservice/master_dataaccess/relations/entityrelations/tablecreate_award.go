package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// TableAward
//----------------------------------------------------
func CreateTableAward(database *gorm.DB) {
	if !database.HasTable(&ent.TableAward{}) {

		database.CreateTable(&ent.TableAward{})
	}
}
