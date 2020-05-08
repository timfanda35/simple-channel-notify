# Simple Channel Notity

## Build

Build conainer image

```
docker build -t simple-channel-notify .
```

## Telegram

Send notify message to Telegram

```
docker run \
 -e NOTIFY_TELEGRAM_TOKEN="$NOTIFY_TELEGRAM_TOKEN" \
 -e NOTIFY_TELEGRAM_CHAT_ID="$NOTIFY_TELEGRAM_CHAT_ID" \
 simple-channel-notify telegram --message="Notify Message~~~"
```

References:

- [從零開始的 Telegram Bot](https://blog.sean.taipei/2017/05/telegram-bot)
- [用 Docker Multi-Stage 編譯出 Go 語言最小 Image](https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/)

## Line Notify

Send notify message to Line Notify

```
docker run \
 -e NOTIFY_LINE_NOTIFY_TOKEN="$NOTIFY_LINE_NOTIFY_TOKEN" \
 simple-channel-notify linenotify --message="Notify Message~~~"
```

References:

- [LINE Notify API Document](https://notify-bot.line.me/doc/en/)
- [自建 LINE Notify 訊息通知](https://www.oxxostudio.tw/articles/201806/line-notify.html)

## Hangouts Chat

Send notify message to Hangouts Chat

```
docker run \
 -e NOTIFY_HANGOUTS_CHAT_WEBHOOK="$NOTIFY_HANGOUTS_CHAT_WEBHOOK" \
 simple-channel-notify hangoutschat --message="Notify Message~~~"
```

References:

- [Using incoming webhooks](https://developers.google.com/hangouts/chat/how-tos/webhooks)

## Slack

Send notify message to Slcak Channel

```
docker run \
 -e NOTIFY_SLACK_WEBHOOK="$NOTIFY_SLACK_WEBHOOK" \
 simple-channel-notify slack   --message="Notify Message~~~"
```

References:

- [Incoming Webhooks for Slack](https://slack.com/intl/en-tw/help/articles/115005265063-Incoming-Webhooks-for-Slack)
