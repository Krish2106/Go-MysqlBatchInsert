func (c *MysqlClient) batchUpdate(data []foo) error {

	sqlStr := "INSERT INTO <Table_Name>(<CN1>,<CN2>,<CN3>...) VALUES"

	vals := make([]interface{}, 0, len(data))

	for _, row := range data {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, row.CN1, row.CN2, row.CN3)
	}

	// Removes the last comma
	sqlStr = sqlStr[0:len(sqlStr)-1]

	stmt, err := c.session.Prepare(sqlStr)
	if err != nil {
		return wraperrors.Wrap(err, "SQL:Prepared statement failed")
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		return wraperrors.Wrap(err, "SQL: Execute statement failed")
	}

	return nil
}
