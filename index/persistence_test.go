package index

import (
	"context"
	"github.com/scutrobotlab/rm-search/svc"
	"testing"
)

func TestIndexer_Persistence(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.Persistence(ctx, 54068)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_BatchPersistenceRangeIfExist(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.BatchPersistenceRangeIfExist(ctx, 20_0000, 50_0000, 100)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_BatchPersistenceRangeIfNotExist(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.BatchPersistenceRangeIfNotExist(ctx, 1, 10000, 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_PersistenceAnnounce(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.PersistenceAnnounce(ctx, 1768)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_BatchPersistenceAnnounce(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.BatchPersistenceAnnounceRange(ctx, 800, 2000, 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_PersistenceAttachment(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithTika())
	idx := NewIndexer(svcCtx)
	const URL = "https://terra-1-g.djicdn.com/b2a076471c6c4b72b574a977334d3e05/RoboMaster%202025%20%E6%9C%BA%E7%94%B2%E5%A4%A7%E5%B8%88%E8%B6%85%E7%BA%A7%E5%AF%B9%E6%8A%97%E8%B5%9B%E5%8F%82%E8%B5%9B%E6%89%8B%E5%86%8CV1.1.0%EF%BC%8820241225%EF%BC%89.pdf"
	err := idx.PersistenceAttachment(ctx, URL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_BatchPersistenceAttachmentFromAnnounce(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb(), svc.WithTika())
	idx := NewIndexer(svcCtx)
	err := idx.BatchPersistenceAttachmentFromAnnounce(ctx, 800, 2000, 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_PersistenceBbsPostItemRange(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.PersistenceBbsPostItemRange(ctx, 1, 40000)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIndexer_PersistenceBbsPostItemPage(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewContextForTest(svc.WithDb())
	idx := NewIndexer(svcCtx)
	err := idx.PersistenceBbsPostItemPage(ctx, 1, 100)
	if err != nil {
		t.Fatal(err)
	}
}
