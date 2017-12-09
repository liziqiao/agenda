package entities

// A Meeting specifies a meeting in the agenda
//
// ID - The id of the meeting
// Title - The title of the meeting
// Host - The host of the meeting
// Members - The users who attend the meeting
// Starttime - The start time of the meeting, in form YYYY/MM/DD/HH:MM
// Endtime - The end time of the meeting, in form YYYY/MM/DD/HH:MM
type Meeting struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Host      string   `json:"host"`
	Members   []string `json:"members"`
	Starttime string   `json:"starttime"`
	Endtime   string   `json:"endtime"`
}

// NewMeeting returns a new instance of a meeting
func NewMeeting(title string, host string, members []string,
	starttime string, endtime string) *Meeting {
	return &Meeting{-1, title, host, members, starttime, endtime}
}
