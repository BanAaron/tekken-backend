package database

func GetCharacters() (characters []Character, err error) {
	rows, err := Db.Query("select id, short_name, long_name, fighting_style, nationality, height, weight, gender from characters order by short_name")
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id            uint8
			shortName     string
			longName      string
			fightingStyle string
			nationality   string
			height        uint16
			weight        uint16
			gender        string
		)

		err = rows.Scan(&id, &shortName, &longName, &fightingStyle, &nationality, &height, &weight, &gender)
		if err != nil {
			return
		}
		char := Character{
			Id:            id,
			ShortName:     shortName,
			LongName:      longName,
			FightingStyle: fightingStyle,
			Nationality:   nationality,
			Height:        height,
			Weight:        weight,
			Gender:        gender,
		}
		characters = append(characters, char)
	}
	err = rows.Close()
	return characters, err
}
