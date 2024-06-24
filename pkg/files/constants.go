package files

const filesFolder = "filesstorage"

const bytesSize = 2

// The first byte indicates if the message is the last one, and the other 4 bytes indicates the filename length
const headerSize = 5

const lastMessageHeaderIndex = 0
const fileNameHeaderStartIndex = 1
const fileNameHeaderEndIndex = 5

const (
	notLastMessage = 0
	lastMessage    = 1
)

type OperationType string

const (
	Read    OperationType = "READ"
	ReadAll OperationType = "READ_ALL"
	Save    OperationType = "SAVE"
	Find    OperationType = "FIND"
)

func (o OperationType) RequiresExclusiveAccess() bool {
	switch o {
	case Read, ReadAll:
		return false
	default:
		return true
	}
}
