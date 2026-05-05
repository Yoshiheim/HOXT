package handlers

import (
	"fmt"
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"text/template"
)

// The Index aka '/' path in website.
// path: 'http://<HOST>:<PORT>/'
func MainPage(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParsePageID(r)
	if err != nil {
		http.Error(w, "error with args", http.StatusBadRequest)
		return
	}

	/*
		var tops []modules.Topic

		act := db.DB.Find(&tops)
		if act.Error != nil {
			log.Println(act.Error.Error())
			http.Error(w, "Error With DB.", http.StatusInternalServerError)
			return
		}
	*/

	var pastes []modules.Paste

	page := int(id)
	limit := 30
	offset := (page - 1) * limit

	//.Order("is_titled DESC").Order("created_at DESC").
	act := db.DB.Order("is_titled DESC").
		Order("created_at DESC").Offset(offset).Limit(limit).Find(&pastes)
	for i := range pastes {
		pastes[i].Title = html.EscapeString(pastes[i].Title)
	}
	if act.Error != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	var count int64

	db.DB.Model(&modules.Paste{}).Count(&count)

	totalPages := (count + int64(limit) - 1) / int64(limit)

	/*var tops []TopicWithCount

	if err := db.DB.
		Model(&modules.Topic{}).
		Select(`
		topics.*,
		(SELECT COUNT(*) FROM pastes WHERE pastes.topic_id = topics.id) as post_count
	`).Scan(&tops); err.Error != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	for i := range tops {
		tops[i].Name = html.EscapeString(tops[i].Name)
		tops[i].Description = html.EscapeString(tops[i].Description)
	}
	*/

	cfg := data.GetDConfig(w)

	tpl, err := template.New("index.html").Funcs(helpers.FuncMap).ParseFiles("./templates/index.html", "./templates/attr.html", "./templates/search.html", "./templates/interface.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"data":       cfg,
		"logo":       string(data.Logo),
		"timer":      helpers.Dest,
		"pastes":     pastes,
		"totalPages": int(totalPages),
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
