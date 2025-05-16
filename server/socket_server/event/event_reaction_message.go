package event

import (
	"forum/constant"
	"forum/models"
	"forum/repository"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
)

func RegisterEventReactionMessage(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSReactionMessage, EventReactionMessage, middleware...)
}

func EventReactionMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csReactionMessage := utils.ConvertMessage[cs.CSReactionMessage](data)
	if csReactionMessage == nil {
		ReactionMessageFailure(client, "Bad request")
		return nil
	}
	room, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csReactionMessage.RoomID, "Users")
	if err != nil {
		ReactionMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if room == nil {
		ReactionMessageFailure(client, "Room chat not found")
		return nil
	}
	inRoom := false
	for _, user_ := range room.Users {
		if user_.ID == client.UserID {
			inRoom = true
			break
		}
	}
	if !inRoom {
		ReactionMessageFailure(client, "Not in room")
		return nil
	}
	message, err := repository.GetRoomMessageRepositoryInstance().FindByIDWithPreloadedField(csReactionMessage.MessageID, "Reactions")
	if err != nil {
		ReactionMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if message == nil {
		ReactionMessageFailure(client, "Room message not found")
		return nil
	}
	if message.Type == constant.MessageTypeNotification {
		ReactionMessageFailure(client, "Can not react to notification")
		return nil
	}
	if !utils.IsReactionTypeValid(csReactionMessage.Type) {
		ReactionMessageFailure(client, "Invalid reaction type")
		return nil
	}
	reaction := message.GetReaction(csReactionMessage.Type)
	if reaction == nil {
		reaction = &models.ContentReaction{
			Type:      csReactionMessage.Type,
			MessageID: &message.ID,
			Count:     0,
		}
		reaction, err = repository.GetReactionRepositoryInstance().Create(reaction)
		if err != nil {
			ReactionMessageFailure(client, "Unexpected error occurred, please try again")
			return nil
		}
		message.Reactions = append(message.Reactions, reaction)
	}
	reaction.Count += 1
	err = repository.GetReactionRepositoryInstance().Update(reaction)
	if err != nil {
		ReactionMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	author, err := repository.GetUserRepositoryInstance().FindByID(*message.UserID)
	if err != nil {
		ReactionMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	BroadcastRoomUpdateMessage(client, room, message, author)
	return nil
}

func ReactionMessageFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCReactionMessage, sc.SCUpdateMessage{Status: sc.StatusError, Message: message})
}
