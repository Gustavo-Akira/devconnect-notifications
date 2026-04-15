package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotification(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		userID  string
		target  string
		message string
		channel Channel
	}{
		{
			name:    "builds notification with provided fields",
			id:      "notification-123",
			userID:  "user-123",
			target:  "user@example.com",
			message: "welcome",
			channel: EmailChannel,
		},
		{
			name:    "preserves empty values because constructor does not validate",
			id:      "",
			userID:  "",
			target:  "",
			message: "",
			channel: SMSChannel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNotification(tt.id, tt.userID, tt.target, tt.message, tt.channel)

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.id, got.Id)
			assert.Equal(t, tt.userID, got.UserId)
			assert.Equal(t, tt.target, got.Target)
			assert.Equal(t, tt.message, got.Message)
			assert.Equal(t, tt.channel, got.Channel)
			assert.Equal(t, Pending, got.Status)
		})
	}
}

func TestNotificationMarkAsSent(t *testing.T) {
	tests := []struct {
		name          string
		initialStatus Status
	}{
		{
			name:          "changes status from pending to sent",
			initialStatus: Pending,
		},
		{
			name:          "overrides failed status to sent",
			initialStatus: Failed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notification := &Notification{Status: tt.initialStatus}

			notification.MarkAsSent()

			assert.Equal(t, Sent, notification.Status)
		})
	}
}

func TestNotificationMarkAsFailed(t *testing.T) {
	tests := []struct {
		name          string
		initialStatus Status
	}{
		{
			name:          "changes status from pending to failed",
			initialStatus: Pending,
		},
		{
			name:          "overrides sent status to failed",
			initialStatus: Sent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notification := &Notification{Status: tt.initialStatus}

			notification.MarkAsFailed()

			assert.Equal(t, Failed, notification.Status)
		})
	}
}
