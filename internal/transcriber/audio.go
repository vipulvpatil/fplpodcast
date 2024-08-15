package transcriber

import (
	"fmt"
	"os"

	"github.com/vipulvpatil/fplpodcast/internal/files"
)

func GetTranscripionFilepath(filename string) string {
	// For now, manually transcribing file.
	// Only checking it exists here.
	filepath := files.PathForTranscriptionFile(filename)
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("Transcription file does not exists: ", err)
		return ""
	}
	return filepath
}
