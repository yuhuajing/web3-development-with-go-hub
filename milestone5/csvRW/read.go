package csvrw

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ReadCSV(filepath string) {
	// Open CSV file
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		panic(err)
	}
	defer f.Close()

	// Read File into *lines* variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("无法读取CSV文件:", err)
		panic(err)
	}

	// 创建地址数组和数量数组
	var quantities []string
	var addressesStr []string
	// 用于检查地址是否已经存在的映射
	seenAddresses := make(map[string]bool)
	// 遍历每一行记录

	for _, line := range lines {
		address := strings.TrimSpace(line[0])
		if address == "" {
			break
		}

		quantity := strings.TrimSpace(line[1])
		if quantity == "" {
			fmt.Println("没有提供数量:", address)
			panic(err)
		}

		if len(address) != 42 {
			fmt.Println("地址长度错误:", address)
			panic(err)
		}
		// 检查地址是否已存在
		if seenAddresses[address] {
			fmt.Println("重复的地址 :", address)
			panic(err)
		}

		addressesStr = append(addressesStr, address)
		quantities = append(quantities, quantity)
		// 标记地址已存在
		seenAddresses[address] = true
	}

	//fmt.Println("[\"" + strings.Join(addressesStr, "\",\n\"") + "\"]")
	receiverdata := "[\"" + strings.Join(addressesStr, "\",\"") + "\"]"
	writeTxt("./addr.txt", receiverdata)
	//fmt.Println("[" + strings.Join(quantities, ",\n") + "]")
	nftAmount := "[" + strings.Join(quantities, ",") + "]"
	writeTxt("./xxx.txt", nftAmount)
}

func ReadBackUpCSV(filepath string) {
	// Open CSV file
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		panic(err)
	}
	defer f.Close()

	// Read File into *lines* variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("无法读取CSV文件:", err)
		panic(err)
	}

	// 创建地址数组和数量数组
	var addressesStr []string
	// 用于检查地址是否已经存在的映射
	seenAddresses := make(map[string]bool)
	// 遍历每一行记录

	for _, line := range lines {
		address := strings.TrimSpace(line[0])
		if len(address) != 42 {
			fmt.Println("地址长度错误:", address)
			panic(err)
		}
		// 检查地址是否已存在
		if seenAddresses[address] {
			fmt.Println("重复的地址 :", address)
			panic(err)
		}

		addressesStr = append(addressesStr, address)
		// 标记地址已存在
		seenAddresses[address] = true
	}

	//fmt.Println("[\"" + strings.Join(addressesStr, "\",\n\"") + "\"]")
	receiverdata := "[\"" + strings.Join(addressesStr, "\",\"") + "\"]"
	writeTxt("./xxx.txt", receiverdata)
}

func writeTxt(fileTxt, data string) {
	file, err := os.Create(fileTxt)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 创建写入器
	writer := bufio.NewWriter(file)

	// 写入数据
	_, err = writer.WriteString(data)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	// 刷新缓冲区，确保所有数据都被写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区时出错:", err)
		return
	}
	fmt.Println("数据已成功写入文件。")
}
