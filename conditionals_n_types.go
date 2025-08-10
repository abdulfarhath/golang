package main

import "fmt"

func main() {
    // Example 1: if statement
    age := 12

    // The 'if' statement checks a condition. If it's true, the code inside runs.
    if age > 18 {
        fmt.Println("You are an adult.")
    } else if age == 18 {
        // 'else if' lets you check another condition if the first was false.
        fmt.Println("You are exactly 18 years old.")
    } else {
        // 'else' runs if none of the above conditions were true.
        fmt.Println("You are a child.")
    }

    // Example 2: switch statement
    // 'switch' is used to compare a value against multiple cases.
    switch age {
    case 18:
        fmt.Println("You are exactly 18 years old.")
    case 19, 20:
        // You can list multiple values in one case.
        fmt.Println("You are 19 or 20 years old.")
    default:
        // 'default' runs if no case matches.
        fmt.Println("Your age does not match any special case.")
    }

    // Example 3: switch without an expression (acts like if-else chain)
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    case score >= 70:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: D or below")
    }

    // Example 4: short statement in if
    if n := 5; n%2 == 0 {
        fmt.Println("n is even")
    } else {
        fmt.Println("n is odd")
    }
}
