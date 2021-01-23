package util

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Header struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"header"`
	Data interface{} `json:"data"`
}

// ResponseJSON is wrapper function to write HTTP response in JSON format
func ResponseJSON(w http.ResponseWriter, httpStatus int, errMsg string, data interface{}) {
	var res jsonResponse
	res.Header.Status = "success"
	if errMsg != "" {
		res.Header.Status = "error"
		res.Header.Message = errMsg
	}
	res.Data = data
	e, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	w.Write(e)
}

func ResponseImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)

	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
