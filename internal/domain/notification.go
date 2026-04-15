package domain

type Notification struct {
	Id      string
	UserId  string
	Email   string
	Message string
	Status  Status
}

func NewNotification(id, userID, email, message string) (*Notification, error) {
	return &Notification{
		Id:      id,
		UserId:  userID,
		Email:   email,
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
