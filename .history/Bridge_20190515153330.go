package Structural

import "fmt"


type TimeImp interface {
	Tell()
}

type BasicTimeImp struct {
	hour   int
	minute int
}

func NewBasicTimeImp(hour, minute int) TimeImp {
	return &BasicTimeImp{hour, minute}
}

func (t *BasicTimeImp) Tell() {
	fmt.Fprintf(outputWriter, "The time is %2.2d:%2.2d\n", t.hour, t.minute)
}

type CivilianTimeImp struct {
	BasicTimeImp
	meridiem string
}

func NewCivilianTimeImp(hour, minute int, pm bool) TimeImp {
	meridiem := "AM"
	if pm {
		meridiem = "PM"
	}
	time := &CivilianTimeImp{meridiem: meridiem}
	time.hour = hour
	time.minute = minute
	return time
}

func (t *CivilianTimeImp) Tell() {
	fmt.Fprintf(outputWriter, "The time is %2.2d:%2.2d %s\n", t.hour, t.minute, t.meridiem)
}

type ZuluTimeImp struct {
	BasicTimeImp
	zone string
}

func NewZuluTimeImp(hour, minute, zoneID int) TimeImp {
	zone := ""
	if zoneID == 5 {
		zone = "Eastern Standard Time"
	} else if zoneID == 6 {
		zone = "Central Standard Time"
	}
	time := &ZuluTimeImp{zone: zone}
	time.hour = hour
	time.minute = minute
	return time
}

func (t *ZuluTimeImp) Tell() {
	fmt.Fprintf(outputWriter, "The time is %2.2d:%2.2d %s\n", t.hour, t.minute, t.zone)
}

type Time struct {
	imp TimeImp
}

func NewTime(hour, minute int) *Time {
	return &Time{imp: NewBasicTimeImp(hour, minute)}
}

func (t *Time) Tell() {
	t.imp.Tell()
}

type CivilianTime struct {
	Time
}

func NewCivilianTime(hour, minute int, pm bool) *Time {
	time := &CivilianTime{}
	time.imp = NewCivilianTimeImp(hour, minute, pm)
	return &time.Time
}

// ZuluTime is a time with a Zulu time implementation.
type ZuluTime struct {
	Time
}

// NewZuluTime creates a new Zulu time.
func NewZuluTime(hour, minute, zoneID int) *Time {
	time := &ZuluTime{}
	time.imp = NewZuluTimeImp(hour, minute, zoneID)
	return &time.Time
}
