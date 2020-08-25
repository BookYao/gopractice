/**
 * @Author: BookYao
 * @Description:
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
}

  