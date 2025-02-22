package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/gui"
)

type api struct {
}

func New() *api {
	return &api{}
}

const (
	vkURLTemplate = "https://api.vk.com/method/users.get?fields=first_name,last_name&access_token=%s&v=5.131"
	githubURL     = "https://api.github.com/user"
	stepikURL     = "https://stepik.org/api/users/me"
	outputFile    = "user_data.txt"
)

type VKResponse struct {
	Response []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

type GitHubResponse struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StepikResponse struct {
	Users []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"users"`
}

type UserData struct {
	VK     string `json:"vk"`
	GitHub string `json:"github"`
	Stepik string `json:"stepik"`
}

func (a *api) GetInfo() {
	gui := gui.New()

	tokens := gui.GetTokens()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vkURL := fmt.Sprintf(vkURLTemplate, tokens.Vk)
	vkJson, err := makeVKRequest(ctx, vkURL)
	if err != nil {
		log.Printf("Ошибка VK: %s", err)
	}

	githubJson, err := MakeRequest(ctx, tokens.Github, githubURL)
	if err != nil {
		log.Printf("Ошибка GitHub: %s", err)
	}

	stepikJson, err := MakeRequest(ctx, tokens.Stepik, stepikURL)
	if err != nil {
		log.Printf("Ошибка Stepik: %s", err)
	}

	var vkData VKResponse
	if err := json.Unmarshal(vkJson, &vkData); err != nil {
		log.Printf("Ошибка парсинга VK: %s", err)
	} else if len(vkData.Response) == 0 {
		log.Printf("VK API вернул пустой ответ")
	}

	var githubData GitHubResponse
	if err := json.Unmarshal(githubJson, &githubData); err != nil {
		log.Printf("Ошибка парсинга GitHub: %s", err)
	}

	var stepikData StepikResponse
	if err := json.Unmarshal(stepikJson, &stepikData); err != nil {
		log.Printf("Ошибка парсинга Stepik: %s", err)
	}

	userData := UserData{
		VK:     fmt.Sprintf("%s %s", vkData.Response[0].FirstName, vkData.Response[0].LastName),
		GitHub: githubData.Login,
		Stepik: fmt.Sprintf("%s %s", stepikData.Users[0].FirstName, stepikData.Users[0].LastName),
	}

	if err := saveToFile(userData); err != nil {
		log.Printf("Ошибка сохранения данных в файл: %s", err)
	} else {
		fmt.Println("Данные успешно сохранены в файл:", outputFile)
	}
}

func makeVKRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("VK API вернул статус %d, ответ: %s", resp.StatusCode, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %w", err)
	}

	return data, nil
}

func MakeRequest(ctx context.Context, token, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API вернул статус %d, ответ: %s", resp.StatusCode, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %w", err)
	}

	return data, nil
}

func saveToFile(data UserData) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf("VK: %s\nGitHub: %s\nStepik: %s\n", data.VK, data.GitHub, data.Stepik)

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("ошибка при записи данных в файл: %w", err)
	}

	return nil
}
