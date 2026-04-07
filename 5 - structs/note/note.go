package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"` //  The struct tags are not mandatory, but it's good to do
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid metadata")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Title: %v\nContent: %v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"

	// json.Marshal only converts exposed fields from the struct
	jsonText, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonText, 0644)
}
