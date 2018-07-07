package main

import (
	// "errors"
	"flag"
	// "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

//全局变量，缓存模板
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var addr = flag.Bool("addr", false, "find open address and print to final-port.txt")

//验证url
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/*func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Tile")
	}
	return m[2], nil
}*/

/**
 * 载入页面
 */
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//保存
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//查看
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// 处理不存在的page
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//编辑
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	thisPage, err := loadPage(title)
	if err != nil {
		thisPage = &Page{Title: title}
	}
	renderTemplate(w, "edit", thisPage)
	// tmplt, _ := template.ParseFiles("edit.htm")
	// tmplt.Execute(w, thisPage)
}

//抽出来，负责渲染模板
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
 * 发现view、edit和save 的Handler都上来包含了getTitle方法进行判断，
 * 我们可以统一做他们重复的代码，然后再该谁是谁，需要一个函数来做这个过程，得益于golang的
 * 函数是第一公民的特性，我们可以很简单的实现，因为重复代码中的结果需要用到，所以
 * 给三个handler加一个参数。在java8以前这种方式可以用设计模式实现，新的java也可以了！
 *
 */

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//主函数！
func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}

// ***TODO****
// 文件目录，
// 数据库
// root redirect to view
// inter link
// https://go-zh.org/doc/articles/wiki/

//https://golang.org/doc/effective_go.html#interface_methods
//http.handle(,)
//可以是一个实现了serveHTTP方法的object
//或者是一个普通的function但是signture和serveHTTP一样，使用http.HandlerFunc
//我草，go的这种模式真是很绕啊！！！不知怎么想出来的！理解就很费劲了！
//看来函数式编程需要好好学学了！
//首先是一下两种方式，对于第二种方式，接收一个参数为这俩的函数作为handler，
//func Handle(pattern string, handler Handler)
//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// 然后http包内定义了一个类型叫HandlerFunc，
//type HandlerFunc func(ResponseWriter, *Request)  并实现了Handler的方法，
//表明他是一种沟通的adapter。也就是说第二种方式里面的普通函数作为这种类型不知不觉就实现了
// handler的方法，可以作为handler了！！
//难点有两个：1、函数作为一等公民的思想！。2、设计模式
