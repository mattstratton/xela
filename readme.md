# xela

A webapp for tracking sponsorship, speaking, and cfps for events.

Built with :heart: by [@mattstratton](https://github.com/mattstratton) in Go.

![xela](https://raw.githubusercontent.com/mattstratton/xela/master/assets/images/xela-logo.png)


This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. We appreciate your contribution. Please refer to the [contributing guidelines](CONTRIBUTING.md) for details on how to help.

[Powered by Buffalo](http://gobuffalo.io)

## Setup

Right now, this is pretty sparse. A few things:

### Local Dev

Make sure you have [Buffalo](http://gobuffalo.io) installed.

### Local database

You'll need access to a Postgres database set up in the same definition as listed in `database.yml`. I recommend just using docker. Run this command to get the database going (note; you'll lose all the data when Docker goes down)

```
docker run --rm -it --publish 0.0.0.0:5432:5432 --name pg -e POSTGRES_PASSWORD=postgres postgres:alpine\n\n
```

### Environment variables

Create a file called `.env` at the root of this project (don't worry, it won't get into git). It should look like this:

```
GOOGLE_KEY=xxxxxx-xxxxxx.apps.googleusercontent.com
GOOGLE_SECRET=xxxx
AUTHORIZED_LOGIN_DOMAIN=pagerduty.com
S3_REGION=us-east-1
S3_BUCKET=xela-dev
AWS_ACCESS_KEY_ID=xxxxxx
AWS_SECRET_ACCESS_KEY=xxxxx
```

Replace the values as appropriate. If you don't know the Google settings, check with @mattstratton (if you're a PagerDuty employee; if you're not, you're on your own right now!)

### Heroku setup

The following environment variables must be set in Heroku; should look something like this:

```
=== pure-taiga-48603 Config Vars
AUTHORIZED_LOGIN_DOMAIN: pagerduty.com
DATABASE_URL:            postgres://xxxxxx:xxxxxx@ec2-54-83-49-109.compute-1.amazonaws.com:5432/xxxxx
GOOGLE_KEY:              xxxxxxxx-xxxxxxxxx.apps.googleusercontent.com
GOOGLE_SECRET:           xxxxxxxxxxx
GO_ENV:                  production
HOST:                    https://xxxx.herokuapp.com
SESSION_SECRET:          xxxxxxxxxxxxxxxxxx
S3_REGION:               us-east-1
S3_BUCKET:               xela-prod
AWS_ACCESS_KEY_ID:       xxxxxx
AWS_SECRET_ACCESS_KEY:   xxxxxxx
```

Deploying to Heroku via Docker uses these commands (TODO)

#### Initial Setup

(you'll need the [Heroku plugin](https://github.com/gobuffalo/buffalo-heroku) for Buffalo)
```
buffalo heroku new
heroku container:push web
heroku container:release web
heroku config:set HOST=xxxx.herokuapp.com
heroku config:set AUTHORIZED_LOGIN_DOMAIN=pagerduty.com
heroku config:set GOOGLE_KEY=xxx-xxx.apps.googleusercontent.com
heroku config:set GOOGLE_SECRET=xxx
heroku config:set S3_REGION=xxx
heroku config:set S3_BUCKET=xxx
heroku config:set AWS_ACCESS_KEY_ID=xxx
heroku config:set AWS_SECRET_ACCESS_KEY=xxx
heroku run /bin/app migrate
heroku open
```

#### Deployments

After initial setup, deployments can be run via:

```
heroku container:push web
heroku container:release web
heroku run /bin/app migrate
```

Or you can run the `heroku-deploy.sh` shell script.

## Authors

- **Matt Stratton** - *Initial work* - [mattstratton](https://github.com/mattstratton)

## License

xela - A webapp for tracking sponsorship, speaking, and cfps for events

|                      |                                          |
|:---------------------|:-----------------------------------------|
| **Author:**          | Matt Stratton (<matt.stratton@gmail.com>)
| **Copyright:**       | Copyright 2018, PagerDuty
| **License:**         | The MIT License

```markdown
The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

```