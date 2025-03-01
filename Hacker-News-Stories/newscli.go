package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type PostIDs struct {
	IDs []int `json:"ids"`
}

type Story struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// fetchJSON makes an HTTP GET request and unmarshals the response into the target.
func fetchJSON(url string, target any) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to fetch data: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	return nil
}

// getIDs fetches the list of story IDs from the Hacker News API.
func getIDs() ([]int, error) {
	url := "https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty"

	var ids []int
	if err := fetchJSON(url, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

// getStory fetches a single story by its ID from the Hacker News API.
func getStory(id int) (Story, error) {
	url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(id) + ".json?print=pretty"

	var story Story
	if err := fetchJSON(url, &story); err != nil {
		return Story{}, err
	}

	return story, nil
}

// fetchStories concurrently fetches stories for the given IDs.
func fetchStories(ids []int) ([]Story, error) {
	var wg sync.WaitGroup
	stories := make([]Story, len(ids))
	errors := make(chan error, len(ids))

	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			story, err := getStory(id)
			if err != nil {
				errors <- fmt.Errorf("error fetching story %d: %v", id, err)
				return
			}
			stories[i] = story
		}(i, id)
	}

	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return stories, nil
}

func main() {
	totalIds := flag.Int("total", 10, "total ids")

	flag.Parse()
	// Fetch story IDs
	ids, err := getIDs()
	if err != nil {
		log.Fatalf("Error fetching IDs: %v", err)
	}

	// Fetch stories concurrently
	stories, err := fetchStories(ids[:*totalIds]) // Limit to the number of stories specified by the user
	if err != nil {
		log.Fatalf("Error fetching stories: %v", err)
	}

	// Print stories
	for _, story := range stories {
		fmt.Println("By: ", story.By)
		fmt.Println("Id: ", story.ID)
		fmt.Println("Score: ", story.Score)
		fmt.Println("Title: ", story.Title)
		fmt.Println("URL: ", story.URL)
		fmt.Println()
	}
}