package API

type Pkg struct {
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

type RequestResponse struct {
	Type int
	UUID string
	Text string
}

type BuildRequest struct {
	PackageName string
}

type AddURLRequest struct {
	UUID   string
	URL    string
	Secret string
}
