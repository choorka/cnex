package telegram

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Rate struct {
	Min  int     `json:"min"`
	Max  int     `json:"max"` // -1 означает "и выше"
	Rate float64 `json:"rate"`
}

type PostData struct {
	Rates   []Rate `json:"rates"`
	Hours   string `json:"hours"`
	RawText string `json:"raw"`
	Date    string `json:"date"`
}

func parseRates(text string) []Rate {
	var rates []Rate

	// Паттерн ищет все вхождения: "от <число>¥ — <число с запятой или точкой>"
	rateRegex := regexp.MustCompile(`от\s*(\d+)\s*¥\s*[—\-]\s*([\d.,]+)`)

	matches := rateRegex.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return rates
	}

	var minValues []int
	for _, m := range matches {
		min, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}

		rateStr := strings.ReplaceAll(m[2], ",", ".")
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			continue
		}

		minValues = append(minValues, min)
		rates = append(rates, Rate{Min: min, Rate: rate})
	}

	// Устанавливаем max для каждой записи
	for i := range rates {
		if i < len(rates)-1 {
			rates[i].Max = minValues[i+1] - 1
		} else {
			rates[i].Max = -1
		}
	}

	return rates
}


func parseHours(text string) string {
	hourRegex := regexp.MustCompile(`(\d{1,2}:\d{2})\s*до\s*(\d{1,2}:\d{2})`)
	if matches := hourRegex.FindStringSubmatch(text); len(matches) == 3 {
		return matches[1] + " - " + matches[2]
	}
	return ""
}

func parseDate(text string) string {
	dateRegex := regexp.MustCompile(`(\d{1,2}\.\d{1,2}\.\d{4})`)
	if matches := dateRegex.FindStringSubmatch(text); len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func GetLatestPostParsed() (*PostData, error) {
    res, err := http.Get("https://t.me/s/CNEXTrader")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

    messages := doc.Find(".tgme_widget_message_text")
    count := messages.Length()

    var result *PostData

    // Проходим с конца к началу
    for i := count - 1; i >= 0; i-- {
        s := messages.Eq(i)
        rawText := s.Text()
        rates := parseRates(rawText)
        if len(rates) == 0 {
            continue
        }

        result = &PostData{
            Rates:   rates,
            Hours:   parseHours(rawText),
            RawText: rawText,
            Date:    parseDate(rawText),
        }
        break // нашли последний валидный пост, выходим
    }

    if result == nil {
        return nil, fmt.Errorf("не найден актуальный пост с курсом")
    }

    return result, nil
}
