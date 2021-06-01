package config

import "fmt"

const (
	DBUser     = "root"
	DBPassword = "root"
	DBName     = "root"
	DBHost     = "0.0.0.0"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConn() string {
	dbBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
	return dbBase
}
