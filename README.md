# Simple Channel Notity

## Build

Build conainer image

```
docker build -t simple-channel-notify .
```

## Telegram

Sent Notify to Telegram

```
docker run \
 -e NOTIFY_TELEGRAM_TOKEN="$NOTIFY_TELEGRAM_TOKEN" \
 -e NOTIFY_TELEGRAM_CHAT_ID="$NOTIFY_TELEGRAM_CHAT_ID" \
 simple-channel-notify telegram --message="Notify Message~~~"
```

Reference:

- [從零開始的 Telegram Bot](https://blog.sean.taipei/2017/05/telegram-bot)
- [用 Docker Multi-Stage 編譯出 Go 語言最小 Image](https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/)
