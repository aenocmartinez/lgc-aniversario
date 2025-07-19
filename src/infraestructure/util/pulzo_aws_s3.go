package util

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/getsentry/sentry-go"
)

type FileDataToS3 struct {
	Name      string
	Extension string
	Content   []byte
	Bucket    string
}

var projectDirName string = "lgc-aniversario"

func createFileLocal(fileName string, content []byte) (filePath string, err error) {
	filePath = "./" + pathFile(fileName)

	f, err := os.Create(filePath)
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	defer f.Close()

	_, err = f.WriteString(string(content))
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	return filePath, err
}

func deleteFileLocal(filePath string) error {
	return os.Remove(filePath)
}

func UploadFile(fileDataToS3 FileDataToS3) (path string, code int, err error) {

	var endpointPulzoAwsS3 string = os.Getenv("AWS_URL_UPLOAD")

	filePath, err := createFileLocal(fileDataToS3.Name, fileDataToS3.Content)
	if err != nil {
		return "", 500, err
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	fileLocal, err := os.Open(filePath)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	defer fileLocal.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	_, err = io.Copy(part1, fileLocal)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	writer.WriteField("path", fileDataToS3.Bucket)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpointPulzoAwsS3, payload)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	deleteFileLocal(filePath)
	return depureBodyResponse(string(body)), 200, nil
}

func depureBodyResponse(body string) string {
	body = strings.ReplaceAll(body, "urlFileUploaded", "")
	body = strings.ReplaceAll(body, "\"", "")
	body = strings.ReplaceAll(body, ":http", "http")
	body = strings.ReplaceAll(body, "{", "")
	body = strings.ReplaceAll(body, "}", "")
	return body
}

func pathFile(filaname string) string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath) + filaname
}
