package brc20

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	Ticker      = "ordi"
	Start       = 0
	Limit       = 300
	BearerToken = "00000"
)

func CalculateBrc20Holders() {
	length := GetHoldersLength(Ticker)
	num := length.Data.HoldersCount / Limit
	value := length.Data.HoldersCount % Limit
	for i := 0; i <= num; i++ {
		var ticksHolders TickResBody
		if i == num {
			ticksHolders = GetHoldersInfo(Ticker, Start+i*Limit, value, BearerToken)
		} else {
			ticksHolders = GetHoldersInfo(Ticker, Start+i*Limit, Limit, BearerToken)
		}

		for _, info := range ticksHolders.Data.Detail {
			writeToCSV(HolderInfo{
				info.Address,
				info.OverallBalance,
				info.TransferableBalance,
				info.AvailableBalance,
				info.AvailableBalanceSafe,
				info.AvailableBalanceUnSafe,
			})
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func GetHoldersInfo(Ticker string, Start, Limit int, BearerToken string) TickResBody {
	url := fmt.Sprintf("https://open-api.unisat.io/v1/indexer/brc20/%s/holders?start=%d&limit=%d", Ticker, Start, Limit)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get request failed:", err)
		//return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error get:%d\n", resp.StatusCode)
		//return false
	}

	blocks := TickResBody{}
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		fmt.Println(err)
	}
	return blocks
}

func GetHoldersLength(Ticker string) Length {
	url := fmt.Sprintf("https://open-api.unisat.io/v1/indexer/brc20/%s/info", Ticker)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get request failed:", err)
		//return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error get:%d\n", resp.StatusCode)
		//return false
	}

	blocks := Length{}
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		fmt.Println(err)
	}
	return blocks
}

type TickResBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Height int `json:"height"`
		Total  int `json:"total"`
		Start  int `json:"start"`
		Detail []struct {
			Address                string `json:"address"`
			OverallBalance         string `json:"overallBalance"`
			TransferableBalance    string `json:"transferableBalance"`
			AvailableBalance       string `json:"availableBalance"`
			AvailableBalanceSafe   string `json:"availableBalanceSafe"`
			AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
		} `json:"detail"`
	} `json:"data"`
}

type Length struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Ticker                 string `json:"ticker"`
		SelfMint               bool   `json:"selfMint"`
		HoldersCount           int    `json:"holdersCount"`
		HistoryCount           int    `json:"historyCount"`
		InscriptionNumber      int    `json:"inscriptionNumber"`
		InscriptionId          string `json:"inscriptionId"`
		Max                    string `json:"max"`
		Limit                  string `json:"limit"`
		Minted                 string `json:"minted"`
		TotalMinted            string `json:"totalMinted"`
		ConfirmedMinted        string `json:"confirmedMinted"`
		ConfirmedMinted1H      string `json:"confirmedMinted1h"`
		ConfirmedMinted24H     string `json:"confirmedMinted24h"`
		MintTimes              int    `json:"mintTimes"`
		Decimal                int    `json:"decimal"`
		Creator                string `json:"creator"`
		Txid                   string `json:"txid"`
		DeployHeight           int    `json:"deployHeight"`
		DeployBlocktime        int    `json:"deployBlocktime"`
		CompleteHeight         int    `json:"completeHeight"`
		CompleteBlocktime      int    `json:"completeBlocktime"`
		InscriptionNumberStart int    `json:"inscriptionNumberStart"`
		InscriptionNumberEnd   int    `json:"inscriptionNumberEnd"`
	} `json:"data"`
}

type HolderInfo struct {
	Address                string
	OverallBalance         string
	TransferableBalance    string
	AvailableBalance       string
	AvailableBalanceSafe   string
	AvailableBalanceUnSafe string
}

func writeToCSV(info HolderInfo) error {
	// 打开 CSV 文件，如果文件不存在则创建文件
	file, err := os.OpenFile("holders.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个CSV写入器
	fileWriter := csv.NewWriter(file)
	defer fileWriter.Flush()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		// 写入表头
		header := []string{
			"Address",
			"OverallBalance",
			"TransferableBalance",
			"AvailableBalance",
			"AvailableBalanceSafe",
			"AvailableBalanceUnSafe",
		}
		if err := fileWriter.Write(header); err != nil {
			return err
		}
	}
	// 写入数据
	line := []string{
		info.Address,
		info.OverallBalance,
		info.TransferableBalance,
		info.AvailableBalance,
		info.AvailableBalanceSafe,
		info.AvailableBalanceUnSafe,
	}
	// 写入行
	if err := fileWriter.Write(line); err != nil {
		return err
	}
	return nil
}
