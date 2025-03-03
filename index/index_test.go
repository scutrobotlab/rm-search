package index

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scutrobotlab/rm-search/svc"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestIndexer_RecreateIndex(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	index, err := idx.CreateIndex()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %s created", index)

	count, err := idx.ScrollAndIndexBbsPost(ctx, index, 1, 1_000_000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d posts", count)

	count, err = idx.ScrollAndIndexAnnounce(ctx, index, 1, 2000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d announces", count)

	count, err = idx.ScrollAndIndexAttachment(ctx, index, 1, 2000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d attachments", count)

	TestIndexer_DeleteUnusedIndices(t)
}

func TestIndexer_ScrollAndIndex(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	index, err := idx.CreateIndex()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %s created", index)

	count, err := idx.ScrollAndIndexBbsPost(ctx, index, 1, 1_000_000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d posts", count)
}

func TestIndexer_DeleteUnusedIndices(t *testing.T) {
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	if err := idx.DeleteUnusedIndices(); err != nil {
		t.Fatal(err)
	}
	t.Log("unused indices deleted")
}

func TestIndexer_ScrollAndIndexAnnounce(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	index, err := idx.CreateIndex()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %s created", index)

	count, err := idx.ScrollAndIndexAnnounce(ctx, index, 1, 2000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d announces", count)
}

func TestIndexer_ScrollAndIndexAttachment(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	index, err := idx.CreateIndex()
	if err != nil {
		t.Fatal(err)
	}

	count, err := idx.ScrollAndIndexAttachment(ctx, index, 1, 2000)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index %d attachments", count)
}

func TestIndexer_ScrollBbsPost(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithElastic())
	idx := NewIndexer(svcCtx)

	p := idx.SvcCtx.Query.BbsPostItem
	items, err := p.WithContext(ctx).
		Where(p.HeadImg.Neq("[]")).
		Where(p.HeadImg.Neq("null")).
		Find()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("items.length: %d", len(items))

	_ = os.Mkdir("images", 0755)

	for _, item := range items {
		var headImages []map[string]interface{}
		if err := json.Unmarshal([]byte(item.HeadImg), &headImages); err != nil {
			t.Fatal(err)
		}
		if len(headImages) == 0 {
			t.Errorf("item.ID: %d, item.HeadImg: %s", item.ID, item.HeadImg)
		}
		for _, headImage := range headImages {
			resp, err := http.Get(headImage["url"].(string))
			if err != nil {
				t.Errorf(err.Error())
				continue
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf(err.Error())
				continue
			}
			resp.Body.Close()

			err = os.WriteFile(fmt.Sprintf("images/%d_%s_%s", item.ID, item.Title, headImage["alt"].(string)), data, 0644)
			if err != nil {
				t.Errorf(err.Error())
				continue
			}
		}
	}
}
