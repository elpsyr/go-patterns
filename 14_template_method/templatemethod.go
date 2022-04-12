package templatemethod

import "fmt"

type Downloader interface {
	Download(uri string)
}

// 一个结构体用于组织流程
// Download
type template struct {
	implement
	uri string
}

func (t *template) prepare() {
	fmt.Print("prepare downloading\n")
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{implement: impl}
}

func (t *template) Download(uri string) {
	t.uri = uri
	t.prepare()
	//fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	// 子类
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	// 子类
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*FTPDownloader) save() {
	fmt.Printf("http save\n")
}
