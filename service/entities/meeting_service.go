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
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns all meetings in the database
func (*MeetingServiceProvider) FindAll() ([]Meeting, error) {
	dao := meetingDAO{db}
	meetings, err := dao.FindAll()
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

// DeleteMeetingsHostedByUser deletes all meetings hosted by
// a specific user
func (*MeetingServiceProvider) DeleteMeetingsHostedByUser(username string) error {
	dao := meetingDAO{db}
	err := dao.DeleteMeetingsHostedByUser(username)
	if err != nil {
		return err
	}
	return nil
}
