package app

func RunApp() {
	myWindow := setupWin()

	content := createContent()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
