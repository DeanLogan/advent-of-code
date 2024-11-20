package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Game struct {
	Hand string
	Bid  int
	Type string
}

var charToValueMapPartOne = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var charToValueMapPartTwo = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func part1(){
    // create a slice for each type of hand
    fiveOfAKinds := []Game{}
    fourOfAKinds := []Game{}
    fullHouse := []Game{}
    threeOfAKind := []Game{}
    twoPair := []Game{}
    onePair := []Game{}
    highCard := []Game{}

    // Convert the .txt into a slice of games in the format of the Game struct
    ans := 0
    lines := libs.FileToSlice("2023/day07/input.txt", "\n")
    games := []Game{}
    for i:=0; i<len(lines);i++{
        game := Game{}
        bidStr := ""
        game.Hand, bidStr = libs.SplitAtChar(lines[i], ' ')
        bidStr = strings.ReplaceAll(bidStr, " ", "")
        game.Hand = strings.ReplaceAll(game.Hand, " ", "")
        game.Bid, _ = strconv.Atoi(bidStr) 
        game.Type = calcType(game.Hand)

        // get the number of occurances for each character in the hand
        count := map[string]int{}
        for _, char := range game.Hand {
            count[string(char)]++
        }

        // store only the number of times a value occurs for the values
        values := []int{}
        for _, value := range count {
            values = append(values, value)
        }
        sort.Ints(values)

        // adds the game to the correct type slice
        switch len(values) {
        case 1:
            fiveOfAKinds = append(fiveOfAKinds, game)
        case 2:
            if values[1] == 4 {
                fourOfAKinds = append(fourOfAKinds, game)
            } else {
                fullHouse = append(fullHouse, game)
            }
        case 3:
            if values[2] == 3 {
                threeOfAKind = append(threeOfAKind, game)
            } else {
                twoPair = append(twoPair, game)
            }
        case 4:
            onePair = append(onePair, game)
        default:
            highCard = append(highCard, game)
        }

        // games = append(games, game)
    } 

    // order each of the types
    fiveOfAKinds = orderHands(fiveOfAKinds, charToValueMapPartOne)
    fourOfAKinds = orderHands(fourOfAKinds, charToValueMapPartOne)
    fullHouse = orderHands(fullHouse, charToValueMapPartOne)
    threeOfAKind = orderHands(threeOfAKind, charToValueMapPartOne)
    twoPair = orderHands(twoPair, charToValueMapPartOne)
    onePair = orderHands(onePair, charToValueMapPartOne)
    highCard = orderHands(highCard, charToValueMapPartOne)

    // games = orderHands(games)

    games = append(games, highCard...)
    games = append(games, onePair...)
    games = append(games, twoPair...)
    games = append(games, threeOfAKind...)
    games = append(games, fullHouse...)
    games = append(games, fourOfAKinds...)
    games = append(games, fiveOfAKinds...)

    //game.Hand, bidStr = libs.SplitAtChar(lines[i], ' ')

    for i := 0; i<len(games); i++{
        ans = ans + (games[i].Bid * (i+1))
    }

    fmt.Println("The answer to part 1 for day 7 is:", ans)
}

func orderHands(games []Game, valueMap map[byte]int) []Game {
    sort.Slice(games, func(i, j int) bool {
        game1, game2 := games[i], games[j]
        hand1, hand2 := game1.Hand, game2.Hand

        for k := 0; k < len(hand1); k++ {
            card1, card2 := hand1[k], hand2[k]
            value1, value2 := valueMap[card1], valueMap[card2]

            if value1 != value2 {
                return value1 < value2
            }
        }

        return false
    })

    return games
}

func calcType(hand string) string {
    // get the number of occurances for each character in the hand
    count := map[string]int{}
    for _, char := range hand{
        count[string(char)]++
    }

    // store only the number of times a value occurs for the values
    values := []int{}
    for _, value := range count {
        values = append(values, value)
    }
    sort.Ints(values)

    // check the type based on how man
    switch len(values) {
    case 1:
        return "Five of a kind"
    case 2:
        if values[1] == 4 {
            return "Four of a kind"
        } else {
            return "Full house"
        }
    case 3:
        if values[2] == 3 {
            return "Three of a kind"
        } else {
            return "Two pair"
        }
    case 4:
        return "One pair"
    default:
        return "High card"
    }
}


func part2(){
    // create a slice for each type of hand
    fiveOfAKinds := []Game{}
    fourOfAKinds := []Game{}
    fullHouse := []Game{}
    threeOfAKind := []Game{}
    twoPair := []Game{}
    onePair := []Game{}
    highCard := []Game{}

    // Convert the .txt into a slice of games in the format of the Game struct
    ans := 0
    lines := libs.FileToSlice("2023/day07/input.txt", "\n")
    games := []Game{}
    for i:=0; i<len(lines);i++{
        game := Game{}
        bidStr := ""
        game.Hand, bidStr = libs.SplitAtChar(lines[i], ' ')
        bidStr = strings.ReplaceAll(bidStr, " ", "")
        game.Hand = strings.ReplaceAll(game.Hand, " ", "")
        game.Bid, _ = strconv.Atoi(bidStr) 
        game.Type = calcType(game.Hand)

        // get the number of occurances for each character in the hand
        count := map[string]int{}
        for _, char := range game.Hand {
            count[string(char)]++
        }

        // store only the number of times a value occurs for the values
        values := []int{}
        for char, value := range count {
            if char != "J"{ // records the highest occurance character that is not J
                values = append(values, value)
            }
        }
        sort.Ints(values)

        if (len(values) > 0){
            values[len(values)-1] = values[len(values)-1]+count["J"] // adds the occurances of J to the highest character to improve the type
        } else { // handles the edge case where all values in the hand is a J
            values = []int{0}
        }

        // adds the game to the correct type slice
        switch len(values) {
        case 1:
            fiveOfAKinds = append(fiveOfAKinds, game)
        case 2:
            if values[1] == 4 {
                fourOfAKinds = append(fourOfAKinds, game)
            } else {
                fullHouse = append(fullHouse, game)
            }
        case 3:
            if values[2] == 3 {
                threeOfAKind = append(threeOfAKind, game)
            } else {
                twoPair = append(twoPair, game)
            }
        case 4:
            onePair = append(onePair, game)
        default:
            highCard = append(highCard, game)
        }
    } 

    // order each of the types
    fiveOfAKinds = orderHands(fiveOfAKinds, charToValueMapPartTwo)
    fourOfAKinds = orderHands(fourOfAKinds, charToValueMapPartTwo)
    fullHouse = orderHands(fullHouse, charToValueMapPartTwo)
    threeOfAKind = orderHands(threeOfAKind, charToValueMapPartTwo)
    twoPair = orderHands(twoPair, charToValueMapPartTwo)
    onePair = orderHands(onePair, charToValueMapPartTwo)
    highCard = orderHands(highCard, charToValueMapPartTwo)

    games = append(games, highCard...)
    games = append(games, onePair...)
    games = append(games, twoPair...)
    games = append(games, threeOfAKind...)
    games = append(games, fullHouse...)
    games = append(games, fourOfAKinds...)
    games = append(games, fiveOfAKinds...)

    for i := 0; i<len(games); i++{
        ans = ans + (games[i].Bid * (i+1))
    }

    fmt.Println("The answer to part 2 for day 7 is:", ans)
}