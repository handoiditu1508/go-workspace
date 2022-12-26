package exercise2

import (
	"fmt"
	"strconv"
	"strings"
)

var typeDict = map[string]int{"music": 0, "images": 0, "movies": 0, "": 0}
var extDict = map[string]string{"mp3": "music", "aac": "music", "flac": "music", "jpg": "images", "bmp": "images", "gif": "images", "mp4": "movies", "avi": "movies", "mkv": "movies"}

func InvokeExercise5(input string) string {
	fileList := strings.Split(input, "\n")

	for _, value := range fileList {
		dotIndex := strings.LastIndex(value, ".")
		spaceIndex := strings.LastIndex(value, " ")
		fileExt := value[dotIndex+1 : spaceIndex]
		fileType := extDict[fileExt]
		fileSize, _ := strconv.ParseInt(value[spaceIndex+1:len(value)-1], 0, 32)
		typeDict[fileType] += int(fileSize)
	}

	output := fmt.Sprintf("%s %db\n", "music", typeDict["music"]) + fmt.Sprintf("%s %db\n", "images", typeDict["images"]) + fmt.Sprintf("%s %db\n", "movies", typeDict["movies"]) + fmt.Sprintf("%s %db", "other", typeDict[""])
	return output
}
