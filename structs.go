package main

type pkg struct {
	ID     int
	UUID   string `gorm:"type:string"`
	Name   string
	Status int
	Secret string `json:"-"`
}

// File represents a file uploaded to S3/Minio
type File struct {
	ID   int
	UUID string `gorm:"type:string"`
	URL  string
}

type requestResponse struct {
	Type int
	UUID string
	Text string
}

type buildRequest struct {
	PackageName string
}

type addUrlRequest struct {
	UUID   string
	URL    string
	Secret string
}
