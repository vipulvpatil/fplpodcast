package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const AUDIO_URL = "https://audioboom.com/posts/8555343.mp3?modified=1723551853&amp;sid=5001585&amp;source=rss"
const FILE_PATH = "S7 Ep4: FPL Pod: Finalising our Gameweek 1 team"

func DownloadPodcast() error {
	fmt.Println("start downloading podcast")
	out, err := os.Create(fmt.Sprintf("internal/files/audio/%s.mp3", FILE_PATH))
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(AUDIO_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
