# Dcard-popular-post-notify

> Send dcard popular post to telegram channel: <https://t.me/dcard_popular_post_notify>

## Config example

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
