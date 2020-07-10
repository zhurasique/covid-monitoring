package collect

import (
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type CollectLinks struct{

}

// All takes a reader object (like the one returned from http.Get())
// It returns a slice of strings representing the "href" attributes from
// anchor links found in the provided html.
// It does not close the reader passed to it.
func (c *CollectLinks) All(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := c.trimHash(attr.Val)
					col = append(col, tl)
					c.resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func (c *CollectLinks) trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func (c *CollectLinks) check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func (c *CollectLinks) resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if c.check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}