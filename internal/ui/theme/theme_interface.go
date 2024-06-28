package theme

type ThemeInterface interface {
	Background() string
	Foreground() string
	Primary() string
	Secondary() string
	Success() string
	Info() string
	Warning() string
	Danger() string
	Light() string
	Dark() string
}