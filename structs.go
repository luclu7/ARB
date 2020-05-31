package main

type pkg struct {
	ID     int
	UUID   string `gorm:"type:string"`
	Name   string
	Status int
	Secret string
}

type File struct {
	ID   int
	UUID string `gorm:"type:string"`
	URL  string
}

type requestResponse struct {
	Type int
	Text string
}
