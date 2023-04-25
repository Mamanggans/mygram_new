package database

// uncomment ini kalau butuh buat di lokal
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "jakarta2017"
// 	dbname   = "finalProject"
// 	port     = "5432"
// 	dialect  = "postgres"

// 	// host     = "localhost"
// 	// port     = "5432"
// 	// user     = "postgres"
// 	// password = "tolong isi password db yang dituju"
// 	// dbname   = "challengeTen"
// 	// dialect  = "postgres"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

// func handleDatabaseConnection() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname,
// 	)

// 	db, err = sql.Open(dialect, psqlInfo)

// 	if err != nil {
// 		log.Panic("error occured while trying to validate database arguments:", err)
// 	}

// 	err = db.Ping()

// 	if err != nil {
// 		log.Panic("error occured while trying to connect to database:", err)
// 	}

// }

// func handleCreateRequiredTables() {
// 	userTable := `
// 	CREATE TABLE IF NOT EXISTS "user" (
// 		id VARCHAR(50) NOT NULL,
// 		username VARCHAR(50) NOT NULL,
// 		email VARCHAR(50) NOT NULL,
// 		password VARCHAR(255) NOT NULL,
// 		age INT UNSIGNED NOT NULL,
// 		profileImageUrl VARCHAR(255),
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 		PRIMARY KEY (id),
// 		UNIQUE KEY user_username_uindex (username),
// 		UNIQUE KEY user_email_uindex (email)
// 	  );
// 	`

// 	image := `
// 	CREATE TABLE IF NOT EXISTS "image" (
// 		id VARCHAR(50) NOT NULL,
// 		title VARCHAR(50) NOT NULL,
// 		caption TEXT,
// 		image_url VARCHAR(255) NOT NULL,
// 		user_id VARCHAR(50) NOT NULL,
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 		PRIMARY KEY (id),
// 		CONSTRAINT image_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ON UPDATE CASCADE
// 	  );
// 	  `
// 	comment := `
// 	CREATE TABLE IF NOT EXISTS "comment" (
// 		id VARCHAR(50) NOT NULL,
// 		user_id VARCHAR(50) NOT NULL,
// 		image_id VARCHAR(50) NOT NULL,
// 		message VARCHAR(255) NOT NULL,
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 		PRIMARY KEY (id),
// 		CONSTRAINT comment_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ON UPDATE CASCADE,
// 		CONSTRAINT comment_image_id_fk FOREIGN KEY (image_id) REFERENCES
// 		image (id) ON DELETE CASCADE ON UPDATE CASCADE
// 	  );
// `
// 	SocialMedia := `
// 				CREATE TABLE IF NOT EXISTS social_medias (
// 					id VARCHAR(50) PRIMARY KEY,
// 					name VARCHAR(50) NOT NULL,
// 					social_media_url VARCHAR(255) NOT NULL,
// 					user_id VARCHAR(50) NOT NULL,
// 					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 					CONSTRAINT fk_social_media_user_id
// 					FOREIGN KEY (user_id)
// 					REFERENCES users (id)
// 					ON UPDATE CASCADE
// 					ON DELETE CASCADE
// 				);
// 			`

// 	createTableQueries := fmt.Sprintf("%s %s", userTable, SocialMedia, comment, image)

// 	_, err = db.Exec(createTableQueries)

// 	if err != nil {
// 		log.Panic("error occured while trying to create required tables:", err)
// 	}
// }

// func InitiliazeDatabase() {
// 	handleDatabaseConnection()
// 	handleCreateRequiredTables()
// }

// func GetDatabaseInstance() *sql.DB {
// 	return db
// }

// uncomment ini kalau butuh buat di rail way

import (
	"fmt"
	"log"
	"mygram-byferdiansyah/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	var (
		env      = os.Getenv("ENV")
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDBNAME")
		port     = os.Getenv("PGPORT")
		timeZone = os.Getenv("TIMEZONE")
		dsn      = ""
		db       *gorm.DB
		err      error
	)

	if env == "production" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=%s", host, user, password, dbname, port, timeZone)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, timeZone)
	}

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true}); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = db.AutoMigrate(&domain.User{}, &domain.Image{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	return db
}
