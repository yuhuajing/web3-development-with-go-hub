package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Symbol string

const (
	SymbolETH     Symbol = "ETH"
	SymbolIMX     Symbol = "IMX"
	SymbolSEPOLIA Symbol = "SEPOLIA"
	SymbolRONIN   Symbol = "RONIN"
	SymbolPOLYGON Symbol = "POLYGON"
	SymbolBNB     Symbol = "BNB"
	SymbolNEAR    Symbol = "NEAR"
)

var _SymbolValue = map[Symbol]Symbol{
	"RONIN":   SymbolRONIN,
	"IMX":     SymbolIMX,
	"ETH":     SymbolETH,
	"POLYGON": SymbolPOLYGON,
	"BNB":     SymbolBNB,
	"NEAR":    SymbolNEAR,
	"SEPOLIA": SymbolSEPOLIA,
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	} // The Original .env
	//env := Env("ENV", "")
	_ = godotenv.Overload(".env")
}

func getRpcPool(c Symbol) []string {
	switch _SymbolValue[c] {
	case SymbolETH:
		return EthRpcPool()
	case SymbolSEPOLIA:
		return SepoliaRpcPool()
	case SymbolIMX:
		return IMXRpcPool()
	case SymbolRONIN:
		return RoninRpcPool()
	case SymbolPOLYGON:
		return PolygonRpcPool()
	case SymbolBNB:
		return BnbRpcPool()
	case SymbolNEAR:
		return NearRpcPool()
	default:
		return nil
	}
}

func IMXRpcPool() []string {
	return rpcPool(Env("IMX_RPC_POOL", Env("ETH_RPC_URL", "")))
}

func EthRpcPool() []string {
	return rpcPool(Env("ETH_RPC_POOL", Env("ETH_RPC_URL", "")))
}
func SepoliaRpcPool() []string {
	return rpcPool(Env("SEPOLIA_RPC_POOL", Env("SEPOLIA_RPC_POOL", "")))
}
func RoninRpcPool() []string {
	return rpcPool(Env("RONIN_RPC_POOL", Env("RONIN_RPC_POOL", "")))
}
func PolygonRpcPool() []string {
	return rpcPool(Env("POLYGON_RPC_POOL", Env("POLYGON_RPC_POOL", "")))
}
func BnbRpcPool() []string {
	return rpcPool(Env("BNB_RPC_POOL", Env("BNB_RPC_POOL", "")))
}
func NearRpcPool() []string {
	return rpcPool(Env("NEAR_RPC_POOL", Env("NEAR_RPC_POOL", "")))
}

func Env(key string, d string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return d
}

func rpcPool(pool string) []string {
	var ret []string
	for _, url := range strings.Split(pool, ",") {
		url = strings.TrimSpace(url)
		if url != "" {
			ret = append(ret, url)
		}
	}

	return ret
}

func NewClient(s Symbol) *ethclient.Client {
	urls := getRpcPool(s)
	if urls == nil {
		fmt.Printf("eth.NewClient: urls not found")
		return nil
	}

	var clients *ethclient.Client
	for index, url := range urls {
		client, err := ethclient.Dial(url)
		if err == nil {
			clients = client
			break
		} else {
			if index == len(urls)-1 {
				fmt.Printf("eth.NewClient: all RPC urls invalid")
				return nil
			}
			continue
		}
	}
	return clients
}
