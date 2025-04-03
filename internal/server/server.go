package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/APIs/github"
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/APIs/stepik"
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/APIs/vk"
)

type App struct {
}

func New() *App {
	return &App{}
}

type tokens struct {
	vk       string
	github   string
	stepik   string
	stepikID string
}

type Resume struct {
	LastName     string
	FirstName    string
	Age          int
	Email        string
	GitHubLogin  string
	Certificates int
}

func getTokensHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/getTokens.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func calcAge(bdate string) int {
	layout := "1.1.2006"
	birthdate, err := time.Parse(layout, bdate)
	if err != nil {
		log.Fatalf("ошибка парсинга даты рождения: %s", err)
	}

	now := time.Now()
	age := now.Year() - birthdate.Year()

	if now.YearDay() < birthdate.YearDay() {
		age--
	}

	return age
}

func processTokens(t tokens) {
	vkData, err := vk.GetData(t.vk)
	if err != nil {
		log.Fatalf("Ошибка обработки данных из вк: %s", err)
	}

	githubData, err := github.GetData(t.github)
	if err != nil {
		log.Fatalf("Ошибка обработки данных из гитхаба: %s", err)
	}

	st, _ := strconv.Atoi(t.stepikID)
	stepikData, err := stepik.GetData(st, t.stepik)
	if err != nil {
		log.Fatalf("Ошибка обработки данных из степика: %s", err)
	}

	age := calcAge(vkData.Bdate)

	data := Resume{
		LastName:     vkData.Lname,
		FirstName:    vkData.Fname,
		Age:          age,
		Email:        githubData.Email,
		GitHubLogin:  githubData.Login,
		Certificates: stepikData.Certificates,
	}

	err = generateAndSaveResume(data)
	if err != nil {
		log.Fatalf("Ошибка геренации резюме: %s", err)
	}
}

func generateAndSaveResume(data Resume) error {
	outputPath := "public/resume.html"

	tmpl, err := template.ParseFiles("templates/resume.html")
	if err != nil {
		return err
	}

	resume, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer resume.Close()

	err = tmpl.Execute(resume, data)
	if err != nil {
		return err
	}

	return nil
}

func showResumeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/resume.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка при обработке формы", http.StatusBadRequest)
		return
	}

	t := tokens{
		vk:       r.FormValue("token1"),
		github:   r.FormValue("token2"),
		stepik:   r.FormValue("token3"),
		stepikID: r.FormValue("login"),
	}

	processTokens(t)

	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
}

func (a *App) GetInfo() {

	http.HandleFunc("/resume/generate", getTokensHandler)
	http.HandleFunc("/submit", generateHandler)
	http.HandleFunc("/", showResumeHandler)

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}
