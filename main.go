package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baaj2109/gin-cloud-storage/router"
)

func init() {

	// var serviceAccountFilePath string = "./google_drive_account.json"

	// sharedService, err := drive.NewService(context.Background(), option.WithCredentialsFile(serviceAccountFilePath))
	// if err != nil {
	// 	panic(err)
	// }

	// r, err := sharedService.Files.List().PageSize(10).Fields("nextPageToken, files(name)").Do()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Files:")
	// if len(r.Files) == 0 {
	// 	fmt.Println("No files found.")
	// } else {
	// 	for _, i := range r.Files {
	// 		fmt.Printf("%s \n", i.Name)
	// 	}
	// }

	// fmt.Print("done")

}

func main() {

	r := router.InitRouters()

	r.LoadHTMLGlob("view/*")
	r.Static("/static", "./static")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
