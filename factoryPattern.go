package main

import "fmt"

// Problem Statement 1: Factory Pattern in Golang
// Problem Statement 2: How to override some function of a class while using base class's method as well

type Notification interface {
	Send()
}

type EmailNotification struct{}

func (n *EmailNotification) Send() {
	fmt.Println("Email Notification")
}

type SMSNotification struct{}

func (n *SMSNotification) Send() {
	fmt.Println("SMS Notification")
}

type SlackNotification struct{}

func (n *SlackNotification) Send() {
	fmt.Println("Slack Notification")
}

type NotificationService interface {
	notificationTypeCreator() Notification
	sendNotification()
}

type BaseNotificationService struct {
	self NotificationService
}

func (b *BaseNotificationService) notificationTypeCreator() Notification {
	panic("notificationTypeCreator must be implemented by derived service")
}

func (b *BaseNotificationService) sendNotification() {
	// Here this calls the actual service's overidden method
	notification := b.self.notificationTypeCreator()
	notification.Send()
}

type EmailNotificationService struct {
	// Here we have embedded this, so we can access its method in a inheritance way
	BaseNotificationService
}

func (e *EmailNotificationService) notificationTypeCreator() Notification {
	notification := EmailNotification{}
	return &notification
}

func NewEmailNotificationService() *EmailNotificationService {
	em := EmailNotificationService{}
	// Here we are setting it as self so it gets used in sendNotification Function
	em.self = &em
	return &em
}

type SMSNotificationService struct {
	BaseNotificationService
}

func (e *SMSNotificationService) notificationTypeCreator() Notification {
	notification := SMSNotification{}
	return &notification
}

func NewSMSNotificationService() *SMSNotificationService {
	em := SMSNotificationService{}
	em.self = &em
	return &em
}

type SlackNotificationService struct {
	BaseNotificationService
}

func (e *SlackNotificationService) notificationTypeCreator() Notification {
	notification := SlackNotification{}
	return &notification
}

func NewSlackNotificationService() *SlackNotificationService {
	em := SlackNotificationService{}
	em.self = &em
	return &em
}

func FactoryPatternExample() {
	emailService := NewEmailNotificationService()
	emailService.sendNotification()

	smsService := NewSMSNotificationService()
	smsService.sendNotification()

	slackService := NewSlackNotificationService()
	slackService.sendNotification()
}
