package main
import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }

    switch age {
    case 18:
        fmt.Println("Just turned adult!")
    default:
        fmt.Println("Different age")
    }
}

