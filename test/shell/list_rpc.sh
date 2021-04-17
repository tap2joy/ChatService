grpcurl -plaintext 127.0.0.1:9101 list ChatService
grpcurl -plaintext 127.0.0.1:9101 describe ChatService.GetChannelList
grpcurl -plaintext 127.0.0.1:9101 describe ChatService.GetChatLog
grpcurl -plaintext 127.0.0.1:9101 describe ChatService.SendMessage