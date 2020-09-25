# privx-secrets

**The command line application is not supported any more. Please use https://github.com/SSHcom/privx-cli**

The `privx-secrets` is a command-line tool for accessing secrets,
stored in the [PrivX](https://www.ssh.com/products/privx/) secrets
vault.

# Synopsis

privx-secrets [**-api** _endpoint_] [**-config** _file_] [**-v**] _command_ _command arguments_

The `privx-secrets` accepts the following arguments which apply for
all commands:

 - **-api** _endpoint_ sets the PrivX API endpoint to _endpoint_.
 - **-config** _file_ reads the global configuration information from the file _file_.
 - **-v** enable verbose output.

The following commands are defined:

 - **login** verify command authentication information by logging in to the PrivX server.
 - **get** gets secrets from the vault.

## The `login` command

## The `get` command

The get command accepts the following arguments:

  - **-c** Generate C-shell commands on stdout
  - **-s** Generate Bourne shell commands on stdout
  - **-separator** _string_ Data element separator (default ".").
  - **-spread** Spread compounds types

# Configuration file

The default configuration file is `.privx-sdk.toml` at user's home
directory. If the home directory can not be resolved, the program
tries to open the system wide default configuration filed called
`/opt/etc/privx/privx-sdk.toml`. The configuration file location can
be changed with the `-config` command line argument.

## API configuration

The `[api]` configuration block contains the PrivX API configuration
information:

 - `endpoint`: the PrivX API endpoint URL
 - `certificate`: the CA root certificate issuing the REST API server
   certificates. If this is unset, the application will use the
   operating system's trusted CA certificates.

## Authentication configuration

The `[auth]` configuration block contains the API authentication
credentials.

# Examples

Sample secret data for the secret `database`:

    {
      "auth_password": {
        "password": "very secret database password",
        "username": "postgresql"
      },
      "url": "postgresql://proddb.ssh.com:5432"
    }

Get a secret value:

    $ privx-secrets get database.url
    postgresql://proddb.ssh.com:5432

Getting multiple named values:

    $ ./privx-secrets get USERNAME=database.auth_password.username PASSWORD=database.auth_password.password
    USERNAME="postgresql"
    PASSWORD="very secret database password"

The `-c` or `-s` options makes it easy to pull multiple values and
bind them in shell scripts:

    $ eval `privx-secrets get -s USERNAME=database.auth_password.username PASSWORD=database.auth_password.password`
    $ echo $USERNAME
    postgresql
    $ echo $PASSWORD
    very secret database password

Getting configuration blobs with automated value _spreading_:

    $ ./privx-secrets get -spread DB=database
    DB_url="postgresql://proddb.ssh.com:5432"
    DB_auth_password_password="very secret database password"
    DB_auth_password_username="postgresql"
