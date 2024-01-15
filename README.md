# Advent Of Code 2023 in Go!

Advent of code is a Advent calendar of programming puzzles, this features my solutions to the programming puzzles presented in the Advent of Code, a festive calendar of intriguing challenges. I've implemented these solutions using the Go programming language.

As I found myself quite busy during december due to some exams/courswork at university I decided not to take part while the challanges were being realeased, but as those exams are finished now (thank god :)), I've finally found the time to delve into the Advent of Code puzzles. Given my limited experience with Go – just a single course under my belt – I saw this as an great opportunity to strengthen my knowledge of Go. 

Visit [https://adventofcode.com/](https://adventofcode.com/) to find out more about advent of code

# Repo Structure

- Each day (puzzle) is seperated into its corresponding folder, there is a single go.mod file so to run the code for each day you can simple cd into the root of the advent-of-code-2023 folder and enter the following:

```console
go run ./day{number of the day you want to see}
```

- The libs folder contains functions that are used across various days, e.g. functions for reading .txt files, string manipulation, etc.

- The template folder contains code that generates the template that I use for each day, it will create a folder for the day selected, along with a go file with some template code, a README.md and then a input.txt file which will connect to the corresponding day on advent of code 2023 and scrap the input data and place it into the input.txt file. To run create a template enter the following into the console:

```console
go run ./template
```

Then you will be prompted with the day you want to create the template for.