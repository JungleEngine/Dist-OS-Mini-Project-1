// gets top 10 most followed on twitter dataset.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Implementing sort interface.
type SortUser struct {
	Followers map[int]int
	Keys      []int
}

// Gets length of map.
func (su *SortUser) Len() int {
	// TODO: Implement Len function.
	return 0
}

// Condition for sorting to compare between values of keys.
func (su *SortUser) Less(i, j int) bool {
	// TODO: Implement Less function.
	return false
}

// Swaps two keys in keys array.
func (su *SortUser) Swap(i, j int) {
	// TODO: Implement Swap function.
}

// Sorts Keys based on number of followers in descending order.
func sortKeys(m map[int]int) []int {
	// TODO: Implement sortKeys function.
	return nil
}

// Calculates top 10 most followed for input file
// and returns array of user id (int) for top 10.
func topTen(dataInput string) []int {
	// TODO: Implement topTen function.
	return nil
}

// Connects to remote service through internet to convert user id
// to username.
func getUsername(userId string) string {
	response, err := http.PostForm("http://tweeterid.com/ajax.php",
		url.Values{"input": {userId}})
	if err != nil {
		fmt.Println("Error getting username("+userId+"): ", err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in response("+userId+"): ", err)
		return ""
	}
	if string(body) == "error" {
		return ""
	}
	return string(body)
}

func main() {
	fmt.Println("Calculating top 10 most followed...")
	topId := topTen("input.txt")
	fmt.Println("topTen length: ", len(topId))

	fmt.Println("Getting and printing screen name for top 10...")
	for i := 0; i < 10 && i < len(topId); i++ {
		fmt.Printf("%-15d%s\n", topId[i], getUsername(strconv.Itoa(topId[i])))
	}
}
