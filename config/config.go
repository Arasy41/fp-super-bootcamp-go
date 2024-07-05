package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	JWTSecret      string
	SupabaseURL    string
	SupabaseKey    string
	SupabaseBucket string
	CloudinaryURL  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBPort:         os.Getenv("DB_PORT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		SupabaseURL:    os.Getenv("SUPABASE_URL"),
		SupabaseKey:    os.Getenv("SUPABASE_KEY"),
		SupabaseBucket: os.Getenv("SUPABASE_BUCKET"),
		CloudinaryURL:  os.Getenv("CLOUDINARY_URL"),
	}
}
