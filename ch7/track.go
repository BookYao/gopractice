/**
 * @Author: BookYao
 * @Description:
 update1: 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次
点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。定义一个sort.Interface的
实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。
 * @File:  track
 * @Version: 1.0.0
 * @Date: 2020/8/25 23:06
 */

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

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

type bySortType []*Track
func (x bySortType) Len() int {
	return len(x)
}

func (x bySortType) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x bySortType) Swap(i, j int) {
	x[i], x[j] =  x[j], x[i]
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
	var track = []*Track {
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	fmt.Println(track, len(track))
	printTrack(track)

	fmt.Println("\n======== Sort =======\n")
	sort.Sort(bySortType(track))
	printTrack(track)

	fmt.Println("\n======== Reverse Sort =======\n")
	sort.Sort(sort.Reverse(bySortType(track)))
	printTrack(track)

	fmt.Println("\n======== Multier Sort =======\n")
	multier := NewMultier(track, "Title", "Year", "Length")
	SetPrimary(multier, "Title")
	sort.Sort(multier)
	printTrack(track)

	fmt.Println("\n======== Set Primary Year. Multier Sort =======\n")
	setPrimary(multier, "Year")
	sort.Sort(multier)
	printTrack(track)
}

  