package trackingSchema

import (
	"github.com/wejick/tego/config"
)

var (
	dbCfg *config.DBConfig
)

//Init init trackingSchema
func Init() (err error) {
	//get db config pointer
	dbCfg = config.Get().DB
	err = initDB()
	if err != nil {
		return
	}
	err = prepareStatement()
	if err != nil {
		return
	}

	return
}
