package uploadFile

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadFiles() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "./index.html")
			return
		}

		// Set the maximum file size to 20MB
		maxFileSize := int64(20 * 1024 * 1024)

		// Parse the multipart form
		if err := r.ParseMultipartForm(maxFileSize); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Retrieve the file from the form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Check if the file size is within the limit
		if handler.Size > maxFileSize {
			http.Error(w, "File size exceeds the maximum limit of 20MB", http.StatusBadRequest)
			return
		}

		// Create the uploads directory if it doesn't exist
		if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new file in the uploads directory
		dst, err := os.Create(filepath.Join("uploads", handler.Filename))
		//fmt.Println(handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the created file on the server
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if strings.HasSuffix(handler.Filename, ".csv") {
			path := filepath.Join("uploads", handler.Filename)
			readCSV(path)
		}

		// Return a success message to the client
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	})

	// Listen on port 8080
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func readCSV(filepath string) {
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

	//var quantities []string
	//var addressesStr []string

	for _, line := range lines {
		address := strings.TrimSpace(line[0])
		if address == "" {
			break
		}
		fmt.Println(address)
		quantity := strings.TrimSpace(line[1])
		if quantity == "" {
			fmt.Println("没有提供数量:", address)
			panic(err)
		}
		fmt.Println(quantity)
	}

}
