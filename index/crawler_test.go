package index

import (
	"fmt"
	"testing"
)

func TestGetBbsPost(t *testing.T) {
	post, err := GetBbsPost(54068)
	if err != nil {
		t.Fatalf("GetBbsPost error: %v", err)
	}
	t.Logf("BbsPost: %+v", post)
}

func TestGetBbsPost2(t *testing.T) {
	const startId = 54050
	const endId = 54100
	for i := startId; i < endId; i++ {
		post, err := GetBbsPost(int64(i))
		if err != nil {
			t.Fatalf("GetBbsPost error: %v", err)
		}
		t.Logf("[%d] BbsPost: %+v", i, post)
		if post.Data != nil {
			t.Logf("Data: %+v", post.Data)
		}
	}
}

func TestGetAnnounce(t *testing.T) {
	announce, err := GetAnnounce(1784)
	if err != nil {
		t.Fatalf("GetAnnounce error: %v", err)
	}
	t.Logf("Announce: %+v", announce)
}

func TestGetAttachment(t *testing.T) {
	const URL = "https://terra-1-g.djicdn.com/b2a076471c6c4b72b574a977334d3e05/RoboMaster%202025%20%E6%9C%BA%E7%94%B2%E5%A4%A7%E5%B8%88%E8%B6%85%E7%BA%A7%E5%AF%B9%E6%8A%97%E8%B5%9B%E5%8F%82%E8%B5%9B%E6%89%8B%E5%86%8CV1.1.0%EF%BC%8820241225%EF%BC%89.pdf"
	attachment, err := GetAttachment(URL)
	if err != nil {
		t.Fatalf("GetAttachment error: %v", err)
	}
	t.Logf("Attachment: %+v", attachment)
}

func TestListBbsPost(t *testing.T) {
	posts, err := ListBbsPost(1, 10)
	if err != nil {
		t.Fatalf("ListBbsPost error: %v", err)
	}
	t.Logf("Posts: %+v", posts)
	if posts.Data != nil {
		for _, item := range posts.Data.List {
			fmt.Printf("%+v\n", item)
		}
	}
}
