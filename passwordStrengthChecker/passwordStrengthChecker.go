package main

import (
	"log"
	"os"
	"strings"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        log.Fatal("Specify a password")
    }
    if len(args[1]) < 9 {
        log.Println("Password is so short")
    } else if !(strings.IndexFunc(args[1], letters) >= 0){
        log.Println("Password must containt lower case (at least one)")
    }  else if !(strings.IndexFunc(args[1], capitals) >= 0){
        log.Println("Password must containt upper case (at least one)")
    } else if !strings.ContainsAny(args[1], "!@#$%&*") {
        log.Println("Password must contain one of these '!@#$%&*'")
    } else {
        log.Println("Password is perfecto")
    }
}

func letters(r rune) bool {
        if (r >= 97 && r <= 122) {
            return true
        }
        return false
    }

func capitals(r rune) bool {
    if (r >= 65 && r <= 90) {
        return true
    } else { return false }
}
