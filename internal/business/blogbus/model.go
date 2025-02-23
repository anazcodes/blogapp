package blogbus

import "time"

type BlogPost struct {
	ID          uint64
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AddBlogPost struct {
	Title       string
	Description string
	Body        string
}

type UpdateBlogPost struct {
	Title       string
	Description string
	Body        string
}

type ID map[string]uint64

func ToID(id uint64) ID {
	return ID{"id": id}
}

func (i ID) ID() uint64 {
	if i == nil {
		return 0
	}

	return i["id"]
}
