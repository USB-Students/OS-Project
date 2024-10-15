package config

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ResultDirectory string
	TcpHost         string
	TcpPort         int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file , Error: %v", err)
	}

	ResultDirectory = GetEnv("RESULTS_DIRECTORY", "")
	if strings.HasSuffix(ResultDirectory, "/") {
		ResultDirectory = ResultDirectory[:len(ResultDirectory)-1]
	}
	if strings.HasSuffix(ResultDirectory, "\\") {
		ResultDirectory = ResultDirectory[:len(ResultDirectory)-1]
	}

	TcpPort = GetEnvAsInt("TCP_PORT", 2000)
	tempHost := GetEnv("TCP_HOST", "127.0.0.1")

	ipPattern := `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipPattern)

	if re.MatchString(tempHost) {
		TcpHost = tempHost
	} else {
		log.Println(tempHost, " is not a valid IP address.\n TCP_HOST will be set to 127.0.0.1")
		TcpHost = "127.0.0.1"
	}
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func GetEnvAsInt(name string, defaultVal int) int {
	valStr := GetEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
