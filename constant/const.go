package constant

const UserIDContextKey = "UserID"
const RequestIDContextKey = "RequestID"
const AuthorizationTokenContextKey = "AuthorizationToken"

const AuthorizationHeader = "Authorization"
const AuthorizationHeaderPrefix = "Bearer"

const SecretToken = "ledeptraivailzz"
const ExpiredTimeInHour = 24
const ExpiredTimeInHourRemember = 24 * 7

const FileFormKey = "files"
const MaxFileSize = 10 << 20 //10MB
var AllowedFileTypes = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
}

var AllowedReactionTypes = map[string]bool{
	"like":  true,
	"love":  true,
	"haha":  true,
	"wow":   true,
	"sad":   true,
	"angry": true,
}

const MessageTypeText = "text"
const MessageTypeNotification = "notification"

const RoomTypePrivate = "private"
const RoomTypeGroup = "group"

const FriendStatusSendAddFriendRequest = -1
const FriendStatusReceiveFriendRequest = -2
const FriendStatusStranger = 0
const FriendStatusAddedFriend = 1
