package api

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/gui"
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/servers/github"
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/servers/stepik"
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/servers/vk"
)

type App struct {
}

func New() *App {
	return &App{}
}

func dataToFile(v *vk.UserData, g *github.GithubData, s *stepik.StepikData) {
	filePath := filepath.Join("output", "data.txt")

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Ошибка при создании файла: %s", err)
	}
	defer file.Close()

	vk := fmt.Sprintf("Данные из Vk\nЛогин: %s\nДата рождения: %s\nИмя: %s\nФамилия: %s\n\n", v.Domain, v.Bdate, v.Fname, v.Lname)
	github := fmt.Sprintf("Данные из GitHub\nЛогин: %s\nПочта: %s\n\n", g.Login, g.Email)
	stepik := fmt.Sprintf("Данные из Stepik\nИмя: %s\nКол-во решенных задач: %d\nКол-во сертификатов: %d\n", s.Name, s.SolvedTasks, s.Certificates)

	data := fmt.Sprintf("%s%s%s", vk, github, stepik)
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatalf("Ошибка при записи данных в файл: %s", err)
	}
}

func (a *App) GetInfo() {
	g := gui.New()
	tokens := g.GetTokens()

	vkData, err := vk.GetData(tokens.Vk)
	if err != nil {
		log.Fatalf("Ошибка получения данных из Vk: %s", err)
	}

	githubData, err := github.GetData(tokens.Github)
	if err != nil {
		log.Fatalf("Ошибка получения данных из GitHub: %s", err)
	}

	stepikData, err := stepik.GetData(tokens.StepikID, tokens.Stepik)
	if err != nil {
		log.Fatalf("Ошибка получения данных из Stepik: %s", err)
	}

	dataToFile(vkData, githubData, stepikData)
}
