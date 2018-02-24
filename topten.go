// gets top 10 most followed on twitter dataset.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

// Implementing sort interface.
type SortUser struct {
	Followers map[int]int
	Keys      []int
}

// Gets length of map.
func (su *SortUser) Len() int {
	return len(su.Followers)
}

// Condition for sorting to compare between values of keys.
func (su *SortUser) Less(i, j int) bool {
	return su.Followers[su.Keys[i]] > su.Followers[su.Keys[j]]
}

// Swaps two keys in keys array.
func (su *SortUser) Swap(i, j int) {
	su.Keys[i], su.Keys[j] = su.Keys[j], su.Keys[i]
}

// Sorts Keys based on number of followers in descending order.
func sortKeys(m map[int]int) []int {
	su := new(SortUser)
	su.Followers = m
	su.Keys = make([]int, len(m)+1)
	i := 0
	for key, _ := range m {
		su.Keys[i] = key
		i++
	}
	sort.Sort(su)
	return su.Keys
}

// Calculates top 10 most followed for input file
// and returns array of user id (int) for top 10.
func topTen(dataInput string) []int {

	//creating user object
	var users SortUser
	//current number of users
	users.Followers = map[int]int{}

	// filling follower map
	// read from input file
	file, err := os.Open(dataInput)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//string to int
		id_2, _ := strconv.Atoi(words[1])
		users.Followers[id_2]++
		// check if user is not added to keys , then add it

	}
	//sort keys of users
	users.Keys = sortKeys(users.Followers)

	//return top 10 (or less than 10 if the followed peoples < 10)
	return users.Keys[0:int64(math.Min(float64(len(users.Keys)),float64(10)))]
}

// Connects to remote service through internet to convert user id
// to username.
func getUsername(userId string) string {

	response, err := http.PostForm("https://tweeterid.com/ajax.php",
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
		//fmt.Println("error -> ", body)
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
