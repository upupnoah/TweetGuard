# TweetGuard

TweetGuard is a command-line tool for protecting your tweets. It provides various functionalities to manage and secure your Twitter account.

## Installation

To install TweetGuard, you need to have Go installed on your system. Then, you can use the following command to install the tool:

```shell
go get -u github.com/upupnoah/TweetGuard
```

## Usage

Once you have installed TweetGuard, you can use it by running the following command:

```shell
TweetGuard [command]
```

Here, `[command]` represents the specific action you want to perform. The available commands are:

- block: This command blocks users on Twitter.
- block pron: This command is used to ban users who post pornographic content on the timeline.
  
### Examples

Pornographic content users on the block timeline.
To block users on Twitter, you can use the block command. Here is an example:

```shell
TweetGuard block pron
```

## Configuration

Before using TweetGuard, you need to configure your Twitter API credentials. Open the block.go file and modify the following constants with your own credentials:

```go
const (
    consumerKey    = "YOUR_CONSUMER_KEY"
    consumerSecret = "YOUR_CONSUMER_SECRET"
    accessToken    = "YOUR_ACCESS_TOKEN"
    accessSecret   = "YOUR_ACCESS_SECRET"
)
```

Replace `YOUR_CONSUMER_KEY`, `YOUR_CONSUMER_SECRET`, `YOUR_ACCESS_TOKEN`, and `YOUR_ACCESS_SECRET` with your actual Twitter API credentials.

## License

This project is licensed under the [Apache-2.0](https://github.com/upupnoah/TweetGuard/blob/main/LICENSE) license.
