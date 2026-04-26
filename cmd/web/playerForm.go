package main

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
)

var ErrTooLarge = errors.New("file too large")
var ErrNotImage = errors.New("file is not an image")

func generateSecureFilename(ext string) string {
	name := uuid.NewString()
	filename := fmt.Sprintf("%s.%s", name, ext)
	return filename
}

// Handler to regsiterPlayerForm moved from main handlers file
// due to the extra pieces
func (app *App) regsiterPlayerForm(w http.ResponseWriter, r *http.Request) {
	var imagePath string
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		app.logger.Error("Failed parsingMultipartFile", "error", err.Error())
		return
	}

	// Get the file from the form data
	file, header, err := r.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			http.Error(w, "Server encountered an error", http.StatusInternalServerError)
			app.logger.Error("Image file upload encountered problem", "Error", err.Error())
			return
		}
		imagePath = ""
		// no file uploaded which is cool
		goto PARSEFORM
	}
	defer file.Close()
	imagePath, err = handleImageFile(file, header)
	if err != nil {
		app.logger.Error(err.Error())
		if errors.Is(err, ErrTooLarge) || errors.Is(err, ErrNotImage) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Server Encountered an Error", http.StatusInternalServerError)
		return
	}

PARSEFORM:
	name := r.PostForm.Get("Name")
	offense, err := strconv.Atoi(r.PostForm.Get("Offense"))
	if err != nil {
		app.logger.Error("Error converting player offense", "Error", err.Error())
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
	defense, err := strconv.Atoi(r.PostForm.Get("Defense"))
	if err != nil {
		app.logger.Error("Error converting player defense", "Error", err.Error())
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
	speed, err := strconv.Atoi(r.PostForm.Get("Speed"))
	if err != nil {
		app.logger.Error("Error converting player speed", "Error", err.Error())
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
	endurance, err := strconv.Atoi(r.PostForm.Get("Endurance"))
	if err != nil {
		app.logger.Error("converting player endurance", "Error", err.Error())
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
	style := r.PostForm.Get("Style")
	app.logger.Info("New player", "Name", name, "Offense", offense, "Defense", defense, "Endurance", endurance, "Style", style)
	// Create new player object
	p := &models.Player{
		Name:      name,
		Image:     imagePath,
		Offense:   offense,
		Defense:   defense,
		Speed:     speed,
		Endurance: endurance,
		Style:     style,
	}
	models.NewPlayer(p)
	// Redirect to player cards
	http.Redirect(w, r, "/players", http.StatusSeeOther)
}

func handleImageFile(file multipart.File, header *multipart.FileHeader) (string, error) {

	// Validate file size
	if header.Size > 5*1024*1024 {
		return "", ErrTooLarge
	}

	// Validate file type
	if !strings.HasPrefix(header.Header.Get("Content-Type"), "image/") {
		return "", ErrNotImage
	}

	// Read first 512 bytes for content type detection
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Detect the content type
	contentType := http.DetectContentType(buffer)
	ext := strings.TrimPrefix(contentType, "image/")

	// Important: Reset the file pointer to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}
	sFileName := generateSecureFilename(ext)
	// Create a new filename
	filename := filepath.Join("./ui/static/uploads/", sFileName)

	// Create the file
	dst, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}
	// parse out the .ui/static
	imageURL := fmt.Sprintf("http://localhost:3000/static/uploads/%s", sFileName)
	fmt.Println(imageURL)
	return imageURL, nil
}
