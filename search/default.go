package search

import "log"

type defaultMatcher struct {
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Match already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
