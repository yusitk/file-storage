package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func loadConfig() (*DBConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/") //доработка

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config DBConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &config, nil
}

func promptForConfig() (*DBConfig, error) {
	var config DBConfig

	fmt.Print("Enter database host: ")
	fmt.Scanln(&config.Host)

	fmt.Print("Enter database port: ")
	fmt.Scanln(&config.Port)

	fmt.Print("Enter database user: ")
	fmt.Scanln(&config.User)

	fmt.Print("Enter database password: ")
	fmt.Scanln(&config.Password)

	fmt.Print("Enter database name: ")
	fmt.Scanln(&config.DBName)

	return &config, nil
}

func InitDB() {
	config, err := loadConfig()
	if err != nil {
		log.Printf("Error loading config from file: %v", err)
		config, err = promptForConfig()
		if err != nil {
			log.Fatal("Error getting database config: ", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)
	log.Printf("Connecting to database with host=%s, port=%d, user=%s, dbname=%s",
		config.Host, config.Port, config.User, config.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	log.Println("Connected to database")
	DB = db
}
