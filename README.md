# HTTP Request Logger

This is a simple HTTP server that logs out various info about any HTTP request
made to it. Use for debugging or whatever.

## Deploy to Heroku

You can obviously deploy to whatever you'd like, but Heroku is simple and free for this. Presuming you have a Heroku account and are logged in via the CLI tool:

```
heroku create
heroku push remote main
```

Then, once it's deployed, send any HTTP request to your Heroku endpoint, and finally view the logs:

```
heroku logs
```

