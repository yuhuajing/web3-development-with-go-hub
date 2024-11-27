package milestone5

import (
	"bufio"
	"fmt"
	"os"
)

func WriteTxt(filePath, data string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
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
