package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/vipulvpatil/fplpodcast/internal/downloader"
	"github.com/vipulvpatil/fplpodcast/internal/files"
	"github.com/vipulvpatil/fplpodcast/internal/model"
	"github.com/vipulvpatil/fplpodcast/internal/transcriber"
	"github.com/vipulvpatil/fplpodcast/services/openai"
)

const AUDIO_URL = "https://audioboom.com/posts/8555343.mp3?modified=1723551853&amp;sid=5001585&amp;source=rss"
const FILENAME = "S7 Ep4: FPL Pod: Finalising our Gameweek 1 team"

func main() {
	viper.AutomaticEnv()
	episode := model.Episode{
		AudioUrl: AUDIO_URL,
		Title:    FILENAME,
	}
	err := downloader.DownloadPodcastEpisode(episode)
	if err != nil {
		fmt.Println(err)
		return
	}

	file := transcriber.GetTranscripionFilepath(episode.Filename())
	transcriptionText, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	analysisFilepath := files.PathForAnalysisFile(episode.Filename())
	_, err = os.Stat(analysisFilepath)
	if err == nil {
		fmt.Println("Analysis file exists: ", analysisFilepath)
		return
	}

	analysis, err := openai.GetAnalysis(string(transcriptionText))
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := os.Create(analysisFilepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = out.WriteString(analysis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
