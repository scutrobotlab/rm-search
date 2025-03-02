package index

import "time"

type BbsPostResp RestResp[any]

type RestResp[T any] struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
}

type BbsPost struct {
	Id                int64           `json:"id"`
	Category          string          `json:"category"`
	State             string          `json:"state"`
	StateDesc         string          `json:"stateDesc"`
	AuditRejectReason *string         `json:"auditRejectReason"`
	Official          bool            `json:"official"`
	Top               bool            `json:"top"`
	Marrow            bool            `json:"marrow"`
	Title             string          `json:"title"`
	ContentType       string          `json:"contentType"`
	HtmlContent       string          `json:"htmlContent"`
	MarkdownContent   string          `json:"markdownContent"`
	HeadImg           *string         `json:"headImg"`
	AuthorId          int64           `json:"authorId"`
	AuthorNickname    *string         `json:"authorNickname"`
	AuthorAvatar      string          `json:"authorAvatar"`
	CreateAt          Time            `json:"createAt"`
	UpdateAt          Time            `json:"updateAt"`
	Views             int64           `json:"views"`
	Approvals         int64           `json:"approvals"`
	Comments          int64           `json:"comments"`
	OriginalAuthor    *string         `json:"originalAuthor"`
	OriginalTitle     *string         `json:"originalTitle"`
	OriginalUrl       *string         `json:"originalUrl"`
	Directory         *string         `json:"directory"`
	DifficultyScore   int64           `json:"difficultyScore"`
	Tags              []Tag           `json:"tags"`
	SolutionId        *int64          `json:"solutionId"`
	Solution          *Solution       `json:"solution"`
	BelongWikis       []BelongWiki    `json:"belongWikis"`
	DynamicCategory   interface{}     `json:"dynamicCategory"`
	Draft             interface{}     `json:"draft"`
	JobCategoryId     interface{}     `json:"jobCategoryId"`
	CanComment        int64           `json:"canComment"`
	History           bool            `json:"history"`
	ContactText       *string         `json:"contactText"`
	ContactPic        *string         `json:"contactPic"`
	Attachments       []BbsAttachment `json:"attachments"`
	References        []Reference     `json:"references"`
	BelongTeamContent bool            `json:"belongTeamContent"`
	TeamInfo          *TeamInfo       `json:"teamInfo"`
}

type Tag struct {
	Id        int64       `json:"id"`
	GroupName string      `json:"groupName"`
	Name      string      `json:"name"`
	HeadImg   interface{} `json:"headImg"`
}

type BelongWiki struct {
	WikiId      int64       `json:"wikiId"`
	WikiName    string      `json:"wikiName"`
	WikiHeadImg string      `json:"wikiHeadImg"`
	NodeId      int64       `json:"nodeId"`
	Type        string      `json:"type"`
	Childs      interface{} `json:"childs"`
}

type Solution struct {
	Id         *int64  `json:"id"`
	UserId     *int64  `json:"userId"`
	Content    *string `json:"content"`
	CreateAt   *Time   `json:"createAt"`
	UserName   *string `json:"userName"`
	UserAvatar *string `json:"userAvatar"`
}

type BbsAttachment struct {
	Id        int64  `json:"id,omitempty"`
	Src       string `json:"src"`
	Name      string `json:"name"`
	Downloads int64  `json:"downloads,omitempty"`
}

type Reference struct {
	Id                int64   `json:"id"`
	Url               string  `json:"url"`
	Type              int64   `json:"type"`
	Title             string  `json:"title"`
	Author            *string `json:"author"`
	Season            *string `json:"season"`
	Category          string  `json:"category"`
	CreateAt          Time    `json:"createAt"`
	OrderNum          int64   `json:"orderNum"`
	UpdateAt          *Time   `json:"updateAt"`
	ContentId         string  `json:"contentId"`
	CitedPostId       *int64  `json:"citedPostId"`
	CitingWikiId      int64   `json:"citingWikiId"`
	CitingWikiNodeId  int64   `json:"citingWikiNodeId"`
	CitedPostCreateAt Time    `json:"citedPostCreateAt"`
}

type TeamInfo struct {
	Id          int64   `json:"id"`
	NameEn      *string `json:"nameEn"`
	NameZh      string  `json:"nameZh"`
	CollegeId   *int64  `json:"collegeId"`
	CollegeName string  `json:"collegeName"`
}

type Announce struct {
	Id          int64        `json:"id"`
	Title       string       `json:"title"`
	Date        time.Time    `json:"date"`
	Context     string       `json:"context"`
	Content     string       `json:"content"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Id           int64     `json:"id,omitempty"`
	URL          string    `json:"url"`
	Name         string    `json:"name,omitempty"`
	Size         int64     `json:"size,omitempty"`
	ContentType  string    `json:"contentType,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
	Data         []byte    `json:"data,omitempty"`
}

type ListBbsPostResp RestResp[ListBbsPostData]

type ListBbsPostData struct {
	List  []BbsPostItem `json:"list"`
	Total int64         `json:"total"`
	Size  int64         `json:"size"`
}

type BbsPostItem struct {
	History        bool          `json:"history"`
	Official       bool          `json:"official"`
	Top            bool          `json:"top"`
	Marrow         bool          `json:"marrow"`
	HeadImg        string        `json:"headImg"`
	Id             int           `json:"id"`
	Category       string        `json:"category"`
	CategoryDesc   string        `json:"categoryDesc"`
	Title          string        `json:"title"`
	Introduction   string        `json:"introduction"`
	AuthorId       int           `json:"authorId"`
	AuthorNickname string        `json:"authorNickname"`
	AuthorAvatar   string        `json:"authorAvatar"`
	CreateAt       time.Time     `json:"createAt"`
	Views          int           `json:"views"`
	Approvals      int           `json:"approvals"`
	Comments       int           `json:"comments"`
	Tags           []interface{} `json:"tags"`
	Solution       interface{}   `json:"solution"`
	SolutionDesc   string        `json:"solutionDesc"`
	State          string        `json:"state"`
	StateDesc      string        `json:"stateDesc"`
	UpdateAt       time.Time     `json:"updateAt"`
	WikiId         int           `json:"wikiId"`
}
