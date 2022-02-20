package app

func RunApp() {
	myWindow := setupWin()

	content := createContent(myWindow)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
