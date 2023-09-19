package database

import (
	"fmt"
	"os"
	"strconv"
)

type CharacterWithId struct {
	Id        int
	ShortName string
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

func NewConnectionString() (cs ConnectionString, err error) {
	cs.username = os.Getenv("DB_USERNAME")
	cs.password = os.Getenv("DB_PASSWORD")
	cs.host = os.Getenv("DB_HOST")
	cs.port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	cs.dbname = os.Getenv("DB_DATABASE_NAME")
	return
}

func (cs ConnectionString) Get() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.dbname)
}
