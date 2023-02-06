package entities

type Video struct {
	Url 			string 	`json:"url"`
	Title 			string 	`json:"title"`
	Thumbnail 		string 	`json:"thumbnail"`
	Uploader 		string 	`json:"uploaderName"`
	Duration 		uint64	`json:"duration"`
	Uploaded		uint64	`json:"uploaded"`
	Views 			uint64 	`json:"views"`
	// manually added
	UploadDate		string 	`json:"uploadDate"`
	ViewsString 	string 	`json:"viewsString"`
	DurationString 	string 	`json:"durationString"`
}