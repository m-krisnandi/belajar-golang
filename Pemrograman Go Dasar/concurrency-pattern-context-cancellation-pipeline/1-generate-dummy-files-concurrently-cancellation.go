package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	// add context package for cancellation
	"context"
)

const totalFile = 3000
const contentLength = 5000

// add timeout for cancellation
const timeoutDuration = 3 * time.Second

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.61-pipeline-cancellation-context")

type FileInfo struct {
    Index       int
    FileName    string
    WorkerIndex int
    Err         error
}

func main() {
    log.Println("start")
    start := time.Now()

    // change the generateFiles() to use context
	// generateFiles()

	// add context with timeout for cancellation (tanpa paksa)
	// ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	// defer cancel()
	// generateFilesWithContext(ctx)

	// cancel the context secara paksa(tanpa timeout)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(timeoutDuration, cancel)
	generateFilesWithContext(ctx)

    duration := time.Since(start)
    log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
    randomizer := rand.New(rand.NewSource(time.Now().Unix()))
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, length)
    for i := range b {
        s := randomizer.Intn(len(letters))
        b[i] = letters[s]
    }

    return string(b)
}

func generateFiles() {
	// change this line to use context

    // os.RemoveAll(tempPath)
    // os.MkdirAll(tempPath, os.ModePerm)

    // // pipeline 1: job distribution
    // chanFileIndex := generateFileIndexes()

    // // pipeline 2: the main logic (creating files)
    // createFilesWorker := 100
    // chanFileResult := createFiles(chanFileIndex, createFilesWorker)

    // // track and print output
    // counterTotal := 0
    // counterSuccess := 0
    // for fileResult := range chanFileResult {
    //     if fileResult.Err != nil {
    //         log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
    //     } else {
    //         counterSuccess++
    //     }
    //     counterTotal++
    // }

    // log.Printf("%d/%d of total files created", counterSuccess, counterTotal)

	// add generateFileWithContext() here
	generateFilesWithContext(context.Background())
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
    os.MkdirAll(tempPath, os.ModePerm)

	// add done channel for cancellation
	done := make(chan int)

	go func() {
		// pipeline 1: job distribution
		chanFileIndex := generateFileIndexes(ctx)

		// pipeline 2: the main logic (creating files)
		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)
	
		// track and print output
		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}

		// notify that the process is done
		done <- counterSuccess
	}()

	// add select for cancellation
	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}
}

func generateFileIndexes(ctx context.Context) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        for i := 0; i < totalFile; i++ {
			// change this line to use context
            // chanOut <- FileInfo{
            //     Index:    i,
            //     FileName: fmt.Sprintf("file-%d.txt", i),
            // }

			// add select for cancellation
			select {
				case <-ctx.Done():
					break
				default:
					chanOut <- FileInfo{
						Index:    i,
						FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}

        }
        close(chanOut)
    }()

    return chanOut
}

func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    wg := new(sync.WaitGroup)
    wg.Add(numberOfWorkers)

    go func() {
        for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
            go func(workerIndex int) {
                for job := range chanIn {
                    // add select for cancellation
					select {
					case <-ctx.Done():
						break
					default:
						filePath := filepath.Join(tempPath, job.FileName)
                    	content := randomString(contentLength)
                    	err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

                    	log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

                    	chanOut <- FileInfo{
                        	FileName:    job.FileName,
                        	WorkerIndex: workerIndex,
                        	Err:         err,
                		}

					}
                }

                wg.Done()
            }(workerIndex)
        }
    }()

    go func() {
        wg.Wait()
        close(chanOut)
    }()

    return chanOut
}