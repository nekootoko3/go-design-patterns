package ticket_maker

type ticketMaker int

var instance *ticketMaker

func GetInstance() *ticketMaker {
	if instance == nil {
		var tm ticketMaker = 1000
		instance = &tm
	}
	return instance
}

func (tm *ticketMaker) GetNextTicketNumber() int {
	*tm += 1
	return int(*tm)
}
