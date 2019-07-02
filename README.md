# Dcard-popular-post-notify

> Send dcard popular post to telegram channel: <https://t.me/dcard_popular_post_notify>

## Config example

### With yaml file

```yaml
debug: true
less_like_count: 3000
dcard:
  posts_url: "https://www.dcard.tw/_api/posts"
  post_url: "https://www.dcard.tw/f/all/p/%d"
telegram:
  token: "token"
  chat_id: "-1001129684762"
  disable_web_page_preview: true
  send_message_url: "https://api.telegram.org/bot%s/sendMessage"
db:
  migrate: true
  url: "host=localhost port=5432 user=root dbname=dcard_notify password=pwd sslmode=disable"
```

### With env

```env
DC_DEBUG=true
DC_LESS_LIKE_COUNT=3000
DC_DCARD_POSTS_URL=https://www.dcard.tw/_api/posts
DC_DCARD_POST_URL=https://www.dcard.tw/f/all/p/%d
DC_TELEGRAM_TOKEN=token
DC_TELEGRAM_CHAT_ID=-1001129684762
DC_TELEGRAM_DISABLE_WEB_PAGE_PREVIEW=true
DC_TELEGRAM_SEND_MESSAGE_URL=https://api.telegram.org/bot%s/sendMessage
DC_DB_MIGRATE=true
DC_DB_URL=host=localhost port=5432 user=root dbname=dcard_notify password=pwd sslmode=disable
```
