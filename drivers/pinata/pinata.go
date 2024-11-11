package pinata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-ads-management/utils"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type PinataFileUploadResponse struct {
	Data Data `json:"data"`
}

type PinataSignedURLResponse struct {
	Data string `json:"data"`
}

type Data struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Cid           string `json:"cid"`
	CreatedAt     string `json:"created_at"`
	Size          int64  `json:"size"`
	NumberOfFiles int64  `json:"number_of_files"`
	MIMEType      string `json:"mime_type"`
	UserID        string `json:"user_id"`
	GroupID       string `json:"group_id"`
	IsDuplicate   bool   `json:"is_duplicate"`
}

type PinataRequest struct {
	URL     string `json:"url"`
	Method  string `json:"method"`
	Expires int64  `json:"expires"`
	Date    int64  `json:"date"`
}

type UploadResponse struct {
	FileID    string
	FileCID   string
	SignedURL string
}

func UploadFile(file *multipart.FileHeader) (UploadResponse, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return UploadResponse{}, err
	}
	defer src.Close()

	// Prepare the multipart form for the Pinata API
	file.Filename = utils.GenerateFilename(file)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return UploadResponse{}, err
	}

	// Copy the file content to the form
	_, err = io.Copy(part, src)
	if err != nil {
		return UploadResponse{}, err
	}

	// Close the writer to finalize the form
	err = writer.Close()
	if err != nil {
		return UploadResponse{}, err
	}

	// Make the request to Pinata API
	url := "https://uploads.pinata.cloud/v3/files"
	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return UploadResponse{}, err
	}

	// Add necessary headers
	auth := fmt.Sprintf("Bearer %s", utils.GetConfig("PINATA_TOKEN"))

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return UploadResponse{}, err
	}
	defer res.Body.Close()

	// Read the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return UploadResponse{}, err
	}

	var uploadResponse PinataFileUploadResponse

	if err := json.Unmarshal(body, &uploadResponse); err != nil {
		return UploadResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return UploadResponse{}, errors.New("file upload failed")
	}

	signedURL, err := getSignedURL(uploadResponse.Data.Cid)

	if err != nil {
		return UploadResponse{}, err
	}

	result := UploadResponse{
		FileID:    uploadResponse.Data.ID,
		FileCID:   uploadResponse.Data.Cid,
		SignedURL: signedURL,
	}

	return result, nil
}

func getSignedURL(cid string) (string, error) {
	url := "https://api.pinata.cloud/v3/files/sign"

	gateway := utils.GetConfig("PINATA_GATEWAY")

	exp, err := strconv.Atoi(utils.GetConfig("PINATA_LINK_EXPIRATION"))

	if err != nil {
		return "", err
	}

	var expire int64 = int64(exp) // in seconds

	fileUrl := fmt.Sprintf(
		"https://%s/files/%s",
		gateway,
		cid,
	)

	payload, err := json.Marshal(PinataRequest{
		URL:     fileUrl,
		Date:    time.Now().Unix(),
		Expires: expire,
		Method:  http.MethodGet,
	})

	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	auth := fmt.Sprintf("Bearer %s", utils.GetConfig("PINATA_TOKEN"))

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var signedURLResponse PinataSignedURLResponse

	if err := json.Unmarshal(body, &signedURLResponse); err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("get signed URL failed")
	}

	return signedURLResponse.Data, nil
}

func DeleteFile(fileID string) error {
	url := fmt.Sprintf("https://api.pinata.cloud/v3/files/%s", fileID)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	auth := fmt.Sprintf("Bearer %s", utils.GetConfig("PINATA_TOKEN"))

	req.Header.Add("Authorization", auth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return nil
}
