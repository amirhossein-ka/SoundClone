package main

import (
	"flag"
	"fmt"
)

type Args struct {
	playlistName string
	userID       string
	clientID     string
	oauth        string
}

func (a *Args) Parse() {
	flag.StringVar(&a.playlistName, "name", "SoundClone", "Name of the cloned playlist")
	flag.StringVar(&a.userID, "uid", "", "ID of user that you want to clone their playlist")
	flag.StringVar(&a.clientID, "cid", "", "Your ClientID from soundcloud web")
	flag.StringVar(&a.oauth, "oauth", "", "Your ouath key")
	flag.Parse()
}


func (a *Args) Validate() error {
  if a.userID == "" {
    return fmt.Errorf("userid cannot be empty")
  }

  if a.clientID == "" {
    return fmt.Errorf("client id cannot be empty")
  }

  if a.oauth == "" {
    return fmt.Errorf("oauth key cannot be empty")
  }

  return nil
}
