package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client = &http.Client{}

func TestTokens(args *Args) (*Me, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"/me", nil)
	if err != nil {
		log.Fatal(err)
	}

	addQueryParam(req, map[string]string{"client_id": args.clientID, "app_version": APP_VERSION})
	addHeaders(req, args.oauth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var me = new(Me)
	if err = json.NewDecoder(resp.Body).Decode(me); err != nil {
		return nil, err
	}

	return me, nil
}

func GetUserPlaylists(args *Args) (*Playlists, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s/playlists", BASE_URL, args.userID), nil)
	if err != nil {
		return nil, err
	}

	addQueryParam(req, map[string]string{
		"client_id":   args.clientID,
		"app_version": APP_VERSION,
		"show_tracks": "false",
	})
	addHeaders(req, args.oauth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	var playlists = new(Playlists)
	if err = json.NewDecoder(resp.Body).Decode(playlists); err != nil {
		return nil, err
	}

	return playlists, nil
}

func GetPlayListTracks(args *Args, playlistID int) ([]*Track, int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/playlists/%d", BASE_URL, playlistID), nil)
	if err != nil {
		return nil, 0, err
	}

	addQueryParam(req, map[string]string{
		"client_id":   args.clientID,
		"app_version": APP_VERSION,
		"show_tracks": "false",
	})
	addHeaders(req, args.oauth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	var playlist = new(Playlist)
	if err = json.NewDecoder(resp.Body).Decode(playlist); err != nil {
		return nil, 0, err
	}

	return playlist.Tracks, playlist.TrackCount, nil
}

func CreatePlaylist(args *Args, pl any) (*Playlist, error) {
	data := map[string]any{"playlist": pl}
	// payloadBuf := new(bytes.Buffer)
	// if err := json.NewEncoder(payloadBuf).Encode(data); err != nil {
	// 	return nil, err
	// }

	o, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/playlists", BASE_URL), bytes.NewReader(o))
	if err != nil {
		return nil, err
	}
	addQueryParam(req, map[string]string{
		"client_id":   args.clientID,
		"app_version": APP_VERSION,
	})
	addHeaders(req, args.oauth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	var playlist = new(Playlist)
	if err = json.NewDecoder(resp.Body).Decode(playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}

func addQueryParam(req *http.Request, params map[string]string) {
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
}

func addHeaders(req *http.Request, oauth string) {
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", oauth))
	req.Header.Add("accept", "application/json; charset=utf-8")
}
