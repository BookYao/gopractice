/**
 * @Author: BookYao
 * @Description: 练习 7.9： 使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。将这
个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表
格
   注意： 这个demo 在 web上还不能获取到排序数据
 * @File:  trackHtml
 * @Version: 1.0.0
 * @Date: 2020/8/26 16:53
 */

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

var trackTable = template.Must(template.New("Track").Parse(`
<h1> Tracks </h1>
<table>
<tr style='text-align: left'>
    <th οnclick="submitform('Title')">Title
        <form action="" name="Title" method="post">
            <input type="hidden" name="orderby" value="Title"/>
        </form>
    </th>
    <th>Artist
        <form action="" name="Artist" method="post">
            <input type="hidden" name="orderby" value="Artist"/>
        </form>
    </th>
    <th>Album
        <form action="" name="Album" method="post">
            <input type="hidden" name="orderby" value="Album"/>
        </form>
    </th>
    <th οnclick="submitform('Year')">Year
        <form action="" name="Year" method="post">
            <input type="hidden" name="orderby" value="Year"/>
        </form>
    </th>
    <th οnclick="submitform('Length')">Length
        <form action="" name="Length" method="post">
            <input type="hidden" name="orderby" value="Length"/>
        </form>
    </th>
</tr>
{{range .T}}
<tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type Multier struct {
	t []*Track
	primary string
	secondary string
	third string
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func printTrack(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout,  0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}

	tw.Flush()
}

func printTracks(w io.Writer, x sort.Interface) {
	if x, ok := x.(*Multier); ok {
		trackTable.Execute(w, x)
	}
}

func (m *Multier) Len() int {
	return len(m.t)
}

func (m *Multier) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

func (m *Multier) Less(i, j int) bool {
	key := m.primary
	for k := 0; k < 3; k++ {
		switch(key) {
		case "Title":
			if m.t[i].Title != m.t[j].Title {
				return m.t[i].Title < m.t[j].Title
			}
		case "Year":
			if m.t[i].Year != m.t[j].Year {
				return m.t[i].Year < m.t[j].Year
			}
		case "Length":
			if m.t[i].Length != m.t[j].Length {
				return m.t[i].Length < m.t[j].Length
			}
		}

		if k == 0 {
			key = m.secondary
		} else if k == 1 {
			key = m.third
		}
	}

	return false
}

func setPrimary(m *Multier, p string) {
	m.primary, m.secondary, m.third = p, m.primary, m.secondary
}

func SetPrimary(x sort.Interface, p string) {
	if x, ok := x.(*Multier); ok {
		setPrimary(x, p)
	}
}

func NewMultier(t []*Track, p, s, th string) *Multier {
	return &Multier{
		t:t,
		primary:p,
		secondary:s,
		third:th,
	}
}

func main() {
	var track = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	}

	fmt.Println(track, len(track))
	printTrack(track)

	fmt.Println("\n======== Multier Sort =======\n")
	multier := NewMultier(track, "Title", "Year", "Length")
	SetPrimary(multier, "Title")
	sort.Sort(multier)
	printTrack(track)

	// start a simple server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm: %v\n", err)
		}

		fmt.Println("xxxxxx")
		for k, v := range r.Form {
			fmt.Printf("k:%s-v:%s\n", k, v)
			if k == "orderby" {
				SetPrimary(multier, v[0])
			}
		}
		fmt.Println("yyyy")
		sort.Sort(multier)
		printTrack(track)
		fmt.Println("zzz")
		printTracks(w, multier)
	})
	http.ListenAndServe("192.168.50.132:8003", nil)

}