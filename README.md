# Simple Channel Notify

Send simple message via channel's RESTful API.

Support channels:
- Telegram
- Line Notify
- Hangouts Chat
- Slack
- Discord

Pre-build container image on [Releases](https://github.com/timfanda35/simple-channel-notify/pkgs/container/simple-channel-notify)

Repository: `ghcr.io/timfanda35/simple-channel-notify`

## Telegram

Send notify message to Telegram

```
docker run \
  -e NOTIFY_TELEGRAM_TOKEN="$NOTIFY_TELEGRAM_TOKEN" \
  -e NOTIFY_TELEGRAM_CHAT_ID="$NOTIFY_TELEGRAM_CHAT_ID" \
  ghcr.io/timfanda35/simple-channel-notify \
  telegram --message="Notify Message~~~"
```

Environment Variables
- `NOTIFY_TELEGRAM_TOKEN`
- `NOTIFY_TELEGRAM_CHAT_ID`

References:

- [從零開始的 Telegram Bot](https://blog.sean.taipei/2017/05/telegram-bot)
- [用 Docker Multi-Stage 編譯出 Go 語言最小 Image](https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/)

## Line Notify

Send notify message to Line Notify

```
docker run \
  -e NOTIFY_LINE_NOTIFY_TOKEN="$NOTIFY_LINE_NOTIFY_TOKEN" \
  ghcr.io/timfanda35/simple-channel-notify \
  linenotify --message="Notify Message~~~"
```

Environment Variables
- `NOTIFY_LINE_NOTIFY_TOKEN`

References:

- [LINE Notify API Document](https://notify-bot.line.me/doc/en/)
- [自建 LINE Notify 訊息通知](https://www.oxxostudio.tw/articles/201806/line-notify.html)

## Hangouts Chat

Send notify message to Hangouts Chat

```
docker run \
  -e NOTIFY_HANGOUTS_CHAT_WEBHOOK="$NOTIFY_HANGOUTS_CHAT_WEBHOOK" \
  ghcr.io/timfanda35/simple-channel-notify \
  hangoutschat --message="Notify Message~~~"
```

Environment Variables
- NOTIFY_HANGOUTS_CHAT_WEBHOOK

References:

- [Using incoming webhooks](https://developers.google.com/hangouts/chat/how-tos/webhooks)

## Slack

Send notify message to Slack Channel

```
docker run \
  -e NOTIFY_SLACK_WEBHOOK="$NOTIFY_SLACK_WEBHOOK" \
  ghcr.io/timfanda35/simple-channel-notify \
  slack --message="Notify Message~~~"
```

Environment Variables
- `NOTIFY_SLACK_WEBHOOK`

References:

- [Incoming Webhooks for Slack](https://slack.com/intl/en-tw/help/articles/115005265063-Incoming-Webhooks-for-Slack)

## Discord

Send notify message to Discord Text Channel

```
docker run \
  -e NOTIFY_DISCORD_WEBHOOK="$NOTIFY_DISCORD_WEBHOOK" \
  ghcr.io/timfanda35/simple-channel-notify \
  discord --message="Notify Message~~~"
```

Environment Variables
- `NOTIFY_DISCORD_WEBHOOK`

References:

- [Webhook Resource](https://discord.com/developers/docs/resources/webhook#execute-webhook)


## GitHub Action Config Sample

Set the required environment variables in GitLab `Settings -> Secrets -> Acctions` page.

- `NOTIFY_TELEGRAM_TOKEN`
- `NOTIFY_TELEGRAM_CHAT_ID`

```
name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Notify
        run: |
          docker run \
          -e NOTIFY_TELEGRAM_TOKEN="${{ secrets.NOTIFY_TELEGRAM_TOKEN }}" \
          -e NOTIFY_TELEGRAM_CHAT_ID="${{ secrets.NOTIFY_TELEGRAM_CHAT_ID }}" \
          ghcr.io/timfanda35/simple-channel-notify \
          telegram --message="Notify Message~~~"
```

Reference: https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions

## GitLab CI Config Sample

Set the required environment variables in GitLab `Settings -> CI/CD -> Variables` page.

- `NOTIFY_TELEGRAM_TOKEN`
- `NOTIFY_TELEGRAM_CHAT_ID`

`.gitlab-ci.yml`

```
stages:
 - run

notify-telegram:
  stage: run
  image:
    name: 'ghcr.io/timfanda35/simple-channel-notify'
    entrypoint: [""]
  script:
    - /app telegram --message="Message from GitLab CI"
```

Reference:
https://docs.gitlab.com/ee/ci/yaml/gitlab_ci_yaml.html

## Google Cloud Build Config Sample

`cloudbuild.yaml`

```
steps:
  - name: 'ghcr.io/timfanda35/simple-channel-notify'
    args: [ 'telegram', '--message=Message from Cloud Build' ]
    env:
      - 'NOTIFY_TELEGRAM_TOKEN=${_NOTIFY_TELEGRAM_TOKEN}'
      - 'NOTIFY_TELEGRAM_CHAT_ID=${_NOTIFY_TELEGRAM_CHAT_ID}'
```

Reference:
https://cloud.google.com/build/docs/configuring-builds/substitute-variable-values#using_user-defined_substitutions

## Local Build

Build container image

```
docker build -t simple-channel-notify .
```
