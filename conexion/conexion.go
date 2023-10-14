package conexion

import( "github.com/jinzhu/gorm"
		"log"
		_ "github.com/jinzhu/gorm/dialects/mysql"
		"gofiber/persona"
)

func Conexion() *gorm.DB{
	db, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/test")

	if err != nil{
		log.Fatal(err)
	}
	return db
}

func Migrate(){
db := Conexion()
defer db.Close()
log.Println("Migrando....")
db.AutoMigrate(persona.Persona{})
}