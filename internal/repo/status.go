package repo

type TagStatus int32

const (
	TagStatusDraft     TagStatus = 0 // Черновик
	TagStatusPublished TagStatus = 1 // Опубликовано
	TagStatusDeleted   TagStatus = 2 // Удален
)
