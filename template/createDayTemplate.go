package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code/template/webScraping"
)

// Create main function to execute the program
func main() {
	year := getinput("Enter the year that the template will be created in: ")
	if year == "" {
		return 
	}

	day := getinput("Enter the number of the day whos template you want to create: ")
	if day == ""{
		return
	}
	dayInt, err := strconv.Atoi(day)
    if err != nil {
        fmt.Println("Invalid day input")
        return
    }

	inputData := webScraping.GetWebScrapedData(year, day, true)
	if inputData == "" {
		fmt.Println("Data not available for selected year and date")
		return
	}

	htmlContent := webScraping.GetWebScrapedData(year, day, false)
	readmeFileContent := webScraping.HtmlToReadme(htmlContent, year, day)+"\n\n## Part 2"

	if dayInt < 10 {
        day = "0" + day
    }

	err = createFolder(year)
	if err != nil {
		fmt.Println("Problem with creating year folder")
		return
	}
	

	err = createFolder(year+"\\day"+day)
	if err != nil {
		fmt.Println("Problem with creating day folder")
		return
	}

	folderName := year+"\\day"+day
	
	writeToFile(folderName, "input.txt", inputData)

	writeToFile(folderName, "README.md", readmeFileContent)

	goFileContent := "package main\n\nimport (\n    \"fmt\"\n\n    \"github.com/DeanLogan/advent-of-code/libs\"\n)\n\nfunc main(){\n    part1()\n    part2()\n}\n\nfunc part1(){\n    ans := 0\n    fmt.Println(\"The answer to part 1 for day "+day+" is:\", ans)\n}\n\nfunc part2(){\n    ans := 0\n    fmt.Println(\"The answer to part 2 for day "+day+" is:\", ans)\n}"

	writeToFile(folderName, "day"+day+".go", goFileContent)
}

func getinput(msg string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading user input:", err)
		return ""
	}

	return strings.TrimSpace(userInput)
}

func createFolder(folderName string) error {
    err := os.Mkdir(folderName, 0755) // 0755 is the permission mode
    if err != nil {
        if os.IsExist(err) {
            fmt.Println("Folder already exists:", folderName)
            return nil
        }
        fmt.Println("Error creating folder:", err)
        return err
    }

    fmt.Println("Folder", folderName, "created successfully.")
    return nil
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