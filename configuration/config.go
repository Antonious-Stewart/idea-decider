package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port          string
	NumberOfScans int
}

func GetValues() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port env is not set")
	}

	numberOfScans, err := strconv.Atoi(os.Getenv("NUMBER_OF_SCANS"))

	if err != nil {
		log.Fatal(err)
	}

	if numberOfScans == 0 {
		log.Fatal("Number of scans is not set")
	}

	return &Config{
		Port:          port,
		NumberOfScans: numberOfScans,
	}
}
