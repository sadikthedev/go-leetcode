package main

import "strings"

//Find the Index of the first occurence in a string

func IndexOfTheFirstOccurenceInAString(haystack string, needle string) int{
	return strings.Index(haystack, needle)
}