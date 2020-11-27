package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	cerror "github.com/KimJeongChul/go-resource-monitor/error"
	"github.com/KimJeongChul/go-resource-monitor/logger"
	"github.com/KimJeongChul/go-resource-monitor/resourceprofiler"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// MonitorConfig
type MonitorConfig struct {
	Port    string `json:"port"`
	Period  int    `json:"period"`
	LogPath string `json:"logPath"`
}

// Read config.json and load MonitorConfig
func LoadConfigJson(fileName *string) (monitorConfig *MonitorConfig, cErr *cerror.CError) {
	monitorConfig = nil
	file, err := os.Open(*fileName)
	if err != nil {
		cErr = cerror.NewCError(cerror.OS_OPEN_ERR, err.Error())
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&monitorConfig)
	if err != nil {
		cErr = cerror.NewCError(cerror.OS_OPEN_ERR, err.Error())
		return
	}
	return
}

func main() {
	var cErr *cerror.CError
	defer func() {
		if cErr != nil {
			logger.LogE("main", "main", cErr.Error())
		}
	}()

	// Load config file
	configFilePath := flag.String("c", "./config.json", "set agent config file")
	config, cErr := LoadConfigJson(configFilePath)
	if cErr != nil {
		logger.LogE("main", "main", "Config File:"+*configFilePath+" load error.")
		os.Exit(-1)
	}

	// Logging file
	logPath := config.LogPath + "/%Y%m%d.log"
	rlLogger, err := rotatelogs.New(logPath)
	if err != nil {
		logger.LogE("main", "UNDEFINED", "Cannot create log file:", logPath, "error:", err)
	}
	log.SetOutput(rlLogger)

	// Create ResourceProfiler
	rp := resourceprofiler.New(config.Period)
	rp.Monitor()
}
