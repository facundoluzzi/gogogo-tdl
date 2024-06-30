package files

const filesFolder = "filesstorage"

const bytesSize = 2

// The first byte indicates if the message is the last one, and the other 4 bytes indicates the filename length
const headerSize = 5

const lastMessageHeaderIndex = 0
const fileNameHeaderStartIndex = 1
const fileNameHeaderEndIndex = 5

const translateServiceURL = "https://e720069868034e7a803e32c2c41d3dab.api.mockbin.io/"

const (
	notLastMessage = 0
	lastMessage    = 1
)

type OperationType string

const (
	NewFile        OperationType = "NEW"
	Read           OperationType = "READ"
	ReadAll        OperationType = "READ_ALL"
	Save           OperationType = "SAVE"
	Find           OperationType = "FIND"
	Delete         OperationType = "DELETE"
	FindAndReplace OperationType = "FIND_AND_REPLACE"
	Append         OperationType = "APPEND"
	DeleteFile     OperationType = "DELETE_FILE"
	Translate      OperationType = "TRANSLATE"
)

func (o OperationType) RequiresExclusiveAccess() bool {
	switch o {
	case Read, ReadAll:
		return false
	default:
		return true
	}
}
