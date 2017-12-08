package entities

import (
	"strings"
)

type meetingDAO DataAccessObject

func (dao *meetingDAO) Insert(m *Meeting) error {
	sqlStmt := `
        INSERT INTO meetings(title, host, members, starttime, endtime) VALUES(
            '` + m.Title + `',
            '` + m.Host + `',
            '` + strings.Join(m.Members, "&") + `',
            '` + m.Starttime + `',
            '` + m.Endtime + `',
        )
    `
	result, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

func (dao *meetingDAO) FindAll() ([]Meeting, error) {
	sqlStmt := `SELECT * FROM meetings`

	rows, err := dao.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetingList := make([]Meeting, 0, 0)
	for rows.Next() {
		m := Meeting{}
		memberList := ""
		err := rows.Scan(
			&m.ID, &m.Title, &m.Host, &memberList, &m.Starttime, &m.Endtime)
		m.Members = strings.Split(memberList, "&")
		if err != nil {
			return nil, err
		}
		meetingList = append(meetingList, m)
	}

	return meetingList, nil
}

func (dao *meetingDAO) DeleteMeetingsHostedByUser(username string) error {
	sqlStmt := `DELETE FROM meetings WHERE host = '` + username + `';`
	_, err := dao.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
