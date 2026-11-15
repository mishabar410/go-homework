package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type Notifier interface {
	Send(msg string) error
}

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Log() {
	fmt.Println("Sending notification via channel:", b.Channel)
}

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (e EmailNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMessage
	}
	fmt.Printf("[%s] %s: %s\n", e.Channel, e.Address, msg)
	return nil
}

var _ Notifier = EmailNotifier{}

type PushNotifier struct {
	BaseNotifier BaseNotifier
	DeviceID     uint64
}

func (p PushNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMessage
	}
	if p.DeviceID == 0 {
		return &SendError{Channel: p.BaseNotifier.Channel, Reason: "не задан device id"}
	}
	fmt.Printf("[%s] устройство %d: %s\n", p.BaseNotifier.Channel, p.DeviceID, msg)
	return nil
}

var _ Notifier = PushNotifier{}

type SendError struct {
	Channel string
	Reason  string
}

func (e *SendError) Error() string {
	return fmt.Sprintf("ошибка отправки через %s: %s", e.Channel, e.Reason)
}

var _ error = (*SendError)(nil)

type BrokenNotifier struct{}

func (BrokenNotifier) Send(string) error {
	panic("драйвер канала повреждён")
}

type Channel struct {
	Name   string
	Sender Notifier
}

func SafeSend(notifier Notifier, msg string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("паника при отправке: %v", r)
		}
	}()
	return notifier.Send(msg)
}

func Dispatch(all []Channel, msg string) []error {
	var failures []error
	for _, channel := range all {
		if err := SafeSend(channel.Sender, msg); err != nil {
			failures = append(failures, fmt.Errorf("канал %s: %w", channel.Name, err))
		}
	}
	return failures
}

func report(failures []error) {
	for _, err := range failures {
		fmt.Println("ошибка:", err)

		if errors.Is(err, ErrEmptyMessage) {
			fmt.Println("  сообщение пустое, введите текст уведомления")
			wrapped := fmt.Errorf("рассылка: %w", err)
			fmt.Println("  errors.Is после ещё одной обёртки:", errors.Is(wrapped, ErrEmptyMessage))
		}

		var sendErr *SendError
		if errors.As(err, &sendErr) {
			fmt.Println("  канал:", sendErr.Channel)
			fmt.Println("  причина:", sendErr.Reason)
		}
	}
	fmt.Println("Всего ошибок:", len(failures))
}

func main() {
	email := EmailNotifier{
		BaseNotifier: BaseNotifier{Channel: "email"},
		Address:      "example@mail.ru",
	}
	push := PushNotifier{
		BaseNotifier: BaseNotifier{Channel: "push"},
		DeviceID:     123456,
	}
	pushNoDevice := PushNotifier{
		BaseNotifier: BaseNotifier{Channel: "push"},
	}

	fmt.Println("Сценарий 1: email, broken, push")
	report(Dispatch([]Channel{
		{Name: "email", Sender: email},
		{Name: "broken", Sender: BrokenNotifier{}},
		{Name: "push", Sender: push},
	}, "Заказ готов"))

	fmt.Println()
	fmt.Println("Сценарий 2: пустое сообщение")
	report(Dispatch([]Channel{
		{Name: "email", Sender: email},
		{Name: "push", Sender: push},
	}, ""))

	fmt.Println()
	fmt.Println("Сценарий 3: push без device id")
	report(Dispatch([]Channel{
		{Name: "push", Sender: pushNoDevice},
	}, "Заказ готов"))
}
