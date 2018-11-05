#!/bin/bash

heroku container:push web
heroku container:release web
heroku run /bin/app migrate