package xmldom

import (
	"testing"
)

func TestSvgParse(t *testing.T) {
	doc, err := ParseFile("test.svg")
	if err != nil {
		t.Fatalf(err.Error())
	}
	root := doc.Root

	imagesNodes := root.FindByName("image")
	if len(imagesNodes) < 4 {
		t.Fatalf("No images")
	}
}

func TestAttrNamespace(t *testing.T) {
	root := Must(ParseFile("test.svg")).Root
	uses := root.FindByName("use")

	var contains bool
	for _, a := range root.Attributes {
		if a.Name == "xmlns:xlink" {
			contains = true
		}
	}

	if !contains {
		t.Fatalf("Expect root to contain xmlns:xlink attribute")
	}

	for _, u := range uses {
		for _, a := range u.Attributes {
			if a.Name == "href" {
				t.Fatalf("Expect use tag to contain xlink:href attribute but got href")
			}
		}
	}
}
