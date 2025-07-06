package domain

import (
	"github.com/tuhin47/go-ems/models"
	"github.com/tuhin47/go-ems/types"
)

type (
	MailService interface {
		SendEmail(reqData types.EmailPayload) error
		SendInvitationEmail(userIds []int, event *models.Event) error
		EnqueueEventReminderEmailNotification(event *models.Event) error
	}

	MailRepository interface {
		SendEmail(reqData *types.EmailPayload) error
	}
)
