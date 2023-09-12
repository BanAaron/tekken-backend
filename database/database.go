package database

import (
	"database/sql"
	"fmt"

	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

type Move struct {
	MoveName            string
	InputCommand        string
	StartupFrames       int
	ActiveFrames        int
	RecoveryFrames      int
	Damage              int
	Height              string
	Launcher            bool
	Throw               bool
	KnockDown           bool
	CounterHitLauncher  bool
	CounterHitKnockDown bool
	HeatEngage          bool
	HeatSmash           bool
	RageArt             bool
	RageDrive           bool
	Unblockable         bool
	GuardBreak          bool
}

type Character struct {
	Id            int
	ShortName     string
	LongName      string
	FightingStyle string
	Nationality   string
	Height        int
	Weight        int
	Gender        string
}

func (c Character) String() string {
	var gender string
	if c.Gender == "m" {
		gender = "Male"
	} else {
		gender = "Female"
	}
	result := fmt.Sprintf(
		"%d, %s (%s), %s, %s, %dcm, %dkg, %s",
		c.Id,
		c.ShortName,
		c.LongName,
		c.FightingStyle,
		c.Nationality,
		c.Height,
		c.Weight,
		gender,
	)
	return result
}

type ConnectionString struct {
	username string
	password string
	host     string
	port     int
	dbname   string
}

func (cs ConnectionString) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		"[PASSWORD]",
		cs.host,
		cs.port,
		"[DBNAME]",
	)
}

func (cs ConnectionString) Get() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.dbname)
}

func NewConnectionString(
	username string,
	password string,
	host string,
	port int,
	dbname string,
) ConnectionString {
	return ConnectionString{
		username: username,
		password: password,
		host:     host,
		port:     port,
		dbname:   dbname,
	}
}

func GetCharacters(db *sql.DB) (characters []Character) {
	res, err := db.Query("select id, short_name, long_name, fighting_style, nationality, height, weight, gender from characters")
	util.CheckError(err)

	defer func(res *sql.Rows) {
		err := res.Close()
		util.CheckError(err)
	}(res)

	for res.Next() {
		var (
			id            int
			shortName     string
			longName      string
			fightingStyle string
			nationality   string
			height        int
			weight        int
			gender        string
		)

		err := res.Scan(
			&id,
			&shortName,
			&longName,
			&fightingStyle,
			&nationality,
			&height,
			&weight,
			&gender,
		)
		util.CheckError(err)

		characters = append(characters, Character{
			Id:            id,
			ShortName:     shortName,
			LongName:      longName,
			FightingStyle: fightingStyle,
			Nationality:   nationality,
			Height:        height,
			Weight:        weight,
			Gender:        gender,
		})

	}
	return characters
}
