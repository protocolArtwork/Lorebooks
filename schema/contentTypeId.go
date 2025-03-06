package schema

type ContentTypeId uint8 // one byte!

const (
	CONT_STORY ContentTypeId = iota
	CONT_CHARACTER
	CONT_LOCATIONTYPEID
	CONT_FACTION
	CONT_ITEM
)
