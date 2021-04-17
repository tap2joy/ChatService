grpcurl -plaintext -d '{"sender_name": "test1", "channel": 1, "content": "Hello world"}' \
    127.0.0.1:9101 ChatService/SendMessage