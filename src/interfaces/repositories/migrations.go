package repositories

// Migrate runs database migrations
func Migrate(db DbHandler) error {
	q := `CREATE TABLE dbname.users(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    name VARCHAR(250) NULL DEFAULT NULL,
	    email VARCHAR(100) NOT NULL,
	    password VARCHAR(250) NOT NULL,
	    role VARCHAR(20) NOT NULL DEFAULT 'user',
	    PRIMARY KEY(id),
	    INDEX(role),
	    UNIQUE(email)
	) ENGINE = InnoDB CHARSET = utf8 COLLATE utf8_general_ci;`

	_, err := db.Execute(q)
	return err
}
