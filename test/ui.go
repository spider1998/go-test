package main

import (
	`bytes`
	`encoding/json`
	`fmt`
	`github.com/mattn/go-gtk/gdkpixbuf`
	`github.com/mattn/go-gtk/glib`
	`github.com/mattn/go-gtk/gtk`
	`github.com/pkg/errors`
	`io/ioutil`
	`net/http`
	`os`
	`os/exec`
	`reflect`
	`regexp`
	`sort`
	`strconv`
	"strings"
	`time`
)

func main() {
	var (
		start, end gtk.TextIter
	)
	textview := gtk.NewTextView()
	swin := gtk.NewScrolledWindow(nil, nil)
	frameParams := gtk.NewFrame("body")
	entryheadersv := gtk.NewEntry()
	labelheader := gtk.NewLabel(":")
	entryheadersk := gtk.NewEntry()
	headers := gtk.NewHBox(false, 1)
	frameheader := gtk.NewFrame("header")
	frameboxheader := gtk.NewVBox(false, 1)
	bt := gtk.NewButtonWithLabel("GO!")
	comboboxentry := gtk.NewComboBoxText()
	combos := gtk.NewHBox(false, 1)
	frame2 := gtk.NewFrame("Response")
	framebox2 := gtk.NewVBox(false, 1)
	frame1 := gtk.NewFrame("Request")
	framebox1 := gtk.NewVBox(false, 1)
	vpaned := gtk.NewVPaned()
	menubar := gtk.NewMenuBar()
	vbox := gtk.NewVBox(false, 1)
	frameboxParams := gtk.NewVBox(false, 1)
	resp := gtk.NewHBox(false, 1)
	basicParams := gtk.NewVBox(false, 1)
	basic := gtk.NewFrame("basic:")
	bs := gtk.NewHBox(false, 1)
	bstime := gtk.NewLabel("times: 0 ms")
	bscode := gtk.NewLabel("status: 200 OK")
	bssize := gtk.NewLabel("size: 0 KB")
	cookies := gtk.NewFrame("cookies:")
	cookieParams := gtk.NewVBox(false, 1)
	swin1 := gtk.NewScrolledWindow(nil, nil)
	textview1 := gtk.NewTextView()
	body := gtk.NewFrame("body:")
	bodyParams := gtk.NewVBox(false, 1)
	swin2 := gtk.NewScrolledWindow(nil, nil)
	textview2 := gtk.NewTextView()
	fontbutton := gtk.NewFontButton()
	cascademenu1 := gtk.NewMenuItemWithMnemonic("_File")
	menuitem1 := gtk.NewMenuItemWithMnemonic("E_xit")
	submenu1 := gtk.NewMenu()
	statusbar := gtk.NewStatusbar()
	submenu2 := gtk.NewMenu()
	submenu3 := gtk.NewMenu()
	cascademenu2 := gtk.NewMenuItemWithMnemonic("_View")
	checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	menuitem2 := gtk.NewMenuItemWithMnemonic("_Font")
	fsd := gtk.NewFontSelectionDialog("Font")
	cascademenu3 := gtk.NewMenuItemWithMnemonic("_Help")
	menuitem3 := gtk.NewMenuItemWithMnemonic("_About")
	dialog := gtk.NewAboutDialog()
	pixbuf, _ := gdkpixbuf.NewPixbufFromFile("mattn-logo.png")
	buffer2 := textview2.GetBuffer()
	entry := gtk.NewEntry()
	buffer := textview.GetBuffer()
	buffer1 := textview1.GetBuffer()



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

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------

	vbox.PackStart(menubar, false, false, 0)
	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------

	frame1.Add(framebox1)
	framebox2.SetSizeRequest(100, 300)
	frame2.Add(framebox2)
	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	//--------------------------------------------------------
	// 请求路径及方法
	//--------------------------------------------------------

	combos.SetBorderWidth(10)
	comboboxentry.AppendText("GET")
	comboboxentry.AppendText("POST")
	comboboxentry.AppendText("PUT")
	comboboxentry.SetActive(1)
	comboboxentry.SetSizeRequest(10, 0)
	comboboxentry.Connect("changed", func() {
		fmt.Println("value:", comboboxentry.GetActiveText())
	})
	combos.Add(comboboxentry)
	entry.SetText("http://127.0.0.1:8080")
	entry.SetSizeRequest(300, 30)
	combos.Add(entry)

	bt.Clicked(func() {
		buffer.GetBounds(&start,&end)
		bo := buffer.GetText(&start,&end,false)
		cookie,bodys,code,size,times := goRequest(comboboxentry.GetActiveText(),entry.GetText(),entryheadersk.GetText(),
			entryheadersv.GetText(),bo)
		buffer1.SetText(cookie)
		buffer2.SetText(bodys)
		bscode.SetText(code)
		bssize.SetText(size)
		bstime.SetText(times + "ms")

	})
	combos.Add(bt)

	framebox1.PackStart(combos, false, false, 0)

	//--------------------------------------------------------
	// 请求头
	//--------------------------------------------------------
	frameheader.Add(frameboxheader)
	framebox1.PackStart(frameheader, false, false, 0)
	headers.SetBorderWidth(10)
	entryheadersk.SetTooltipText("KEY(sep`,`)")
	headers.Add(entryheadersk)
	headers.Add(labelheader)
	entryheadersv.SetTooltipText("Value(sep`,`)")
	headers.Add(entryheadersv)
	frameboxheader.PackStart(headers, false, false, 0)

	//--------------------------------------------------------
	// 请求体
	//--------------------------------------------------------
	frameParams.Add(frameboxParams)
	framebox1.PackStart(frameParams, false, false, 0)

	swin.SetSizeRequest(100, 100)
	swin.SetBorderWidth(10)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "{")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "}")
	swin.Add(textview)
	buffer.Connect("changed", func() {
		buffer.GetBounds(&start,&end)
		//buffer.GetIterAtOffset(&end,100)
		fmt.Println("sssssssssssss")
		fmt.Println(buffer.GetText(&start,&end,false))
		fmt.Println("-----------==============")
	})
	frameboxParams.PackStart(swin, false, false, 0)

	//--------------------------------------------------------
	// 响应
	//--------------------------------------------------------


	basic.SetBorderWidth(5)
	basic.Add(basicParams)
	framebox2.PackStart(basic, false, false, 0)

	bs.SetBorderWidth(5)
	bs.Add(bstime)
	bs.Add(bscode)
	bs.Add(bssize)
	basicParams.PackStart(bs, false, false, 0)

	//bodys:
	resp.SetBorderWidth(5)

	//cookies
	cookies.SetBorderWidth(5)
	cookies.Add(cookieParams)
	swin1.SetSizeRequest(100, 200)
	swin1.SetBorderWidth(10)
	swin1.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin1.SetShadowType(gtk.SHADOW_IN)
	swin1.Add(textview1)
	buffer1.Connect("changed", func() {
		fmt.Println("changed")
	})
	cookieParams.PackStart(swin1, false, false, 0)

	//body
	body.SetBorderWidth(5)
	body.Add(bodyParams)
	swin2.SetSizeRequest(100, 200)
	swin2.SetBorderWidth(10)
	swin2.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin2.SetShadowType(gtk.SHADOW_IN)
	swin2.Add(textview2)
	buffer2.Connect("changed", func() {
		fmt.Println("changed")
	})
	bodyParams.PackStart(swin2, false, false, 0)

	resp.Add(cookies)
	resp.Add(body)
	framebox2.PackStart(resp, false, false, 0)

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	fontbutton.Connect("font-set", func() {
		fmt.Println("title:", fontbutton.GetTitle())
		fmt.Println("fontname:", fontbutton.GetFontName())
		fmt.Println("use_size:", fontbutton.GetUseSize())
		fmt.Println("show_size:", fontbutton.GetShowSize())
	})
	menubar.Append(cascademenu1)
	cascademenu1.SetSubmenu(submenu1)

	menuitem1.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu1.Append(menuitem1)
	menubar.Append(cascademenu2)

	cascademenu2.SetSubmenu(submenu2)
	checkmenuitem.Connect("activate", func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	})
	submenu2.Append(checkmenuitem)


	menuitem2.Connect("activate", func() {
		fsd.SetFontName(fontbutton.GetFontName())
		fsd.Response(func() {
			fmt.Println(fsd.GetFontName())
			fontbutton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	submenu2.Append(menuitem2)

	menubar.Append(cascademenu3)
	cascademenu3.SetSubmenu(submenu3)

	menuitem3.Connect("activate", func() {
		dialog.SetName("API-TEST!")
		dialog.SetProgramName("API-TEST")
		dialog.SetLogo(pixbuf)
		dialog.SetLicense("Test License")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})
	submenu3.Append(menuitem3)

	//--------------------------------------------------------
	// 底部
	//--------------------------------------------------------
	contextId := statusbar.GetContextId("api-test")
	statusbar.Push(contextId, "@Go!")
	vbox.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}




//--------------------------------------------------------
// Request
//--------------------------------------------------------
const (
	TAB    string = "	"
	ENTER  string = "\n"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func goRequest(method,url,headersk,headersv,bodys string)(cookie,body,code,size,times string)  {
	fmt.Println(stringToMap(bodys))
	by,err := json.Marshal(stringToMap(bodys))
	if err != nil{
		body = err.Error()
		return
	}
	rtime := time.Now()
	reqs, err := http.NewRequest(method, url, bytes.NewReader(by))
	if err != nil {
		body = err.Error()
		return
	}

	resp, err := do(reqs,headersk,headersv)
	if err != nil {
		body = err.Error()
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = err.Error()
		return
	}

	code = resp.Status
	for _,co := range resp.Cookies(){
		cookie = co.Name+":"+co.Value+";"+"\n"
	}
	size = strconv.Itoa(int(resp.ContentLength)/1024) + " KB"
	times = strconv.Itoa(int(time.Now().Sub(rtime).Nanoseconds())/1000000)
	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		body = err.Error()
		return
	}
	body = mapFormatString(res,0)
	return
}

func do(req *http.Request,headersk,headersv string) (*http.Response, error) {
	k := strings.Split(headersk,",")
	v := strings.Split(headersv,",")
	if len(k) != len(v){
		return nil,errors.New("header mismatch!")
	}
	for i,key := range k{
		if key != ""{
			req.Header.Set(key, v[i])
		}
	}
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

func mapFormatString(m map[string]interface{},counter int)(mapstr string)  {
	var value string
	mapstr += "{"+ENTER
	for k,v := range m{
		if reflect.TypeOf(v).Kind() == reflect.Map {
			value = mapFormatString(v.(map[string]interface{}),counter+1)
		}
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			value = "["
			for i,s := range v.([]interface{}){
				if i != len(v.([]interface{}))-1 {
					value += s.(string) + ","
					continue
				}
				value += s.(string) +"]"
			}
		}
		if reflect.TypeOf(v).Kind() == reflect.String {
			value = v.(string)
		}
		mapstr += strings.Repeat(TAB,counter+1) + k + ":" + value +ENTER
	}
	mapstr += strings.Repeat(TAB,counter)+"}"
	return mapstr
}

func stringToMap(s string)(m map[string]interface{})  {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "	", "", -1)
	m = make(map[string]interface{})
	if len(s) == 2{
		m = nil
		return
	}
	sl := s[2:len(s)-2]
	sliceM := strings.Split(sl,",\n")
	for _,sli := range sliceM{
		var value interface{}
		sls := strings.Split(sli,":")
		if strings.HasSuffix(sls[1],"[") && strings.HasSuffix(sls[1],"]"){
			value = strings.Split(sls[1][1:len(sls[1])-2],",")
		}else{
			value = sls[1]
		}
		m[sls[0]] = value
	}
	return
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