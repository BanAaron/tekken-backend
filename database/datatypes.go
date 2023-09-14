package database

import (
	"fmt"
)

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

type Move struct {
	Id                  int
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

type CharacterMoves struct {
	Character
	Moves []Move
}

func NewCharacter(
	id int,
	shortName string,
	longName string,
	fightingStyle string,
	nationality string,
	height int,
	weight int,
	gender string,
) Character {
	return Character{
		Id:            id,
		ShortName:     shortName,
		LongName:      longName,
		FightingStyle: fightingStyle,
		Nationality:   nationality,
		Height:        height,
		Weight:        weight,
		Gender:        gender,
	}
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

func (cs ConnectionString) Get() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.dbname)
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

func NewConnectionString(username string, password string, host string, port int, dbname string) ConnectionString {
	return ConnectionString{
		username: username,
		password: password,
		host:     host,
		port:     port,
		dbname:   dbname,
	}
}
