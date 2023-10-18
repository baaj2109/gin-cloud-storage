package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var serviceAccountFilePath string = "./google_drive_account.json"

type Drive struct {
	service *drive.Service
}

// connect to drive server
func NewDrive() (*Drive, error) {
	s, err := drive.NewService(
		context.Background(),
		option.WithCredentialsFile(serviceAccountFilePath),
	)
	if err != nil {
		return nil, err
	}
	return &Drive{s}, nil
	
}

func (d *Drive) ShowAllFiles() {
	r, err := d.service.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		panic(err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for index, value := range r.Files {
			fmt.Printf("%d: %s (%s)\n", index, value.Name, value.Id)
		}
	}
}

func (d *Drive) CreateFolder(folderName string, parentId string) (string, error) {
	parents := []string{}
	if parentId != "" {
		parents = []string{parentId}
	}
	f := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  parents,
	}
	createdFolder, err := d.service.Files.Create(f).Do()
	if err != nil {
		log.Print("failed to create folder: ", err)
		return "", err
	}
	return createdFolder.Id, nil
}

func (d *Drive) Delete(fileId string) error {
	err := d.service.Files.Delete(fileId).Do()
	if err != nil {
		log.Printf("failed to delete file: %v", err)
		return err
	}
	return nil
}

// upload file, return file link, file id, error
func (d *Drive) Upload(bytes []byte, fileName, fileType string, permission *drive.Permission) (string, string, error) {
	file := strings.NewReader(string(bytes))
	fileSize := len(bytes)

	f := &drive.File{
		Name:     fileName,
		MimeType: fileType,
	}

	var err error
	var res *drive.File

	if fileSize > 5*1024*1024 {
		res, err = d.service.Files.Create(f).Media(file).
			SupportsAllDrives(true).
			ProgressUpdater(func(now, size int64) {
				fmt.Printf("%d, %d\r", now, size)
			}).
			Do()
	} else {
		res, err = d.service.Files.Create(f).Media(file).Do()
	}
	if err != nil {
		log.Fatalf("Unable to upload media file: %v", err)
		return "", "", err
	}
	if permission != nil {
		_, err = d.service.Permissions.Create(res.Id, permission).Do()
		if err != nil {
			log.Fatalf("Failed to create permission for the file: %v", err)
			return "", "", err
		}
	}
	fileLink := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", res.Id)
	return fileLink, res.Id, nil
}

func (d *Drive) Download(fileId, localPath, outputFileName string) (err error) {

	metaData, err := d.service.Files.Get(fileId).
		Fields("name,mimeType,size,id,md5Checksum").
		SupportsAllDrives(true).Do()
	if err != nil {
		log.Fatalf("failed to get metadata with file id: %s , %v", fileId, err)
		return err
	}
	log.Printf("Name: %s, MimeType: %s\n", metaData.Name, metaData.MimeType)

	if outputFileName == "" {
		outputFileName = metaData.Name
	}

	outputPath := path.Join(localPath, outputFileName)
	writer, err := os.Create(outputPath)

	if err != nil {
		log.Fatal("failed to create file: ", err)
		return err
	}
	defer writer.Close()

	go func() {
		req, err := d.service.Files.Get(fileId).SupportsTeamDrives(true).Download()

		if err != nil {
			log.Fatal("failed to download file: ", err)
			return
		}
		io.Copy(writer, req.Body)
		defer req.Body.Close()
	}()
	return nil
}
