package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	d := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	n := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return d.After(n)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("неверный формат даты: %v", err)
	}

	parts := strings.Split(repeat, " ")
	ruleType := parts[0]

	switch ruleType {
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateFormat), nil

	case "d":
		if len(parts) < 2 {
			return "", fmt.Errorf("отсутствует интервал для правила d")
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval <= 0 {
			return "", fmt.Errorf("неверный интервал дней: %s", parts[1])
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateFormat), nil

	default:
		return "", fmt.Errorf("неподдерживаемое правило повторения: %s", ruleType)
	}
}
