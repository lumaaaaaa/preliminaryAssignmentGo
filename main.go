package main

import (
	"fmt"
	"log"

	"preliminaryAssignmentGo/parse"
)

func main() {
	// The provided JSON example from the assignment
	jsonData := "{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.layout\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.appName\", \"ContentDescription\": \"Cool app\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.Username\", \"ContentDescription\": \"username\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.logout\", \"ContentDescription\": \"Logout\", \"children\": []}]}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece\", \"ContentDescription\": \"Amazing News!!\", \"children\":[{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece.content\", \"ContentDescription\": \"Today at the NIO, we have great news\", \"children\":[]}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece.content\", \"ContentDescription\": \"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vitae elit sed nunc accumsan finibus id id metus. Sed elit.\", \"children\":[]}]}]}"

	// Parse the JSON data and return the content
	content, err := parse.ParseNewsContent(jsonData)
	if err != nil {
		// If there's an error, log it and exit
		log.Fatal(err)
	}

	// Print the content
	fmt.Println("The returned content is:", content)
}
