package dbmaxmind

import(
	"fmt"
	"github.com/oschwald/geoip2-golang"
)

func GetDB(path string) *geoip2.Reader {
	//Get the DB
	db, err := geoip2.Open(path)
	if err != nil {
		fmt.Println("Error: Database Not Found")
	}
	return db
}