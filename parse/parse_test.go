package parse

import (
	"testing"
)

// TestParseNewsContentInvalidJSON tests the ParseNewsContent function when passed an invalid JSON string.
func TestParseNewsContentInvalidJSON(t *testing.T) {
	jsonData := "invalid json"
	content, err := ParseNewsContent(jsonData)
	if content != "" || err == nil {
		t.Error("Expected an error when parsing invalid JSON data")
	}
}

// TestParseNewsContentEmptyString tests the ParseNewsContent function when passed an empty string.
func TestParseNewsContentEmptyString(t *testing.T) {
	jsonData := ""
	content, err := ParseNewsContent(jsonData)
	if content != "" || err == nil {
		t.Error("Expected an error when parsing an empty string")
	}
}

// TestParseNewsContentGivenString tests the ParseNewsContent function when passed valid JSON data.
func TestParseNewsContentGivenString(t *testing.T) {
	jsonData := "{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.layout\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.appName\", \"ContentDescription\": \"Cool app\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.Username\", \"ContentDescription\": \"username\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.logout\", \"ContentDescription\": \"Logout\", \"children\": []}]}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece\", \"ContentDescription\": \"Amazing News!!\", \"children\":[{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece.content\", \"ContentDescription\": \"Today at the NIO, we have great news\", \"children\":[]}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece.content\", \"ContentDescription\": \"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vitae elit sed nunc accumsan finibus id id metus. Sed elit.\", \"children\":[]}]}]}"
	content, err := ParseNewsContent(jsonData)
	if content != "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vitae elit sed nunc accumsan finibus id id metus. Sed elit." || err != nil {
		t.Error("Expected the newspiece content to be returned")
	}
}

// TestParseNewsContentNoNewsPiece tests the ParseNewsContent function when passed valid JSON data with no newspiece.
func TestParseNewsContentNoNewsPiece(t *testing.T) {
	jsonData := "{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.layout\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.appName\", \"ContentDescription\": \"Cool app\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.Username\", \"ContentDescription\": \"username\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.logout\", \"ContentDescription\": \"Logout\", \"children\": []}]}]}"
	content, err := ParseNewsContent(jsonData)
	if content != "" || err == nil {
		t.Error("Expected an error when no newspiece element is found")
	}
}

// TestParseNewsContentNoNewsPieceChildren tests the ParseNewsContent function when passed valid JSON data with no
// newspiece children.
func TestParseNewsContentNoNewsPieceChildren(t *testing.T) {
	jsonData := "{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.layout\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu\", \"ContentDescription\": \"null\", \"children\": [{\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.appName\", \"ContentDescription\": \"Cool app\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.Username\", \"ContentDescription\": \"username\", \"children\": []}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.topMenu.logout\", \"ContentDescription\": \"Logout\", \"children\": []}]}, {\"ActivityName\": \"org.nio.MainActivity\", \"ElementName\": \"org.nio.newspiece\", \"ContentDescription\": \"Amazing News!!\", \"children\":[]}]}"
	content, err := ParseNewsContent(jsonData)
	if content != "" || err == nil {
		t.Error("Expected an error when no newspiece element is found")
	}
}
