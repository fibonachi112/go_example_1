package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// Константы звуков уведомлений
const (
	SOUND_PING   = "Ping"   // Звук пинга
	SOUND_BEEP   = "Beeep"  // Звук бипа
	SOUND_FISH   = "Fish"   // Звук рыбы
	SOUND_GLASS  = "Glass"  // Звук стекла
	SOUND_BOTTLE = "Bottle" // Звук бутылки
	SOUND_FROG   = "Frog"   // Звук лягушки
	SOUND_MORSE  = "Morse"  // Звук морзянки
	SOUND_POP    = "Pop"    // Звук хлопка
	SOUND_PURR   = "Purr"   // Звук мурлыканья
)

// Notification описывает параметры уведомления
type Notification struct {
	Title         string // Заголовок уведомления
	Subtitle      string // Подзаголовок уведомления
	Sound         string // Звук уведомления
	EveryNMinutes int    // Периодичность показа в минутах
}

// sendOnce отправляет одиночное уведомление
func sendOnce(n Notification) error {
	script := fmt.Sprintf(`display notification %q with title %q subtitle %q sound name %q`,
		n.Title, n.Title, n.Subtitle, n.Sound)
	log.Println(script)
	return exec.Command("osascript", "-e", script).Run()
}

// schedule запускает периодическую отправку уведомлений
func schedule(ctx context.Context, n Notification) {
	ticker := time.NewTicker(time.Duration(n.EveryNMinutes) * time.Minute)
	defer ticker.Stop()

	// показываем уведомление
	_ = sendOnce(n)

	for {
		select {
		case <-ticker.C:
			if err := sendOnce(n); err != nil {
				fmt.Println("notify error:", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	glaza := Notification{"Глаза", "гимнастика глаз", SOUND_PING, 5}
	spina := Notification{"Спина", "Упражнения для спины", SOUND_BOTTLE, 2}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer stop()

	go schedule(ctx, glaza)
	go schedule(ctx, spina)

	<-ctx.Done()
}
