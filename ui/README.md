# Oh the pain

This is how you compile it, browserify main.js -t babelify --outfile static/bundle.js

From http://babeljs.io/docs/setup/#browserify

Watchify ftw

watchify main.js -t babelify --outfile static/bundle.js
