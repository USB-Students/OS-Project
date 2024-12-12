package server

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var (
	directory = "../data"
)

func TestCompareSyncAndParallel(t *testing.T) {
	start := time.Now()
	college, score, err := ProcessFilesParallel(directory)
	if err != nil {
		t.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("parallel elapsed: %d\n", elapsed.Nanoseconds())
	log.Println(college.String(), "score:", score)

	start = time.Now()
	syncCollege, score, err := ProcessFilesSync(directory)
	if err != nil {
		t.Fatal(err)
	}
	elapsed = time.Since(start)
	fmt.Printf("sync elapsed: %d\n", elapsed.Nanoseconds())
	log.Println(syncCollege.String(), "score:", score)
}
