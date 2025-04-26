package constants

import "github.com/charmbracelet/lipgloss"

var MainStyles = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Background(lipgloss.Color("236")).
	Padding(1, 2).
	Margin(1, 0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63"))

var SuccessStyles = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).
	Background(lipgloss.Color("22")).
	Padding(1, 2).
	Margin(1, 0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("28"))

var ErrorStyles = lipgloss.NewStyle().
	Foreground(lipgloss.Color("15")).
	Background(lipgloss.Color("1")).
	Padding(1, 2).
	Margin(1, 0).
	Border(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("9"))

var WarningStyles = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("3")).
	Padding(1, 2).
	Margin(1, 0).
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("11"))

var ELoadEnv = "Error loading .env file"
var Start = "___START PROJECT___"
var EAviaUrl = "Error loading AVIA_URL"
var SRunning = "✅ Service running"
var ExtraS = "initial"
var EAviaApi = "Error loading AVIA_API_KEY"
