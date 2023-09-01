package main

import (
	"net/http"
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

	// shutdownError := make(chan error)
	// go func() {
	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 	signal := <-quit
	// 	fmt.Println("shutting down server", map[string]string{
	// 		"signal": signal.String(),
	// 	})
	// 	// app.logger.Printf("shutting down server %s", s)
	// 	// os.Exit(0)
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()

	// 	shutdownError <- s.Shutdown(ctx)
	// }()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

	// err := <-shutdownError
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("starting %s server on 8080")
}
