package services

import (
	"app/container"
	"app/core/logger"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//GetMySqlConnection connects to mysql database
func GetMySqlConnection() {
	mysqlSlaveConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	fmt.Println(os.Getenv("DB_USER"))
	dbr, err := sqlx.Connect("mysql", mysqlSlaveConnectionString)
	if err != nil {
		logger.Error.Fatalln("Mysql Slave DB connection failed : ", err)
	}

	maxIdleConnections, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		logger.Error.Fatalln("ENV read failed - DB_MAX_IDLE_CONNECTIONS : ", err)
	}

	maxOpenConnections, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		logger.Error.Fatalln("ENV read failed - DB_MAX_OPEN_CONNECTIONS : ", err)
	}
	dbr.SetMaxIdleConns(maxIdleConnections)
	dbr.SetMaxOpenConns(maxOpenConnections)

	err = dbr.Ping()
	if err != nil {
		logger.Error.Fatalln("Mysql Slave Ping failed : ", err)
	}

	container.SetReader(dbr)

	mysqlMasterConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	dbw, err := sqlx.Connect("mysql", mysqlMasterConnectionString)
	if err != nil {
		logger.Error.Fatalln("Mysql Master DB connection failed : ", err)
	}
	dbw.SetMaxIdleConns(maxIdleConnections)
	dbw.SetMaxOpenConns(maxOpenConnections)

	err = dbw.Ping()
	if err != nil {
		logger.Error.Fatalln("Mysql Master Ping failed : ", err)
	}

	container.SetWriter(dbw)

	logger.Info.Println("Mysql Connection Success")

	_, err = dbw.Exec("CREATE TABLE IF NOT EXISTS `inventory_management`.`items` ( `id` INT NOT NULL AUTO_INCREMENT , `name` VARCHAR(30) NOT NULL , `warehouse_id` INT NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	if err != nil {
		logger.Error.Fatalln("Error while creating items table : ", err)
	}

	_, err = dbw.Exec("CREATE TABLE IF NOT EXISTS `inventory_management`.`warehouses` ( `id` INT NOT NULL AUTO_INCREMENT , `name` VARCHAR(30) NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	if err != nil {
		logger.Error.Fatalln("Error while creating warehouses table : ", err)
	}
}
