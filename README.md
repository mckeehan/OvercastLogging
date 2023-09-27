# OvercastLogging

This project was created to fetch a list of played podcasts from [Overcast.fm](https://overcast.fm) using their "All Data" export which provides OPML with playlists and episode data added.

To run this, you will need to create a config.yaml with your account information for Overcast (e-mail and password):

```yaml
email: someone@email-service.com
password: monkey123
```
You can optionally provide a date on the command line for the date to be processed - when no date is provided, the current date will be used.
