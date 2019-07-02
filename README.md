# Dcard-popular-post-notify

> Send dcard popular post to telegram channel: <https://t.me/dcard_popular_post_notify>

## Config example

```yaml
debug: true
lessLikeCount: 3000
dcard:
  postsURL: "https://www.dcard.tw/_api/posts"
  postURL: "https://www.dcard.tw/f/all/p/%d"
telegram:
  token: "token"
  chatID: "-1001129684762"
  disableWebPagePreview: true
  sendMessageURL: "https://api.telegram.org/bot%s/sendMessage"
db:
  migrate: true
  url: "host=localhost port=5432 user=root dbname=dcard_notify password=pwd sslmode=disable"
```
