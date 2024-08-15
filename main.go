package main

import (
	"fmt"

	"github.com/vipulvpatil/fplpodcast/internal/downloader"
	"github.com/vipulvpatil/fplpodcast/internal/model"
)

const AUDIO_URL = "https://audioboom.com/posts/8555343.mp3?modified=1723551853&amp;sid=5001585&amp;source=rss"
const FILENAME = "S7 Ep4: FPL Pod: Finalising our Gameweek 1 team"

func main() {

	episode := model.Episode{
		AudioUrl: AUDIO_URL,
		Title:    FILENAME,
	}
	err := downloader.DownloadPodcastEpisode(episode)
	if err != nil {
		fmt.Println(err)
	}
}
