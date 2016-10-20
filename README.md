# CLInotes client [![CircleCI](https://img.shields.io/circleci/project/clinotes/client.svg)]() [![GitHub release](https://img.shields.io/github/release/clinotes/client.svg)]() [![license](https://img.shields.io/github/license/clinotes/client.svg)]()

Use the [CLInotes](https://clinot.es) client application `cn` to create notes from the command line. Per default the command line interface uses the API hosted at [`https://api.clinot.es`](https://api.clinot.es) but you can configure a custom API endpoint within the `~/.clinotes.yaml` configuration file.

See the [CLInotes server component](https://github.com/clinotes/server) for more information about hosting the API endpoint at your own infrastructure. It works fine at Heroku for example! If you don't want to run your own endpoint, just register for a [free CLInotes account](https://clinot.es) from your command line.

## Instalaltion

#### Homebrew on MacOS

```bash
$ > brew tap clinotes/cn git@github.com:clinotes/homebrew-cn.git
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

#### Notes

- `add` - Add note - *not available yet*
- `start` - Start timer - *not available yet*
- `stop` - Stop timer - *not available yet*
- `today` - Show today's notes - *not available yet*
- `yesterday` - Show yesterday's notes - *not available yet*

## Usage

#### Create account

```bash
$ > cn signup --mail "mail@example.com"
```

You wil receive an email at the provided address with a token to verify you email address.

#### Verify account

```bash
$ > cn signup verify --mail "mail@example.com" --token "leOhEHjDJh"
```

Together with the token from the account verification request mail you can now verify that you are the owner of the provided email address.

#### Authorization token

You need a valid token to use the [CLInotes](https://clinot.es) command line application. After verifying your account, just request a new authorization token:

```bash
$ > cn auth request --mail "mail@example.com"
```

The token will then be delivered right into your mailbox.

#### Authorization

```bash
$ > cn auth --mail "mail@example.com" --token "ncMqN4VXSN"
```

If the provided token is valid to authorize access to your account, the configuration will be stored in your `~/.clinotes.yaml` file and you are ready to use `cn` for writing notes.

# License

The [CLInotes](https://clinot.es) client is available under [MIT License](https://github.com/clinotes/client/blob/master/LICENSE.md).

# Contributors

- [Sebastian Müller](https://sbstjn.com) **//** [GitHub](https://github.com/sbstjn) - [Twitter](https://twitter.com/sbstjn)
