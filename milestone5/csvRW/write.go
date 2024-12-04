package csvrw

import (
	"encoding/csv"
	"os"
)

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
	file, err := os.OpenFile("xxx.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
