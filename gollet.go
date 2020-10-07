package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "github.com/hilgardvr/gollet/create"
)

func printOptions() {
    fmt.Println("\nWelcome to Gollet\n")
    fmt.Println("Do you want to:")
    fmt.Println("\t1. Create a wallet")
    fmt.Println("\t2. Import a wallet")
    fmt.Println("\t3. Clear Data and Quit")
    fmt.Print("\tPlease select your choice -> ")
}

func createGollet() {
    fmt.Println("createing a wallet")
    create.CreateGollet()
    return
}

func importGollet() {
    fmt.Println("imporint a wallet")
    return
}

func cleanAndQuit() {
    fmt.Println("Cheerio")
    return
}

func readUserChoice() (string, error) {
    reader:= bufio.NewReader(os.Stdin)
    text, err:= reader.ReadString('\n')
    trimText:= strings.TrimSpace(text)
    return trimText, err
}

func printAndGetChoice() {
    var choice string
    var err error

    printOptions()
    choice, err = readUserChoice()

    if err != nil {
        fmt.Println("Error reading your input:", err)
        choice, err = readUserChoice()
    }

    if "1" == choice {
        createGollet()
        printAndGetChoice()
    } else if "2" == choice {
        importGollet()
        printAndGetChoice()
    } else if "3" == choice {
        cleanAndQuit()
        return
    }

    return
}

func main() {
    printAndGetChoice()
    return
}
