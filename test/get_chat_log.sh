grpcurl -plaintext -d '{"channel": 1}' \
    127.0.0.1:9101 ChatService/GetChatLog