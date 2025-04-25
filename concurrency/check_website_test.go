package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://github.com/amitc9der"
}

func TestCheckWebSites(t *testing.T) {
	websites := []string{
		"https://github.com/amitc9der",
		"https://google.com/",
		"https://youtube.com/",
	}

	want := map[string]bool{
		"https://github.com/amitc9der": false,
		"https://google.com/":          true,
		"https://youtube.com/":         true,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func slowStudWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i, _ := range urls {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStudWebsiteChecker, urls)
	}
}
