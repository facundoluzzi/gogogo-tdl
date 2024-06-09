package files

const filesFolder = "filesstorage"

const bytesSize = 2

// The first byte indicates if the message is the last one, and the other 4 bytes indicates the filename length
const headerSize = 5

const lastMessageHeaderIndex = 0
const filenameHeaderStartIndex = 1
const filenameHeaderEndIndex = 5

const (
	notLastMessage = 0
	lastMessage    = 1
)
