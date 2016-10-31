# clinot.es client [![CircleCI](https://circleci.com/gh/clinotes/client.svg?style=svg)](https://circleci.com/gh/clinotes/client)

This is a little side-project to learn `Go` and write a command line application using [Cobra](https://github.com/spf13/cobra) with a remote API endpoint deployed using [Heroku](https://heroku.com).

The [client](https://github.com/clinotes/client) uses the backend at [`api.clinot.es`](https://clinot.es) per default, but you can host the [server](https://github.com/clinotes/server) by yourself, it works fine with Heroku and the default PostgreSQL add-on.

Make sure to set `CLINOTES_API_HOSTNAME` in your environment after setting up a custom endpoint, or just configure your API endpoint within the `~/.clinotes.yaml` configuration file.

## License

Feel free to use the client code, it's released using the [GPLv3 license](https://github.com/clinotes/client/blob/master/LICENSE.md).

## Contributors

- [Sebastian Müller](https://sbstjn.com)

## Install

#### Homebrew on MacOS

```bash
$ > brew tap clinotes/cn
$ > brew install cn
```

#### Binaries

* [cn_latest_linux_amd64.tar.gz](https://dl.clinot.es/latest/cn_latest_linux_amd64.tar.gz)
* [cn_latest_darwin_amd64.zip](https://dl.clinot.es/latest/cn_latest_darwin_amd64.zip)

## Commands

#### General

- `auth` - Authorize client
- `auth request` - Request new authorization token
- `config` - Show client configuration
- `signup` - Create new account
- `signup verify` - Verify created account
- `version` - Show client version

#### Account

- `me` - Show information for your account
- `subscribe` - Subscribe account to paid plan

#### Notes

- `add` - Add note - *not available yet*
- `start` - Start timer - *not available yet*
- `stop` - Stop timer - *not available yet*
- `today` - Show today's notes - *not available yet*
- `yesterday` - Show yesterday's notes - *not available yet*

## Usage

#### Create account

Configure `CLINOTES_API_HOSTNAME` if you do not want to use the default API endpoint `https://api.clinot.es` for your requests.

```bash
$ > cn signup --mail "mail@example.com"
```

You wil receive an email at the provided address with a token to verify your account.

#### Verify account

```bash
$ > cn signup verify --mail "mail@example.com" --token "leOhEHjDJh"
```

Together with the token from the verification request mail you can now verify that you are the owner of the provided email address.

#### Token

You need a valid token to use the [clinot.es](https://clinot.es) command line application. After verifying your account, just request a new token:

```bash
$ > cn auth request --mail "mail@example.com"
```

The token will be delivered right into your mailbox again.

#### Authorization

```bash
$ > cn auth --mail "mail@example.com" --token "ncMqN4VXSN"
```

If the provided token is valid to authorize access to your account, the configuration will be stored in your `~/.clinotes.yaml` file and you are ready to use `cn` for writing notes.
