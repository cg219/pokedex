package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cg219/pokedex/internal/pokeapi"
	"github.com/charmbracelet/log"
)

type command struct {
    name string
    desc string
    cb func() error
}

var nextL pokeapi.LocationQuery
var prevL pokeapi.LocationQuery

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
        "map": {
            name: "map",
            desc: "load the next 10 maps",
            cb: mapCommand(true),
        },
        "mapb": {
            name: "mapb",
            desc: "load the prev 10 maps",
            cb: mapCommand(false),
        },
    }
}

func mapCommand(next bool) func() error {
    lq := pokeapi.LocationQuery{}

    if next && nextL != (pokeapi.LocationQuery{}) {
        lq = nextL
    }

    if !next && prevL != (pokeapi.LocationQuery{}) {
        lq = prevL
    }

    return func() error {
        locs, nl, pl, err := pokeapi.GetLocation(lq)

        nextL = nl
        prevL = pl

        for _, l := range locs {
            fmt.Println(l.Name)
        }

        if err != nil {
            log.Error(err)    
        }

        return nil

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
