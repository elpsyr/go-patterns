package templatemethod

import "testing"

func TestTemplate_Download(t *testing.T) {
	httpDownloader := NewHTTPDownloader()
	httpDownloader.Download("www.example.com/abc.zip")

	ftpDownloader := NewFTPDownloader()
	ftpDownloader.Download("www.example.com/abc.zip")
}
