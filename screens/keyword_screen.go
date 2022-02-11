package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/indigo-sadland/well_aware/utils"
	"time"
)

func KeywordSearchScreen(_ fyne.Window) fyne.CanvasObject {

	var err error
	var url string

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter keyword to search for")

	ipfsButton := widget.NewButton("ipfs_search", func() {

		url = "https://ipfs-search.com/#/search?q=" + input.Text + "&page=1&last_seen"
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	pastesButton := widget.NewButton("pastes_leaks", func() {

		url = "https://www.google.com/search?q=site:pastebin.com | site:paste2.org | site:pastehtml.com " +
			"| site:slexy.org | site:snipplr.com | site:snipt.net | site:textsnip.com | site:bitpaste.app " +
			"| site:justpaste.it | site:heypasteit.com | site:hastebin.com | site:dpaste.org | site:dpaste.com " +
			"| site:phpfiddle.org | site:ide.geeksforgeeks.org | site:repl.it | site:ideone.com | site:paste.debian.net " +
			"| site:paste.org | site:paste.org.ru | site:codebeautify.org | site:trello.com " +
			"| site:papaly.com | site:0bin.net | site:paste-bin.xyz " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	shortUrlButton := widget.NewButton("short_url_search", func() {

		url = "https://shorteners.grayhatwarfare.com/shorteners?keywords=" + input.Text + "&order=&ext="
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	telegramSearchButton := widget.NewButton("telegram_search", func() {

		url = "https://lyzem.com/search?q=" + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	awsBucketsButton := widget.NewButton("s3_buckets", func() {

		url = "https://buckets.grayhatwarfare.com/results/" + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	mmntButton := widget.NewButton("mmnt_ftp", func() {

		url = "https://www.mmnt.net/get?cx=partner-pub-7093288927147322%3Asqzu2pvo68m&cof=FORID%3A10&ie=UTF-8&q=" +
			input.Text + "&sa=MMNT.net+search"
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	dorkftpButton := widget.NewButton("google_dork_ftp", func() {

		url = "https://www.google.com/search?&q=inurl:ftp -inurl:(http|https) " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	gdriveButton := widget.NewButton("google_drive", func() {

		url = "https://www.google.com/search?q=site:drive.google.com " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	gdocsButton := widget.NewButton("google_docs", func() {

		url = "https://www.google.com/search?q=site:docs.google.com " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}
	})

	eyedexButton := widget.NewButton("eyedex_search", func() {

		url = "https://eyedex.org/search/?q=" + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	googleOpenDirButton := widget.NewButton("open_dir", func() {

		url = "https://www.google.com/search?q=intitle:\"index of\" -inurl:(jsp|pl|php|html|aspx|htm|cf|shtml) " +
			"-inurl:(hypem|unknownsecret|sirens|writeups|trimediacentral|articlescentral|listen77" +
			"|mp3raid|mp3toss|mp3drug|theindexof|index_of|wallywashis|indexofmp3) " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	trelloButton := widget.NewButton("trello", func() {

		url = "https://www.google.com/search?q=inurl:\"https://trello.com\" " + input.Text
		utils.Cmd.Args = append(utils.CmdArgs, url)

		cmd := utils.Cmd
		err = cmd.Start()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

		time.Sleep(500 * time.Millisecond)
		err = cmd.Process.Kill()
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}

	})

	topContainer := container.NewGridWithColumns(1, input)

	group := container.NewVBox(
		ipfsButton, pastesButton, shortUrlButton, telegramSearchButton, awsBucketsButton, mmntButton, dorkftpButton,
		gdriveButton, gdocsButton, googleOpenDirButton, eyedexButton, trelloButton,
	)

	buttonsContainer := container.NewVScroll(group)

	return container.NewBorder(topContainer, nil, nil, nil, buttonsContainer)

}
