package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UIApp struct {
	app        fyne.App
	window     fyne.Window
	logBinding binding.String
	results    []Result
}

func NewUIApp() *UIApp {
	return &UIApp{
		app:        app.New(),
		logBinding: binding.NewString(),
		results:    make([]Result, 0),
	}
}

func (ui *UIApp) Run() {
	ui.window = ui.app.NewWindow(fmt.Sprintf("%s %s", appName, appVersion))
	ui.window.Resize(fyne.NewSize(900, 600))

	// Create tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Search", ui.createSearchTab()),
		container.NewTabItem("Settings", ui.createSettingsTab()),
		container.NewTabItem("Log", ui.createLogTab()),
	)

	ui.window.SetContent(tabs)
	ui.window.ShowAndRun()
}

func (ui *UIApp) createSearchTab() fyne.CanvasObject {
	// NZBLNK Input
	nzblnkEntry := widget.NewEntry()
	nzblnkEntry.SetPlaceHolder("nzblnk://?h=...")

	// Manual input fields
	headerEntry := widget.NewEntry()
	headerEntry.SetPlaceHolder("Subject/Header to search for")

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Title for the NZB file")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password (optional)")

	groupsEntry := widget.NewEntry()
	groupsEntry.SetPlaceHolder("Newsgroups (comma separated)")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("DD.MM.YYYY or Unix timestamp")

	categoryEntry := widget.NewEntry()
	categoryEntry.SetPlaceHolder("Category (optional)")

	// Results area
	resultsLabel := widget.NewLabel("Results will appear here after search")
	resultsLabel.Wrapping = fyne.TextWrapWord
	resultsScroll := container.NewScroll(resultsLabel)
	resultsScroll.SetMinSize(fyne.NewSize(800, 200))

	// Search button
	searchBtn := widget.NewButtonWithIcon("Start Search", theme.SearchIcon(), func() {
		ui.performSearch(nzblnkEntry.Text, headerEntry.Text, titleEntry.Text, 
			passwordEntry.Text, groupsEntry.Text, dateEntry.Text, categoryEntry.Text, resultsLabel)
	})
	searchBtn.Importance = widget.HighImportance

	// Layout
	form := container.NewVBox(
		widget.NewCard("NZBLNK", "", container.NewVBox(
			nzblnkEntry,
		)),
		widget.NewSeparator(),
		widget.NewCard("Manual Search Parameters", "", container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("Header/Subject", headerEntry),
				widget.NewFormItem("Title", titleEntry),
				widget.NewFormItem("Password", passwordEntry),
				widget.NewFormItem("Groups", groupsEntry),
				widget.NewFormItem("Date", dateEntry),
				widget.NewFormItem("Category", categoryEntry),
			),
		)),
		searchBtn,
		widget.NewCard("Results", "", resultsScroll),
	)

	return container.NewScroll(form)
}

func (ui *UIApp) createSettingsTab() fyne.CanvasObject {
	// Load current config
	targetEntry := widget.NewEntry()
	targetEntry.SetPlaceHolder("execute, sabnzbd, nzbget, synologyds")
	if conf.General.Target != "" {
		targetEntry.SetText(conf.General.Target)
	}

	debugCheck := widget.NewCheck("Enable Debug Logging", nil)
	debugCheck.Checked = conf.General.Debug

	skipFailedCheck := widget.NewCheck("Skip Failed NZBs", nil)
	skipFailedCheck.Checked = conf.Nzbcheck.SkipFailed

	bestNzbCheck := widget.NewCheck("Find Best NZB", nil)
	bestNzbCheck.Checked = conf.Nzbcheck.BestNZB

	// SABnzbd settings
	sabHostEntry := widget.NewEntry()
	sabHostEntry.SetPlaceHolder("localhost")
	if conf.SABnzbd.Host != "" {
		sabHostEntry.SetText(conf.SABnzbd.Host)
	}

	sabPortEntry := widget.NewEntry()
	sabPortEntry.SetPlaceHolder("8080")
	if conf.SABnzbd.Port > 0 {
		sabPortEntry.SetText(fmt.Sprintf("%d", conf.SABnzbd.Port))
	}

	sabKeyEntry := widget.NewPasswordEntry()
	sabKeyEntry.SetPlaceHolder("API Key")
	if conf.SABnzbd.Nzbkey != "" {
		sabKeyEntry.SetText(conf.SABnzbd.Nzbkey)
	}

	sabSSLCheck := widget.NewCheck("Use SSL", nil)
	sabSSLCheck.Checked = conf.SABnzbd.Ssl

	// Save button
	saveBtn := widget.NewButton("Save Settings", func() {
		// This would save to config file
		ui.logMessage("Settings saved (config file update not implemented in UI mode)")
	})
	saveBtn.Importance = widget.HighImportance

	form := container.NewVBox(
		widget.NewCard("General", "", container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("Target", targetEntry),
			),
			debugCheck,
		)),
		widget.NewSeparator(),
		widget.NewCard("NZB Check", "", container.NewVBox(
			skipFailedCheck,
			bestNzbCheck,
		)),
		widget.NewSeparator(),
		widget.NewCard("SABnzbd", "", container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("Host", sabHostEntry),
				widget.NewFormItem("Port", sabPortEntry),
				widget.NewFormItem("API Key", sabKeyEntry),
			),
			sabSSLCheck,
		)),
		layout.NewSpacer(),
		saveBtn,
	)

	return container.NewScroll(form)
}

