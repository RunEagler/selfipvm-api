package error_message


type errorMessageActivity int

const (
	ActivityInsert errorMessageActivity = iota
)


func (em errorMessageActivity) String() string{

	switch em{
	case ActivityInsert:
		return "Activity insert error."
	}
	return ""
}