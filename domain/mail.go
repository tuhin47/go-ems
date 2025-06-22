package domain

import "github.com/tuhin47/go-ems/models"

type (
	MailService interface {
		SendInvitationEmail(userIds []int, event *models.Event) error
		EnqueueEventReminderEmailNotification(event *models.Event) error
	}
)
