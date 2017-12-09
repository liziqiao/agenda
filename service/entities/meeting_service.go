package entities

// A MeetingServiceProvider provides a series of
// operations to Entity Meeting
type MeetingServiceProvider struct{}

// MeetingService is an instance of MeetingServiceProvider
var MeetingService = MeetingServiceProvider{}

// Insert inserts a new meeting to the database
func (*MeetingServiceProvider) Insert(m *Meeting) error {
	dao := meetingDAO{db}
	err := dao.Insert(m)
	panicIfErr(err)
	return nil
}

// FindAll returns all meetings in the database
func (*MeetingServiceProvider) FindAll() ([]Meeting, error) {
	dao := meetingDAO{db}
	meetings, err := dao.FindAll()
	panicIfErr(err)
	return meetings, nil
}

// FindBy returns the first meeting whose value of col is val
func (*MeetingServiceProvider) FindBy(col string, val string) (Meeting, error) {
	dao := meetingDAO{db}
	meeting, err := dao.FindBy(col, val)
	panicIfErr(err)
	return meeting, nil
}

// DeleteMeetingsHostedByUser deletes all meetings hosted by
// a specific user
func (*MeetingServiceProvider) DeleteMeetingsHostedByUser(username string) error {
	dao := meetingDAO{db}
	err := dao.DeleteMeetingsHostedByUser(username)
	panicIfErr(err)
	return nil
}
