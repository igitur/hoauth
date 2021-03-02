# HOAuth

> Get OpenId Connect tokens from the command line

![A demo of HOAuth in a terminal window](docs/demo.gif)

HOAuth provides a simple way to interact with OpenId Connect identity providers from your local CLI. Many OIDC providers only support the Authorisation Code grant - and that means running a local web server to receive the authorisation response, or using something like [Postman](https://www.postman.com/). These can be tricky to fit into a scripted workflow in a shell.

This tool saves you time, by:
* Helping you configure clients and manage scopes
* Storing client secrets [securely in your OS keychain](https://medium.com/@calavera/stop-saving-credential-tokens-in-text-files-65e840a237bb)
* Managing a local web server to receive the OpenId Connect callback
* Opening a browser to allow users to grant consent
* Using [metadata discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) to build the Authorisation Request
* Verifying the token integrity with the providers's [JWKS](https://tools.ietf.org/html/draft-ietf-jose-json-web-key-41) public keys
* Piping the `access_token`, `id_token` and `refresh_token` to `stdout`, so you can use them in a script workflow

### Supported grant types
* [Authorisation code](https://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth)
* [PKCE](https://tools.ietf.org/html/rfc7636)

## Installation
Download the binary for your platform:
* [Linux](https://github.com/igitur/hoauth/releases/download/v1.1.0/hoauth_1.1.0_linux_amd64.tar.gz)
* [Mac OS](https://github.com/igitur/hoauth/releases/download/v1.1.0/hoauth_1.1.0_darwin_amd64.tar.gz)
* [Windows](https://github.com/igitur/hoauth/releases/download/v1.1.0/hoauth_1.1.0_windows_amd64.tar.gz)

You can run the binary directly:
```sh
./hoauth
```

Or add it to your OS `PATH`:

### Mac/Linux
```sh
mv hoauth /usr/local/bin/hoauth && chmod +x /usr/local/bin/hoauth
```

Alternatively you can use `brew` on Mac OS:

```
brew tap xeroapi/homebrew-taps
brew install hoauth
```

### Windows

The easiest way to get started on Windows is to use [scoop](https://scoop.sh/) to install hoauth:

```sh
scoop bucket add xeroapi https://github.com/XeroAPI/scoop-bucket.git
scoop install hoauth
```

## Quick start

### Prerequisites
* An OpenId Connect Client Id and Secret
* A `redirect_url` of `http://localhost:8080/callback` configured in your OpenId Connect provider's settings (_you can change the port if the default doesn't suit_).

Once the tool is installed, and you have configured your client with the OpenId Provider, run these two commands to receive an access token on your command line:

```shell script
hoauth setup [clientName]
hoauth connect [clientName]
```

## Command reference

### Setup

Creates a new connection

```shell script
hoauth setup [clientName]
# for instance
hoauth setup hike
```

This will guide you through setting up a new client configuration.


#### add-scope

Adds a scope to an existing client configuration

```shell script
hoauth setup add-scope [clientName] [scopeName...]
# for instance
hoauth setup add-scope hike accounting.transactions.read files.read
```

#### remove-scope

Removes a scope from a client configuration

```shell script
hoauth setup remove-scope [clientName] [scopeName...]
# for instance
hoauth setup remove-scope hike accounting.transactions.read files.read
```

#### update-secret

Replaces the client secret, which is stored in your OS keychain

```shell script
hoauth setup update-secret [clientName] [secret]
# for instance
hoauth setup update-secret hike itsasecret!
```

### List

Lists all the connections you have created

```shell script
hoauth list
```

##### Flags

`--secrets`, `-s` - Includes the client secrets in the output (disabled by default)

```shell script
hoauth list --secrets
```


### Delete

Deletes a given client configuration (with a prompt to confirm, we're not barbarians)

```shell script
hoauth delete [clientName]
```

### Connect

Starts the authorisation flow for a given client configuration

```shell script
hoauth connect [clientName]
# for instance
hoauth connect hike
```

##### Flags

`--port`, `-p` - Change the localhost port that is used for the redirect URL

```shell script
# for instance
hoauth connect hike --port 8080
```

`--dry-run`, `-d` - Output the Authorisation Request URL, without opening a browser window or listening for the callback

```shell script
# for instance
hoauth connect hike --dry-run
```

### Token

Output the last set of tokens that were retrieved by the `connect` command

```shell script
hoauth token [clientName]
```

##### Flags

`--refresh`, `-r' - Force a refresh of the access token
```shell script
# for instance
hoauth token hike --refresh
```

`--env`, `-e` - Export the tokens to the environment. By convention, these will be exported in an uppercase format.

```shell script
[CLIENT]_ACCESS_TOKEN
[CLIENT]_ID_TOKEN
[CLIENT]_REFRESH_TOKEN
```

```shell script
# for instance
eval "$(hoauth token hike --env)"
echo $HIKE_ACCESS_TOKEN
```

## Global configuration

### Changing the default web server port

You can modify the default web server port by setting the `HOAUTH_PORT` environment variable:

```shell script
# for instance
HOAUTH_PORT=9999 hoauth setup
```

## Troubleshooting

Run the doctor command to check for common problems:

```shell script
hoauth doctor
```

hoauth stores client configuration in a JSON file at the following location:

```shell script
$HOME/.hoauth/hoauth.json
```

You may want to delete this file if problems persist.

#### Entries in the OS Keychain
Client secrets are saved as application passwords under the common name `igitur.hoauth`


## Contributing

* PRs welcome
* Be kind
