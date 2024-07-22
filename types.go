package guolai

// CodeLanguage https://www.wolai.com/wolai/ikDbLea69NRHyTGEsgZwvs
type CodeLanguage string

// BlockAlign https://www.wolai.com/wolai/cxwoKscjb7CZvW3gH8niyr
type BlockAlign string

// BlockBackColors https://www.wolai.com/wolai/fNb4SHWY1bV2s8Xg5JYUE4
type BlockBackColors string

// BlockFrontColors https://www.wolai.com/wolai/o2v1vrLkP2qUuZTH6iDZY9
type BlockFrontColors string

// BlockTypes https://www.wolai.com/wolai/jrmyYiJbEwmF65iQzLFZuK
type BlockTypes string

// HeadingLevel https://www.wolai.com/wolai/4SCaydJc3BwKBAtsW6Fsar
type HeadingLevel int

// InlineTitleType https://www.wolai.com/wolai/wC59T4LoA7zHo1jseXpAWX
type InlineTitleType string

// TextAlign https://www.wolai.com/wolai/wsnMsouhzaSANLwCJcEPYr
type TextAlign string

// TodoListProStatus https://www.wolai.com/wolai/pGch2SkbNjy3cM2Dk8bzxB
type TodoListProStatus string

type Block struct {
	Type            BlockTypes        `json:"type"`
	Content         []RichText        `json:"content"`
	BlockFrontColor *BlockFrontColors `json:"block_front_color"`
	BlockBackColor  *BlockBackColors  `json:"block_back_color"`
	TextAlignment   *TextAlign        `json:"text_alignment"`
	BlockAlignment  *BlockAlign       `json:"block_alignment"`
	// type "heading"
	Level  *HeadingLevel `json:"level"`
	Toggle *bool         `json:"toggle"`
	// type "page"
	Icon        *PageIcon    `json:"icon"`
	PageCover   *LinkCover   `json:"page_cover"`
	PageSetting *PageSetting `json:"page_setting"`
	// type "code"
	Language    *CodeLanguage `json:"language"`
	CodeSetting *CodeSetting  `json:"code_setting"`
	Caption     *string       `json:"caption"`
	// type "quote", no new fields
	// type "callout"
	MarqueeMode *bool `json:"marquee_mode"`
	// type "image" | "video" | "audio"
	Link *string `json:"link"`
	// type "divider", no new fields
	// type "progress_bar"
	Progress   *int  `json:"progress"`
	AutoMode   *bool `json:"auto_mode"`
	HideNumber *bool `json:"hide_number"`
	// type "bookmark", no new fields
	// type "enum_list", no new fields
	// type "todo_list"
	Checked *bool `json:"checked"`
	// type "todo_list_pro"
	TaskStatus *TodoListProStatus `json:"task_status"`
	// type "bull_list", no new fields
	// type "toggle_list", no new fields
	// type "block_equation", no new fields
	// type "embed"
	OriginalLink *string `json:"original_link"`
	EmbedLink    *string `json:"embed_link"`
}

// BlockApiResponse The field `data` in return type of `/blocks/{id}`.
//
// PageID `nil` if the page is top level page in workspace, ParentType will be `work_space`
type BlockApiResponse struct {
	Block
	ID         string     `json:"id"`
	ParentID   string     `json:"parent_id"`
	PageID     *string    `json:"page_id"`
	ParentType BlockTypes `json:"parent_type"`
	Children   struct {
		Ids    []string `json:"ids"`
		APIUrl *string  `json:"api_url"`
	} `json:"children"`
	Version   int    `json:"version"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
	EditedBy  string `json:"edited_by"`
	EditedAt  int64  `json:"edited_at"`
}

// CodeSetting https://www.wolai.com/wolai/uyEE87wHMaSDpNUUSgKvhW
type CodeSetting struct {
	LineNumber    bool   `json:"line_number"`
	LineBreak     bool   `json:"line_break"`
	Ligatures     bool   `json:"ligatures"`
	PreviewFormat string `json:"preview_format"`
}

// LinkCover https://www.wolai.com/wolai/5PCC5XwEgPgBQ9ttTs9eLn
type LinkCover struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// PageIcon https://www.wolai.com/wolai/cWVf9i7LAkAtegbojdfQ2H
//
// The type of link icon and emoji icon is the same in Golang type system.
type PageIcon struct {
	Type string `json:"type"`
	Icon string `json:"icon"`
}

// PageSetting https://www.wolai.com/wolai/fDKkJuUhVrMXb2GgXkAkGA
type PageSetting struct {
	IsFullWidth        bool   `json:"is_full_width"`
	IsSmallText        bool   `json:"is_small_text"`
	HasFloatingCatalog bool   `json:"has_floating_catalog"`
	FontFamily         string `json:"font_family"`
	LineSpacing        string `json:"line_leading"`
}

// RichText https://www.wolai.com/wolai/uPvBQMVskPHHhxzKQBzt2a
type RichText struct {
	Type          InlineTitleType  `json:"type"`
	Title         string           `json:"title"`
	Bold          bool             `json:"bold"`
	Italic        bool             `json:"italic"`
	Underline     bool             `json:"underline"`
	Highlight     bool             `json:"highlight"`
	StrikeThrough bool             `json:"strikethrough"`
	InlineCode    bool             `json:"inline_code"`
	FrontColor    BlockFrontColors `json:"front_color"`
	BackColor     BlockBackColors  `json:"back_color"`
}

type WolaiResponse struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	ErrorCode  int    `json:"error_code"`
	StatusCode int    `json:"status_code"`
}