func (ui *UIApp) createLogTab() fyne.CanvasObject {
	logEntry := widget.NewEntryWithData(ui.logBinding)
	logEntry.MultiLine = true
	logEntry.Wrapping = fyne.TextWrapWord
	logEntry.Disable()

	clearBtn := widget.NewButton("Clear Log", func() {
		ui.logBinding.Set("")
	})

	return container.NewBorder(nil, clearBtn, nil, nil, container.NewScroll(logEntry))
}

func (ui *UIApp) logMessage(message string) {
	current, _ := ui.logBinding.Get()
	timestamp := time.Now().Format("15:04:05")
	newLog := fmt.Sprintf("[%s] %s\n%s", timestamp, message, current)
	ui.logBinding.Set(newLog)
}

func (ui *UIApp) performSearch(nzblnk, header, title, password, groups, date, category string, resultsLabel *widget.Label) {
	ui.logMessage("Starting search...")
	resultsLabel.SetText("Searching...")

	// Parse inputs
	if nzblnk != "" {
		args.Nzblnk = nzblnk
	}
	if header != "" {
		args.Header = header
	}
	if title != "" {
		args.Title = title
	}
	if password != "" {
		args.Password = password
	}
	if groups != "" {
		args.Groups = strings.Split(groups, ",")
		for i := range args.Groups {
			args.Groups[i] = strings.TrimSpace(args.Groups[i])
		}
	}
	if date != "" {
		args.Date = date
	}
	if category != "" {
		args.Category = category
	}

	// Check arguments
	if err := validateSearchArgs(); err != nil {
		ui.logMessage(fmt.Sprintf("Error: %v", err))
		resultsLabel.SetText(fmt.Sprintf("Error: %v", err))
		return
	}

	// Run search in goroutine to not block UI
	go func() {
		ui.results = make([]Result, 0)
		
		for _, name := range conf.Searchengines {
			ui.logMessage(fmt.Sprintf("Searching on %s...", searchEngines[name].name))
			
			if err := searchEngines[name].search(searchEngines[name], searchEngines[name].name); err != nil {
				ui.logMessage(fmt.Sprintf("Warning: %v", err))
			}
		}

		// Display results
		if len(results) > 0 {
			resultText := fmt.Sprintf("Found %d result(s):\n\n", len(results))
			for i, result := range results {
				resultText += fmt.Sprintf("%d. Source: %s\n", i+1, result.SearchEngine)
				if result.Nzb != nil && len(result.Nzb.Files) > 0 {
					resultText += fmt.Sprintf("   Subject: %s\n", result.Nzb.Files[0].Subject)
					resultText += fmt.Sprintf("   Size: %s\n", prettyByteSize(result.Nzb.Bytes))
					resultText += fmt.Sprintf("   Files: %d/%d (Missing: %d)\n", 
						result.Nzb.Files.Len(), result.Nzb.TotalFiles, result.FilesMissing)
					resultText += fmt.Sprintf("   Segments: %d/%d (Missing: %.2f%%)\n\n", 
						result.Nzb.Segments, result.Nzb.TotalSegments, result.SegmentsMissingPercent)
				}
			}
			resultsLabel.SetText(resultText)
			ui.logMessage(fmt.Sprintf("Search completed. Found %d result(s)", len(results)))
		} else {
			resultsLabel.SetText("No results found")
			ui.logMessage("Search completed. No results found.")
		}
	}()
}

func validateSearchArgs() error {
	if args.Header == "" && args.Nzblnk == "" {
		return fmt.Errorf("you must provide either a subject or a NZBLNK URI")
	}
	
	// Parse NZBLNK if provided
	if args.Nzblnk != "" {
		if err := parseNzblnk(); err != nil {
			return err
		}
	}
	
	return nil
}

func parseNzblnk() error {
	// This should call the existing NZBLNK parsing logic from arguments.go
	// For now, just a basic implementation
	if !strings.HasPrefix(args.Nzblnk, "nzblnk://") {
		return fmt.Errorf("invalid NZBLNK URI format")
	}
	return nil
}
