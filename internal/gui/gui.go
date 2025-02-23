package gui

import (
	"embed"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed icons/*
var iconFS embed.FS

var (
	vkIcon     fyne.Resource
	githubIcon fyne.Resource
	stepikIcon fyne.Resource
)

type Tokens struct {
	Vk       string
	Github   string
	Stepik   string
	StepikID int
}

type gui struct{}

func New() *gui {
	vkIcon = loadIcon("icons/vk.svg")
	githubIcon = loadIcon("icons/github.svg")
	stepikIcon = loadIcon("icons/stepik.svg")
	return &gui{}
}

func loadIcon(name string) fyne.Resource {
	data, err := iconFS.ReadFile(name)
	if err != nil {
		fmt.Printf("Ошибка загрузки иконки: %v\n", err)
		return theme.ErrorIcon()
	}
	return fyne.NewStaticResource(name, data)
}

func (g *gui) GetTokens() *Tokens {
	myApp := app.New()
	myWindow := myApp.NewWindow("Token Manager")
	myWindow.Resize(fyne.NewSize(400, 300))

	var (
		vkToken     string
		githubToken string
		stepikToken string
		stepikID    string
	)

	header := container.NewCenter(
		container.NewHBox(
			widget.NewIcon(vkIcon),
			widget.NewIcon(githubIcon),
			widget.NewIcon(stepikIcon),
			widget.NewLabel("Получение данных по токенам"),
		),
	)

	createInputField := func(icon fyne.Resource, labelText string, isPassword bool) (*widget.Entry, fyne.CanvasObject) {
		var entry *widget.Entry
		if isPassword {
			entry = widget.NewPasswordEntry()
			entry.SetPlaceHolder("Введите токен...")
			entry.Validator = func(s string) error {
				if len(s) < 10 {
					return fmt.Errorf("минимум 10 символов")
				}
				return nil
			}
		} else {
			entry = widget.NewEntry()
			entry.SetPlaceHolder("Введите логин...")
		}

		return entry, container.NewBorder(
			nil,
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewIcon(icon),
				widget.NewLabel(labelText),
				layout.NewSpacer(),
			),
			nil,
			entry,
		)
	}

	vkEntry, vkBox := createInputField(vkIcon, "VK:", true)
	githubEntry, githubBox := createInputField(githubIcon, "GitHub:", true)
	stepikEntry, stepikBox := createInputField(stepikIcon, "Stepik Token:", true)
	stepikIDEntry, stepikIDBox := createInputField(stepikIcon, "Stepik ID:", false)

	saveBtn := widget.NewButtonWithIcon("Сохранить", theme.DocumentSaveIcon(), func() {
		if err := vkEntry.Validate(); err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		if err := githubEntry.Validate(); err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		if err := stepikEntry.Validate(); err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		vkToken = vkEntry.Text
		githubToken = githubEntry.Text
		stepikToken = stepikEntry.Text
		stepikID = stepikIDEntry.Text

		dialog.ShowInformation("Успех!", "Токены и логин сохранены", myWindow)
	})

	content := container.NewVBox(
		header,
		layout.NewSpacer(),
		vkBox,
		githubBox,
		stepikBox,
		stepikIDBox,
		layout.NewSpacer(),
		container.NewCenter(saveBtn),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	stepikIDint, _ := strconv.Atoi(stepikID)
	return &Tokens{
		Vk:       vkToken,
		Github:   githubToken,
		Stepik:   stepikToken,
		StepikID: stepikIDint,
	}
}
