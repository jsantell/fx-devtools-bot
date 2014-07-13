fx-devtools-bot
===============

Tweets for [@FirefoxDevtools](https://twitter.com/firefoxdevtools) on every landed patch.

Twitter account: [@fxdevtoolsbot](https://twitter.com/fxdevtoolsbot)

## Using 

Install dependencies:

```
$ go get github.com/jsantell/go-githubstream
$ go get github.com/ChimeraCoder/anaconda
$ go get github.com/bitly/go-simplejson
```

The following environment variables will need to be set with their obvious values:

```
FX_DEVTOOLS_BOT_GITHUB_TOKEN
FX_DEVTOOLS_BOT_CONSUMER_KEY
FX_DEVTOOLS_BOT_CONSUMER_SECRET
FX_DEVTOOLS_BOT_ACCESS_TOKEN
FX_DEVTOOLS_BOT_ACCESS_TOKEN_SECRET
```

Build and run

```
$ go build . && ./fx-devtools-bot

```

## Extending

To extend this for your own commit bot, modify `FilterCommits` to ignore commits as type `github.RepositoryCommit` that can be done on the first pass of filtering (like ignoring merge commits, etc.). Then instantiate a custom Commit object (just needs to implement `FormatMessage() string` and `IsValid() bool` for filtering that may require calling an external API, like in this case, calling Bugzilla's API to ensure it's in the Developer Tools component. You may also need to change the environment variables set mentioned in `Using`.

## License

MIT License, Copyright (c) 2014 Jordan Santell
