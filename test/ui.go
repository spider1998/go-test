package main

import (
	`fmt`
	`github.com/mattn/go-gtk/gdkpixbuf`
	`github.com/mattn/go-gtk/glib`
	`github.com/mattn/go-gtk/gtk`
	`os`
	`os/exec`
	`regexp`
	`sort`
	`strings`
)

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("API Test!")
	window.SetIconName("go-api-test")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.NewFrame("Request")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Response")
	framebox2 := gtk.NewVBox(false, 1)
	framebox2.SetSizeRequest(100,300)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)


	//--------------------------------------------------------
	// 请求路径及方法
	//--------------------------------------------------------
	combos := gtk.NewHBox(false, 1)
	combos.SetBorderWidth(10)
	comboboxentry := gtk.NewComboBoxText()
	comboboxentry.AppendText("GET")
	comboboxentry.AppendText("POST")
	comboboxentry.AppendText("PUT")
	comboboxentry.SetActive(1)
	comboboxentry.SetSizeRequest(10,0)
	comboboxentry.Connect("changed", func() {
		fmt.Println("value:", comboboxentry.GetActiveText())
	})
	combos.Add(comboboxentry)
	entry := gtk.NewEntry()
	entry.SetText("http://127.0.0.1:8080")
	entry.SetSizeRequest(300,30)
	combos.Add(entry)

	bt := gtk.NewButtonWithLabel("GO!")
	bt.Clicked(func() {println(comboboxentry.GetActiveText()+entry.GetText())})
	combos.Add(bt)

	framebox1.PackStart(combos, false, false, 0)


	//--------------------------------------------------------
	// 请求头
	//--------------------------------------------------------
	frameheader := gtk.NewFrame("header")
	frameboxheader := gtk.NewVBox(false, 1)
	frameheader.Add(frameboxheader)
	framebox1.PackStart(frameheader, false, false, 0)
	headers := gtk.NewHBox(false, 1)
	headers.SetBorderWidth(10)
	entryheadersk := gtk.NewEntry()
	entryheadersk.SetTooltipText("KEY(sep`,`)")
	headers.Add(entryheadersk)
	labelheader := gtk.NewLabel(":")
	headers.Add(labelheader)
	entryheadersv := gtk.NewEntry()
	entryheadersv.SetTooltipText("Value(sep`,`)")
	headers.Add(entryheadersv)
	frameboxheader.PackStart(headers, false, false, 0)







	//--------------------------------------------------------
	// 请求体
	//--------------------------------------------------------
	frameParams := gtk.NewFrame("body")
	frameboxParams := gtk.NewVBox(false, 1)
	frameParams.Add(frameboxParams)
	framebox1.PackStart(frameParams, false, false, 0)



	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetSizeRequest(100,100)
	swin.SetBorderWidth(10)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	textview := gtk.NewTextView()
	var start, end gtk.TextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "{")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "}")
	swin.Add(textview)
	buffer.Connect("changed", func() {
		fmt.Println("changed")
	})
	frameboxParams.PackStart(swin, false, false, 0)


	//--------------------------------------------------------
	// 响应
	//--------------------------------------------------------

	basic := gtk.NewFrame("basic:")
	basic.SetBorderWidth(5)
	basicParams := gtk.NewVBox(false, 1)
	basic.Add(basicParams)
	framebox2.PackStart(basic, false, false, 0)

	bs := gtk.NewHBox(false,1)
	bs.SetBorderWidth(5)
	bstime := gtk.NewLabel("times: 0 ms")
	sstime := gtk.NewLabel("status: 200 OK")
	sistime := gtk.NewLabel("size: 0 KB")
	bs.Add(bstime)
	bs.Add(sstime)
	bs.Add(sistime)
	basicParams.PackStart(bs, false, false, 0)



	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	fontbutton := gtk.NewFontButton()
	fontbutton.Connect("font-set", func() {
		fmt.Println("title:", fontbutton.GetTitle())
		fmt.Println("fontname:", fontbutton.GetFontName())
		fmt.Println("use_size:", fontbutton.GetUseSize())
		fmt.Println("show_size:", fontbutton.GetShowSize())
	})


	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	var menuitem *gtk.MenuItem
	menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	checkmenuitem.Connect("activate", func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	})
	submenu.Append(checkmenuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("_Font")
	menuitem.Connect("activate", func() {
		fsd := gtk.NewFontSelectionDialog("Font")
		fsd.SetFontName(fontbutton.GetFontName())
		fsd.Response(func() {
			fmt.Println(fsd.GetFontName())
			fontbutton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("_About")
	menuitem.Connect("activate", func() {
		dialog := gtk.NewAboutDialog()
		dialog.SetName("API-TEST!")
		dialog.SetProgramName("API-TEST")
		pixbuf, _ := gdkpixbuf.NewPixbufFromFile("mattn-logo.png")
		dialog.SetLogo(pixbuf)
		dialog.SetLicense("Test License")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})
	submenu.Append(menuitem)


	//--------------------------------------------------------
	// 底部
	//--------------------------------------------------------
	statusbar := gtk.NewStatusbar()
	context_id := statusbar.GetContextId("api-test")
	statusbar.Push(context_id, "@Go!")
	vbox.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}


func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Yasuhiro Matsumoto <mattn.jp@gmail.com>"}
}
