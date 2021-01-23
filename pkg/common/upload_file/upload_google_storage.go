package upload_file

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/storage"

	//"github.com/cpartogi/izyai/internal/menu/helper"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"

	"google.golang.org/api/option"
)

func UploadFile(ctx context.Context, fileHeader *bytes.Reader, fileName string) error {
	// Prevent log from printing out time information
	log.SetFlags(0)

	var projectID, bucket string
	var public bool

	bucket = mustGetEnv("GOOGLE_STORAGE_IMAGE_BUCKET", bucket)
	projectID = mustGetEnv("GCP_PROJECT_ID", projectID)
	public = true

	_, _, err := upload(ctx, fileHeader, projectID, bucket, fileName, public)
	if err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			log.Println("Please create the bucket first e.g. with `gsutil mb`")
		default:
			log.Println(err)
		}
		return err
	}

	return nil
}
func mustGetEnv(envKey, defaultValue string) string {
	val := os.Getenv(envKey)
	if val == "" {
		val = defaultValue
	}
	if val == "" {
		log.Fatalf("%q should be set", envKey)
	}
	return val
}

func objectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", objAttrs.Bucket, objAttrs.Name)
}

func upload(ctx context.Context, r *bytes.Reader, projectID, bucket, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	var credFile = os.Getenv("STORAGE_CREDENTIAL")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	bh := client.Bucket(bucket)
	// Next check if the bucket exists
	if _, err = bh.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}

func CompressImage(img image.Image, imageFormat string) ([]byte, error) {
	switch imageFormat {
	case "png":
		return compressPNG(img)
	case "jpeg":
		return compressJPEG(img)
	default:
		return nil, errors.New("unsupported image file format")
	}
}

func compressPNG(img image.Image) ([]byte, error) {
	var buff bytes.Buffer

	w := bufio.NewWriter(&buff)
	pngEncoder := png.Encoder{CompressionLevel: png.BestCompression}
	if err := pngEncoder.Encode(w, img); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func compressJPEG(img image.Image) ([]byte, error) {
	var buff bytes.Buffer

	w := bufio.NewWriter(&buff)
	if err := jpeg.Encode(w, img, &jpeg.Options{Quality: 60}); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
