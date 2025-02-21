package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	vkIcon     = loadIcon("vk.svg")
	githubIcon = loadIcon("github.svg")
	stepikIcon = loadIcon("stepik.svg")
)

func loadIcon(name string) fyne.Resource {
	data, err := fyne.LoadResourceFromPath("icons/" + name)
	if err != nil {
		fmt.Printf("Error loading icon: %v\n", err)
		return theme.ErrorIcon()
	}
	return data
}

func Gui() []string {
	myApp := app.New()
	myWindow := myApp.NewWindow("Token Manager")
	myWindow.Resize(fyne.NewSize(400, 300))

	var (
		vkToken     string
		githubToken string
		stepikToken string
	)

	header := container.NewCenter(
		container.NewHBox(
			widget.NewIcon(vkIcon),
			widget.NewIcon(githubIcon),
			widget.NewIcon(stepikIcon),
			widget.NewLabel("Получение данных по токенам"),
		),
	)

	createInputField := func(icon fyne.Resource, labelText string) (*widget.Entry, fyne.CanvasObject) {
		entry := widget.NewPasswordEntry()
		entry.SetPlaceHolder("Введите токен...")
		entry.Validator = func(s string) error {
			if len(s) < 10 {
				return fmt.Errorf("минимум 10 символов")
			}
			return nil
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

	vkEntry, vkBox := createInputField(vkIcon, "VK:")
	githubEntry, githubBox := createInputField(githubIcon, "GitHub:")
	stepikEntry, stepikBox := createInputField(stepikIcon, "Stepik:")

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

		dialog.ShowInformation("Успех!", "Токены сохранены", myWindow)
	})

	content := container.NewVBox(
		header,
		layout.NewSpacer(),
		vkBox,
		githubBox,
		stepikBox,
		layout.NewSpacer(),
		container.NewCenter(saveBtn),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	tokens := []string{vkToken, githubToken, stepikToken}
	return tokens
}
