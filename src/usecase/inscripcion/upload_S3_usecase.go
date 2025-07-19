package usecase

import (
	"fmt"
	"io"
	"lgc/src/infraestructure/util"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type UploadS3UseCase struct{}

func NewUploadS3UseCase() *UploadS3UseCase {
	return &UploadS3UseCase{}
}

func (uc *UploadS3UseCase) Execute(fileHeader *multipart.FileHeader) (string, error) {
	const maxSize = 5 * 1024 * 1024

	if fileHeader.Size > int64(maxSize) {
		return "", fmt.Errorf("el archivo excede el tamaño máximo permitido de 5 MB")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".pdf" && ext != ".heic" {
		return "", fmt.Errorf("tipo de archivo no permitido")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error al leer el archivo: %v", err)
	}

	nombreArchivo := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	fileDataToS3 := util.FileDataToS3{
		Name:      nombreArchivo,
		Extension: ext,
		Content:   content,
		Bucket:    "pulzo-dev/lgc-aniversario",
	}

	path, _, err := util.UploadFile(fileDataToS3)
	if err != nil {
		return "", err
	}

	fmt.Println(path)
	return path, nil
}
