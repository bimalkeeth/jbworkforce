package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Campus
//----------------------------------------------------
func CreateCampus(database *gorm.DB) {
	if !database.HasTable(&ent.TableCampus{}) {

		database.CreateTable(&ent.TableCampus{})
		database.Model(&ent.TableCampus{}).Association("Divisions")
	}
}
