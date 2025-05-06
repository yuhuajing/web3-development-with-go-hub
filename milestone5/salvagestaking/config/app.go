package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	LoadAllConfigs(".env")
}

type App struct {
	Chain     string
	Server    string
	Register  string
	Funder    string
	Receiver  string
	UnlockGas int
	TransGas  int
	Ratio     int
	Dev       bool
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadApp() {
	chain := os.Getenv("CHAIN")
	server := os.Getenv("ETH_SERVER")
	register := os.Getenv("ETH_REGISTER")

	if strings.Compare(chain, "eth") == 0 {

	} else if strings.Compare(chain, "imx") == 0 {
		server = os.Getenv("IMX_SERVER")
		register = os.Getenv("IMX_REGISTER")
	} else {
		log.Fatalf("不支持的链")
	}

	unlock := os.Getenv("UNLOCK_LIMIT")
	unlockgas, _ := strconv.Atoi(unlock)
	trans := os.Getenv("NFT_TRANS_LIMIT")
	transgas, _ := strconv.Atoi(trans)

	ratio := os.Getenv("RATIO")
	ratios, _ := strconv.Atoi(ratio)

	dev := os.Getenv("DEV")
	devres := strings.Compare(dev, "true")

	funder := os.Getenv("FUNDER")
	receiver := os.Getenv("RECEIVER")
	app.Chain = chain
	app.Server = server
	app.Register = register
	app.UnlockGas = unlockgas
	app.TransGas = transgas
	app.Ratio = ratios
	app.Dev = devres == 0
	app.Funder = funder
	app.Receiver = receiver
}

// LoadAllConfigs set various configs
func LoadAllConfigs(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	LoadApp()
}
