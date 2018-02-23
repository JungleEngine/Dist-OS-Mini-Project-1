// gets top 10 most followed on twitter dataset.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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

	//creating user object
	var users SortUser
	//current number of users
	users.Followers = map[int]int{}

	visitedUsers := map[int]bool{}
	// filling follower map
	// read from input file
	file, _ := os.Open(dataInput)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//string to int
		id_1, _ := strconv.Atoi(words[0])
		id_2, _ := strconv.Atoi(words[1])
		users.Followers[id_2]++
		// check if user is not added to keys , then add it
		if !visitedUsers[id_1] {
			visitedUsers[id_1] = true
			users.Keys = append(users.Keys, id_1)

		}
		if !visitedUsers[id_2] {
			visitedUsers[id_2] = true
			users.Keys = append(users.Keys, id_2)

		}
	}
	fmt.Println(len(users.Followers))

	// TODO: Implement topTen function.
	//TODO: 1- read file 2- fill follower map 3-fill keys array
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
	topId := topTen("dataset.txt")
	fmt.Println("topTen length: ", len(topId))

	fmt.Println("Getting and printing screen name for top 10...")
	for i := 0; i < 10 && i < len(topId); i++ {
		fmt.Printf("%-15d%s\n", topId[i], getUsername(strconv.Itoa(topId[i])))
	}
}
