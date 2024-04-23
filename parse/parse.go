package parse

import (
	"encoding/json"
	"errors"
)

// Activity is used to represent a JSON activity object
type Activity struct {
	ActivityName       string     `json:"ActivityName"`
	ElementName        string     `json:"ElementName"`
	ContentDescription string     `json:"ContentDescription"`
	Children           []Activity `json:"children"`
}

// ParseNewsContent takes a JSON representation of an activity and returns the
// innermost content description which contains newspiece content.
func ParseNewsContent(jsonData string) (string, error) {
	// Unmarshal the JSON data into an Activity struct so we can access the relevant fields
	var activity Activity
	err := json.Unmarshal([]byte(jsonData), &activity)
	if err != nil {
		return "", err
	}

	// Create an error message for the case where the newspiece element is not found
	retErr := errors.New("No newspiece element found in the activity")

	// Iterate through the children of the activity
	for _, child := range activity.Children {
		// Find the newspiece element
		if child.ElementName == "org.nio.newspiece" {
			// Handle the case where the newspiece element has no children
			if len(child.Children) != 0 {
				// Return the content description of the last newspiece content element
				return child.Children[len(child.Children)-1].ContentDescription, nil
			} else {
				// Set the error message to indicate that no newspiece content was found
				retErr = errors.New("No newspiece content found in the activity")
			}
		}
	}

	// Return an empty string and the error message if the newspiece content was not found
	return "", retErr
}
