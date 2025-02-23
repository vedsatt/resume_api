package stepik

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StepikData struct {
	Name         string `json:"name"`
	SolvedTasks  int    `json:"solved_tasks"`
	Certificates int    `json:"certificates"`
}

type UserResponse struct {
	Users []struct {
		FullName string `json:"full_name"`
	} `json:"users"`
}

type CourseGradeResponse struct {
	CourseGrades []struct {
		Score               int  `json:"score"`
		IsCertificateIssued bool `json:"is_certificate_issued"`
	} `json:"course-grades"`
}

func GetData(userID int, accessToken string) (*StepikData, error) {
	userURL := fmt.Sprintf("https://stepik.org/api/users/%d/", userID)
	req, err := http.NewRequest("GET", userURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка Stepik API при запросе данных пользователя: %d", resp.StatusCode)
	}

	var userData UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return nil, err
	}

	userName := "Не указано"
	if len(userData.Users) > 0 {
		userName = userData.Users[0].FullName
	}

	coursesURL := fmt.Sprintf("https://stepik.org/api/course-grades?user=%d", userID)
	req, err = http.NewRequest("GET", coursesURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка Stepik API при запросе данных о курсах: %d", resp.StatusCode)
	}

	var coursesData CourseGradeResponse
	err = json.NewDecoder(resp.Body).Decode(&coursesData)
	if err != nil {
		return nil, err
	}

	totalSolvedTasks := 0
	totalCertificates := 0

	for _, course := range coursesData.CourseGrades {
		totalSolvedTasks += course.Score
		if course.IsCertificateIssued {
			totalCertificates++
		}
	}

	return &StepikData{
		Name:         userName,
		SolvedTasks:  totalSolvedTasks,
		Certificates: totalCertificates,
	}, nil
}
