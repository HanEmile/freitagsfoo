package db

import (
	"fmt"
	"time"
	"git.darknebu.la/chaosdorf/freitagsfoo/src/structs"
	pg "github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

// InsertTalk inserts the given talk into the database
func InsertTalk(db *pg.DB, talk *structs.Talk) error {
	err := db.Insert(talk)
	if err != nil {
		return fmt.Errorf("could not insert talk into the db: %s", err)
	}

	return nil
}

// UpcomingTalksLimited returns the next 3 upcoming talks
func UpcomingTalksLimited(db *pg.DB) ([]structs.Talk, error) {
	var talks []structs.Talk
	err := db.Model(&talks).Order("date ASC").Where("date > ?", time.Now()).Limit(3).Select()
	if err != nil {
		return []structs.Talk{}, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return talks, nil
}

// UpcomingTalks returns the next upcoming talks
func UpcomingTalks(db *pg.DB) ([]structs.Talk, error) {
	var talks []structs.Talk
	err := db.Model(&talks).Order("date ASC").Where("date > ?", time.Now()).Select()
	if err != nil {
		return []structs.Talk{}, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return talks, nil
}

// PastTalksLimited returns the past 3 talks
func PastTalksLimited(db *pg.DB) ([]structs.Talk, error) {
	var talks []structs.Talk
	err := db.Model(&talks).Order("date DESC").Where("date < ?", time.Now()).Limit(3).Select()
	if err != nil {
		return []structs.Talk{}, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return talks, nil
}

// PastTalks returns the next upcoming talks
func PastTalks(db *pg.DB) ([]structs.Talk, error) {
	var talks []structs.Talk
	err := db.Model(&talks).Order("date DESC").Where("date < ?", time.Now()).Select()
	if err != nil {
		return []structs.Talk{}, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return talks, nil
}

// CountUpcomingTalks counts the amount of talks upcoming
func CountUpcomingTalks(db *pg.DB) (int, error) {
	var talks []structs.Talk
	count, err := db.Model(&talks).Where("date > ?", time.Now()).Where("upcoming = ?", true).SelectAndCount()
	if err != nil {
		return -1, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return count, nil
}

// CountPastTalks counts the amount of past talks
func CountPastTalks(db *pg.DB) (int, error) {
	var talks []structs.Talk
	count, err := db.Model(&talks).Where("date < ?", time.Now()).SelectAndCount()
	if err != nil {
		return -1, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return count, nil
}

// TalkByUUID returns the talk with the given UUID
func TalkByUUID(db *pg.DB, uuidString string) (structs.Talk, error) {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return structs.Talk{}, fmt.Errorf("could not parse the UUID: %s", err)
	}

	var talk structs.Talk
	err = db.Model(&talk).Where("uuid = ?", parsedUUID).Select()
	if err != nil {
		return structs.Talk{}, fmt.Errorf("could not get the talks from the db: %s", err)
	}

	return talk, nil
}
