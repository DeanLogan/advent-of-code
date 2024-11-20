package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/DeanLogan/advent-of-code/template/webScraping"
)

// Create main function to execute the program
func main() {
	day := getinput()

	if day == ""{
		return
	}

	folderName := createFolder(day)
	createFile(folderName, "README.md")
	
	writeToFile(folderName, "input.txt", webScraping.GetWebScrapedData(day+"/input"))

	goFileContent := "package main\n\nimport (\n    \"fmt\"\n)\n\nfunc main(){\n    part1()\n    part2()\n}\n\nfunc part1(){\n    ans := 0\n    fmt.Println(\"The answer to part 1 for day "+day+" is:\", ans)\n}\n\nfunc part2(){\n    ans := 0\n    fmt.Println(\"The answer to part 2 for day "+day+" is:\", ans)\n}"

	writeToFile(folderName, folderName+".go", goFileContent)
}

func getinput() string{
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of the day whos template you want to create: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading user input:", err)
		return ""
	}

	return strings.TrimSpace(userInput)
}

func createFolder(day string) string {
	folderName := "day"+day

	// Create the folder
	err := os.Mkdir(folderName, 0755) // 0755 is the permission mode
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return ""
	}

	fmt.Println("Folder", folderName, "created successfully.")
	return folderName
}

func writeToFile(folder string, fileName string, content string) {	
	err := os.WriteFile(folder+"/"+fileName, []byte(content), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File is created successfully.")
}

func createFile(folder string, fileName string){
	file, err := os.Create(folder+"/"+fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("File is created successfully.")  
}