package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

type command struct {
    name string
    desc string
    cb func() error
}

func getCommandList() map[string]command {
    return map[string]command{
        "help": {
            name: "help",
            desc: "displays help guide",
            cb: helpCommand,
        },
        "exit": {
            name: "exit",
            desc: "close pokedex",
            cb: exitCommand,
        },
    }
}

func exitCommand() error {
    return errors.New("exit")
}

func helpCommand() error {
    fmt.Println("Welcome to the Pokedex!!!")
    fmt.Println("Usage:")
    fmt.Println("")

    for _, c := range getCommandList() {
        fmt.Printf("%s: %s\n", c.name, c.desc)
    }
    fmt.Println("")
    return nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Pokedex > ")

        c, err := reader.ReadString('\n')

        if err != nil {
            panic(err)
        }

        c = strings.TrimRight(c, "\n")

        cmd, ok := getCommandList()[c]

        if !ok {
            fmt.Println("unknown command")
            continue
        }

        err = cmd.cb()

        if err != nil {
            if err.Error() == "exit" {
                return
            }

            panic(err)
        }
    }
}
