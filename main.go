package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

var (
	modelsDir = filepath.Join(dataDir, "models")
	imagesDir = filepath.Join(dataDir, "images")
)

func main() {
	// Initialisasi
	fmt.Println("Facial Recognition System v0.01")

	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")

	// Check Count Faces On File avengers-02.jpeg
	avengersImage := filepath.Join(imagesDir, "avengers-02.jpeg")

	faces, err := rec.RecognizeFile(avengersImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	fmt.Println("Number of Faces in Image: ", len(faces))

	// Create Slice Samples and Avengers
	var samples []face.Descriptor
	var avengers []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own category.
		avengers = append(avengers, int32(i))
	}

	// Name person on file avengers-02.jpeg
	labels := []string{
		"Dr Strange",
		"Tony Stark",
		"Bruce Banner",
		"Wong",
	}

	fmt.Println("Name of person: ", labels)

	// Pass samples to the recognizer.
	rec.SetSamples(samples, avengers)

	// Find Tony Stark
	testTonyStark := filepath.Join(imagesDir, "tony-stark.jpg")
	tonyStark, err := rec.RecognizeSingleFile(testTonyStark)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if tonyStark == nil {
		log.Fatalf("Not a single face on the image")
	}
	avengerID := rec.Classify(tonyStark.Descriptor)
	if avengerID < 0 {
		log.Fatalf("Can't classify")
	}

	fmt.Println(avengerID)
	fmt.Println(labels[avengerID])
}
