package htmldoc

import (
	"golang.org/x/net/html"
)

// From attrs extract single attr
func GetAttr(attrs []html.Attribute, key string) string {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

// From attrs extract the asked for keys
func ExtractAttrs(attrs []html.Attribute, keys []string) map[string]string {
	attrMap := make(map[string]string)
	for _, attr := range attrs {
		for i, key := range keys {
			if attr.Key == key {
				attrMap[key] = attr.Val
				// Delete the key from the keys slice as found
				keys = append(keys[:i], keys[i+1:]...)
			}
		}
	}
	return attrMap
}

// Is key present within attrs?
func AttrPresent(attrs []html.Attribute, key string) bool {
	for _, attr := range attrs {
		if attr.Key == key {
			return true
		}
	}
	return false
}

// Get hash/fragment id from node.Attrs
func GetId(attrs []html.Attribute) string {
	for _, attr := range attrs {
		if attr.Key == "id" {
			return attr.Val
		} else if attr.Key == "name" {
			return attr.Val
		}
	}
	return ""
}
