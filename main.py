import json
import requests

from os import path, system

# os settings
PWD = path.dirname(path.abspath(__file__))
SENT_JSON_FILE_PATH = path.join(PWD, 'sent.json')

# telegram settings
CHAR_ID = '-1001129684762'
DISABLE_WEB_PAGE_PREVIEW = 'true'
with open(path.join(PWD, 'token.txt')) as f:
    SEND_MESSAGE_URL = 'https://api.telegram.org/bot{}/sendMessage'.format(
        f.read().replace('\n', '')
    )

# dcard settings
TARGET_URL = 'https://www.dcard.tw/_api/posts'
POST_URL = 'https://www.dcard.tw/f/all/p/{}'


def get_popular_post_list(min_like_count=3000):
    before = None
    params = {'popular': 'true'}
    while True:
        if before:
            params.update({'before': before})

        posts = requests.get(TARGET_URL, params).json()
        for post in posts:
            if post['likeCount'] < min_like_count:
                break

            yield (post['title'], post['id'])

        if posts[-1]['likeCount'] < min_like_count:
            break

        before = posts[-1]['id']


def get_sent_post_id_list():
    if not path.exists(SENT_JSON_FILE_PATH):
        system('touch {}'.format(SENT_JSON_FILE_PATH))
        return []

    with open(SENT_JSON_FILE_PATH, 'r') as f:
        sent_post_id_list = json.loads(f.read())

    return sent_post_id_list  # post that have already send


def send_popular_post(post_title, post_id):
    params = {
        'chat_id': CHAR_ID,
        'disable_web_page_preview': DISABLE_WEB_PAGE_PREVIEW,
        'text': '{} - {}'.format(post_title, POST_URL.format(post_id))
    }
    requests.get(SEND_MESSAGE_URL, params)


def write_sent_post_id_list(sent_post_id_list):
    with open(SENT_JSON_FILE_PATH, 'w') as f:
        f.write(json.dumps(sent_post_id_list, indent=4))


def main():
    popular_post_list = list(get_popular_post_list())
    sent_post_id_list = get_sent_post_id_list()
    for post_title, post_id in popular_post_list:
        if post_id in sent_post_id_list:
            continue

        send_popular_post(post_title, post_id)
        sent_post_id_list.append(post_id)

    write_sent_post_id_list(sent_post_id_list)


if __name__ == '__main__':
    main()
