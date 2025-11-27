package lessons

type Event struct {
	Title    string
	Date     string
	Location string
}

func CreateGoEvent() Event {
	event := Event{
		Title:    "День рождения Golang",
		Date:     "10 ноября 2009",
		Location: "GoogleLand",
	}
	return event
}
