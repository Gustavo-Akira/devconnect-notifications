package domain

type Notification struct {
	Id      string
	UserId  string
	Target  string
	Message string
	Channel Channel
	Status  Status
}

type Channel string

const (
	EmailChannel Channel = "EMAIL"
	SMSChannel   Channel = "SMS"
)

func NewNotification(id, userID, target, message string, channel Channel) (*Notification, error) {
	return &Notification{
		Id:      id,
		UserId:  userID,
		Target:  target,
		Channel: channel,
		Message: message,
		Status:  Pending,
	}, nil
}

func (n *Notification) MarkAsSent() {
	n.Status = Sent
}

func (n *Notification) MarkAsFailed() {
	n.Status = Failed
}

type Status string

const (
	Pending Status = "PENDING"
	Sent    Status = "SENT"
	Failed  Status = "FAILED"
)
