package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/indigo-sadland/well_aware/utils"
	"time"
)

// domainReconScreen represents the part of window with set of dorks for target domain.
func domainReconScreen(_ fyne.Window) fyne.CanvasObject {

	var err error
	var url string

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter target domain")

	domainBigDataButton := widget.NewButton("domain_big_data", func() {

		url = "https://domainbigdata.com/" + input.Text
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

	googleSubdomainButton := widget.NewButton("google_dork_subdomains", func() {

		url = "https://google.com/search?q=site:*" + input.Text + " -www"
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

	spyOnWebButton := widget.NewButton("spy_on_web", func() {

		url = "https://spyonweb.com/" + input.Text
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
			"| site:codepad.org | site:jsitor.com | site:codepen.io | site:jsfiddle.net | site:dotnetfiddle.net " +
			"| site:phpfiddle.org | site:ide.geeksforgeeks.org | site:repl.it | site:ideone.com | site:paste.debian.net " +
			"| site:paste.org | site:paste.org.ru | site:codebeautify.org | site:codeshare.io | site:trello.com " +
			"| site:npmjs.com | site:papaly.com | site:0bin.net | site:paste-bin.xyz " + input.Text
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

	directoryListingButton := widget.NewButton("directory_listing", func() {

		url = "https://google.com/search?q=site:" + input.Text + " intitle:index+of /"
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

	configFilesButton := widget.NewButton("config_files", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " ext:xml | ext:conf | ext:cnf | ext:reg " +
			"| ext:inf | ext:rdp | ext:cfg | ext:txt | ext:ora | ext:ini | ext:bak | ext:old"
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

	forumsButton := widget.NewButton("forums_search", func() {

		url = "https://www.google.com/search?q=site:forum.infostart.ru | site:sql.ru | site:stackoverflow.com " + input.Text
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

	databaseFilesButton := widget.NewButton("database_files", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " ext:sql | ext:dbf | ext:mdb"
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

	gitlabButton := widget.NewButton("gitlab", func() {

		url = "https://www.google.com/search?q=inurl:gitlab " + input.Text
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

	wordpressButton := widget.NewButton("wordpress", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " inurl:wp- | inurl:wp-content | inurl:plugins " +
			"| inurl:uploads | inurl:themes | inurl:download"
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

	backupFilesButton := widget.NewButton("backup_files", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " ext:bkf | ext:bkp | ext:bak | ext:old " +
			"| ext:backup | ext:zip | ext:rar | ext:tar.gz"
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

	exposedDocumentsButton := widget.NewButton("exposed_documents", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " ext:doc | ext:docx | ext:odt | ext:pdf " +
			"| ext:rtf | ext:sxw | ext:psw | ext:ppt | ext:pptx | ext:pps | ext:csv | ext:xlsx | ext:xls"
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

	backdoorSearchButton := widget.NewButton("backdoor_search", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + "  inurl:shell | inurl:backdoor | inurl:wso " +
			"| inurl:cmd | shadow | passwd | boot.ini | inurl:backdoor"
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

	phpInfoButton := widget.NewButton("php_info", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " ext:php intitle:phpinfo \"published by the PHP Group\""
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

	bufferoverButton := widget.NewButton("bufferover_subdomains", func() {

		url = "https://dns.bufferover.run/dns?q=." + input.Text
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

	installFilesButton := widget.NewButton("installation_files", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " inurl:readme | inurl:license | inurl:install " +
			"| inurl:setup | inurl:config"
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

	sqlErrorsButton := widget.NewButton("sql_errors", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " intext:\"sql syntax near\" " +
			"| intext:\"syntax error has occurred\" | intext:\"incorrect syntax near\" " +
			"| intext:\"unexpected end of SQL command\" | intext:\"Warning: mysql_connect()\" " +
			"| intext:\"Warning: mysql_query()\" | intext:\"Warning: pg_connect()\""
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

	jiraButton := widget.NewButton("jira", func() {

		url = "https://www.google.com/search?q=inurl:jira " + input.Text
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

	openRedirButton := widget.NewButton("open_redirects", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " inurl:redir | inurl:url | inurl:redirect " +
			"| inurl:return | inurl:src=http | inurl:r=http"
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

	apiEndpointsButton := widget.NewButton("api_endpoints", func() {

		url = "https://google.com/search?q=inurl:\"/api/v1\" | \"/api/v2\"" + " site:" + input.Text
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

	apiWSDLButton := widget.NewButton("wsdl", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " filetype:wsdl | filetype:WSDL | ext:svc " +
			"| inurl:wsdl | Filetype: ?wsdl | inurl:asmx?wsdl | inurl:jws?wsdl | intitle:_vti_bin/sites.asmx?wsdl " +
			"| inurl:_vti_bin/sites.asmx?wsdl"
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

	gistSearchButton := widget.NewButton("gist_search", func() {

		url = "https://gist.github.com/search?q=*." + input.Text
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

	gitDirButton := widget.NewButton(".git_dir", func() {

		url = "https://www.google.com/search?q=inurl:" + "/.git " + input.Text + " -github"
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

	atlasbucketButton := widget.NewButton("atlassian/bitbucket", func() {

		url = "https://www.google.com/search?q=site:atlassian.net | site:bitbucket.org " + input.Text
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

	subSubdomainsButton := widget.NewButton("sub-subdomains_dork", func() {

		url = "https://www.google.com/search?q=site:*.*." + input.Text
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

	digioceanButton := widget.NewButton("digitalocean_spaces", func() {

		url = "https://www.google.com/search?q=site:digitaloceanspaces.com " + input.Text
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

	awsButton := widget.NewButton("aws_search", func() {

		url = "https://www.google.com/search?q=site:.s3.amazonaws.com " + input.Text
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

	multipleSearchButton := widget.NewButton("cse_search", func() {

		url = "https://cse.google.com/cse?cx=002972716746423218710:veac6ui3rio#gsc.tab=0&gsc.q=" + input.Text
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

	vulnerableParamsButton := widget.NewButton("vulnerable_params", func() {

		url = "https://www.google.com/search?q=site:" + input.Text + " inurl:php?id= | inurl:index.php?id= " +
			"| inurl:pageid= | inurl:.php?"
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

	waybackOneButton := widget.NewButton("wayback_machine_1", func() {

		url = "http://wwwb-dedup.us.archive.org:8083/cdx/search?url=" + input.Text +
			"/&matchType=domain&collapse=digest&output=text&fl=original,timestamp&filter=urlkey:.*wp[-].*&limit=1000000&xx="
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

	waybackTwoButton := widget.NewButton("wayback_machine_2", func() {

		url = "https://web.archive.org/cdx/search?url=" + input.Text +
			"/&matchType=domain&collapse=urlkey&output=text&fl=original&filter=urlkey:.*swf&limit=100000"
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
		domainBigDataButton, googleSubdomainButton, bufferoverButton, subSubdomainsButton, spyOnWebButton,
		exposedDocumentsButton, backupFilesButton, gitlabButton, jiraButton, atlasbucketButton, wordpressButton,
		directoryListingButton, databaseFilesButton, installFilesButton, configFilesButton, phpInfoButton,
		backdoorSearchButton, sqlErrorsButton, openRedirButton, gitDirButton, apiEndpointsButton, apiWSDLButton,
		vulnerableParamsButton, awsButton, digioceanButton, multipleSearchButton, waybackOneButton, waybackTwoButton,
		forumsButton, pastesButton, gistSearchButton,
	)

	buttonsContainer := container.NewVScroll(group)

	return container.NewBorder(topContainer, nil, nil, nil, buttonsContainer)
}
