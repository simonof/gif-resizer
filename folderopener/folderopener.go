package folderopener

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func OpenFolder() (string, int, int, int) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Folder Opener")

	// Set the window size here
	myWindow.Resize(fyne.NewSize(600, 400))

	var (
		folderPath string
		value1     int
		value2     int
		value3     int
	)

	openFolderButton := widget.NewButton("Open Folder", func() {
		folderDialog := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
			if err == nil && folder != nil {
				folderPath = folder.Path()
				myWindow.Close()
			}
		}, myWindow)
		folderDialog.Show()
	})

	slider1Label := widget.NewLabel("Image: 0")
	slider1 := widget.NewSlider(1, 999)
	slider1.OnChanged = func(value float64) {
		value1 = int(value)
		slider1Label.SetText(fmt.Sprintf("Slider 1: %d", value1))
	}

	slider2Label := widget.NewLabel("Delay: 0")
	slider2 := widget.NewSlider(1, 999)
	slider2.OnChanged = func(value float64) {
		value2 = int(value)
		slider2Label.SetText(fmt.Sprintf("Slider 1: %d", value2))
	}

	slider3Label := widget.NewLabel("Disposal: 0")
	slider3 := widget.NewSlider(1, 999)
	slider3.OnChanged = func(value float64) {
		value3 = int(value)
		slider3Label.SetText(fmt.Sprintf("Slider 1: %d", value3))
	}

	content := container.NewVBox(
		openFolderButton,
		slider1Label,
		slider1,
		slider2Label,
		slider2,
		slider3Label,
		slider3,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	return folderPath, value1, value2, value3
}
