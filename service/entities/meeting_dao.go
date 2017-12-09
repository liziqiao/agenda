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
            '` + m.Endtime + `'
        )
    `
	result, err := db.Exec(sqlStmt)
	panicIfErr(err)
	id, err := result.LastInsertId()
	panicIfErr(err)
	m.ID = int(id)
	return nil
}

func (dao *meetingDAO) FindAll() ([]Meeting, error) {
	sqlStmt := `SELECT * FROM meetings`

	rows, err := dao.Query(sqlStmt)
	defer rows.Close()
	panicIfErr(err)

	meetingList := make([]Meeting, 0, 0)
	for rows.Next() {
		m := Meeting{}
		memberList := ""
		err := rows.Scan(
			&m.ID, &m.Title, &m.Host, &memberList, &m.Starttime, &m.Endtime)
		m.Members = strings.Split(memberList, "&")
		panicIfErr(err)
		meetingList = append(meetingList, m)
	}

	return meetingList, nil
}

func (dao *meetingDAO) FindBy(col string, val string) (Meeting, error) {
	sqlStmt := `SELECT * FROM meetings WHERE ` + col + ` = '` + val + `';`

	rows, err := dao.Query(sqlStmt)
	defer rows.Close()
	panicIfErr(err)

	m := Meeting{}
	if rows.Next() {
		memberList := ""
		err := rows.Scan(
			&m.ID, &m.Title, &m.Host, &memberList, &m.Starttime, &m.Endtime)
		panicIfErr(err)
		m.Members = strings.Split(memberList, "&")
	}

	return m, nil
}

func (dao *meetingDAO) DeleteMeetingsHostedByUser(username string) error {
	sqlStmt := `DELETE FROM meetings WHERE host = '` + username + `';`
	_, err := dao.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
