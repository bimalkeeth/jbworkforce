package entityrelations

import "github.com/jinzhu/gorm"
import ent "jbworkforce/masterdataservice/master_dataaccess/entities"

func CreateAllowance(database *gorm.DB) {
	if !database.HasTable(&ent.TableAllowance{}) {
		database.CreateTable(&ent.TableAllowance{})
		database.Model(&ent.TableAgencyContacts{}).AddForeignKey("allowancegroupid", "table_allowance(id)", "RESTRICT", "RESTRICT")
	}
}
