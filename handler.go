package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


type User struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{
			Name:   "Chris",
			Age:    22,
			Gender: "Male",
		},
		{
			Name:   "Annie",
			Age:    23,
			Gender: "Female",
		},
		{
			Name:   "Jane",
			Age:    25,
			Gender: "Female",
		},
	}

	usersJson, _ := json.Marshal(users)
	w.Write(usersJson)
}

// GetTexts возвращает текстовое представление элементов
// по selector на странице c переданным url
func GetTexts(url, selector string) ([]string, error) {

	// Скачиваем html-страницу
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Считываем страницу в goquery-документ
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Находим все элементы с переданным селектором
	// и сохраняем их содержимое в список
	var texts []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})

	return texts, nil
}

func wikiHandler(w http.ResponseWriter, r *http.Request) {

	urlwiki := "https://ru.wikipedia.org/wiki/%D0%9F%D1%83%D1%82%D0%B8%D0%BD,_%D0%92%D0%BB%D0%B0%D0%B4%D0%B8%D0%BC%D0%B8%D1%80_%D0%92%D0%BB%D0%B0%D0%B4%D0%B8%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B8%D1%87"

	titles, err := GetTexts(urlwiki, "b")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Последние заголовки:")
	for _, t := range titles {
		fmt.Println("-", t)
	}

	usersJson, _ := json.Marshal(titles)
	w.Write(usersJson)
}