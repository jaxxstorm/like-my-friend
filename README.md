# Like My Friend

like-my-friend is a tool to look at the latest instagram feed for a user, and like their photos automatically without having to enable post notifications etc. It is inspired by [this](https://github.com/gulzar1996/auto-like-my-gf-insta-pic) javascript tool and [this](https://github.com/cyandterry/Like-My-GF) python tool, but written in Go.

# Usage

Run the binary with a few command line options. The help output is below:

```
If your friend likes you to be quick to like their photos,
this app will poll the latest feed, determine if you've liked their feed recently
and like the photo if you haven't

Usage:
  like-my-friend [flags]

Flags:
  -a, --account string    The instagram account to watch
      --config string     config file (default is $HOME/.like-my-friend.yaml)
  -d, --dryrun            Show what photos you'd like
  -p, --password string   The password for your account
  -P, --posts int         Number of posts to like (default 10)
  -u, --username string   The username for your account
```

## Environment Variables

Because the app uses [viper](https://github.com/spf13/viper) & [cobra](https://github.com/spf13/cobra), you can use environment variables easily. Simple set them using the prefix `INSTA_`:

```
export INSTA_USERNAME="jaxxstorm"
export INSTA_PASSWORD="my-password"
export INSTA_ACCOUNT="my-friend"
like-my-friend
```

# Building

If you want to contribute, we use glide for dependency management, so it should be as simple as:

cloning this repo into `$GOPATH/src/github.com/jaxxstorm/like-my-friend`
run `glide install` from the directory
run `go build -o hookpick main.go`

