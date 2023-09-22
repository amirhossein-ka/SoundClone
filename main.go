package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	args   = &Args{}
	reader = bufio.NewReader(os.Stdin)

	// TODO use colors...
	// info = color.New(color.FgCyan).PrintfFunc()
	// warning = color.New(color.FgYellow).PrintfFunc()
	// errors = color.New(color.FgRed).PrintfFunc()
)

func init() {
	args.Parse()
	args.Validate()
}

func main() {
	fmt.Println("Welcome To SoundClone!\nPlease wait while i'm checking your tokens...")
	me, err := TestTokens(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("i'm using account named: %s, is it correct ? (y/n): ", me.FullName)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	in = strings.TrimRight(in, "\n")

	switch in {
	case "y":
	case "n":
		fmt.Printf("Please check your oauth token again\nExiting...")
		return
	default:
		fmt.Printf("unknown option\nExiting...")
		return
	}

	fmt.Printf("fetching playlists of user %s...\n", args.userID)
	playlists, err := GetUserPlaylists(args)
	if err != nil {
		log.Fatal(err)
	}

	if len(playlists.Collection) == 0 {
		fmt.Printf("user has no playlists\nExiting.")
		return
	}

	fmt.Printf("Playlists of user %s:\n", playlists.Collection[0].User.Username)
	for i, pl := range playlists.Collection {
		fmt.Printf("%d. Name: %s,Link: %s,TrackCount:%d\n", i, pl.Title, pl.PermalinkURL, pl.TrackCount)
	}

	fmt.Printf("Please enter the playlist id you want to clone: ")
	pid, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	pid = strings.TrimRight(pid, "\n")

	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatal("invalid id")
	}

	fmt.Println("Fetching playlist tracks...")
	tracks, count, err := GetPlayListTracks(args, playlists.Collection[id].ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fetched %d Tracks from playlist\nCloning to %s...\n", count, args.playlistName)

	tid := make([]int, count)
	for i, t := range tracks {
		tid[i] = t.ID
	}

	pl := map[string]any{
		"title":       args.playlistName,
		"description": "soundclone",
		"sharing":     "public",
		"tracks":      tid,
	}

	createdPlaylist, err := CreatePlaylist(args, pl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done !\nLink to your cloned playlist: %s\nHave Fun !", createdPlaylist.PermalinkURL)
}
