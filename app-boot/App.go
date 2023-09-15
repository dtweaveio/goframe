package app_boot

// App is the main structure of a cli application.
// It is recommended that an app be created with the app_boot.NewApp() function.
type App struct {
	name      string
	shortName string
	desc      string
}
