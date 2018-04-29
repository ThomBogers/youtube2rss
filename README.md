
# Project Config file
Configuration is read from a file name config.json placed in the root of this project
```
{
  "RetreiveUrl": "https://www.youtube.com/feeds/videos.xml?channel_id=example",
  "PublishUrl": "http://casts.example.com/feed/rss.xml",
  "Description": "Rss feed for podcast based on youtube channel",
  "Title": "My custom rss feed",
  "AuthorName": "example",
  "AuthorEmail": "example@example.com"
}
```

# Nginx configuration file
```
server {
       listen         80;
       server_name    casts.example.com;

        root /etc/nginx/www/casts;

        ## auth_basic "Restricted";
        ## auth_basic_user_file /etc/nginx/basic_auth/casts;
}
```

# Todo

- Fix rss for podcast
- Fix nginx config

