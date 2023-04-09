// Write a program that takes a list of strings as input and returns the longest string in the list.

package main
import (
    "fmt"
)
// findLongestString takes a list of strings and returns the longest string in the list.
func findLongestString(list []string) string {
    var longest string
    for _, str := range list {
        if len(str) > len(longest) {
            longest = str
        }
    }
    return longest
}

func main() {
    // Define a list of strings.
    stringsList := []string{"Virginia", "Alaska", "Georgia", "Ohio", "Maine"}

    // Call findLongestString to get the longest string in the list.
    longestString := findLongestString(stringsList)

    // Print the longest string to the console.
    fmt.Println("The longest string is:", longestString)
}