package helpers

import (
	"fmt"
	"hoxt/data"
	"html"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Function For HTML-Templates on path '/HOXT/templates/*'
var FuncMap = template.FuncMap{
	"upper": strings.ToUpper,
	"formatDate": func(t time.Time) string {
		return t.Format("Monday, Jan 2, 2006 15:04:05")
	},
	"Escape":    html.EscapeString,
	"CutString": TruncateByte,
	"JoinEscape": func(text []string) string {
		var s []string
		for _, v := range text {
			s = append(s, html.EscapeString(v))
		}
		return strings.Join(s, "\n")
	},
	"Rand": rand.Intn,
	"RGB2String": func(r data.RGB) string {
		return fmt.Sprintf("%s,%s,%s", strconv.Itoa(r.R), strconv.Itoa(r.G), strconv.Itoa(r.B))
	},
	"DestroySpaces": DestroySpaces,
	"Split": func(s string) []string {
		return strings.Split(s, "\n")
	},
	"seq": func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i + 1 // Start from 1 instead of 0
		}
		return s
	},
}
