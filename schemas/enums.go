package schemas

type Frequency string

const (
	Daily          Frequency = "daily"
	MondayToFriday Frequency = "mondayToFriday"
	Weekly         Frequency = "weekly"
	Monthly        Frequency = "monthly"
	Yearly         Frequency = "yearly"
	Unique         Frequency = "unique"
)

type Priority string

const (
	Important Priority = "important"
	Essential Priority = "essential"
	Normal    Priority = "normal"
)
