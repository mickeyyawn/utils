# utils

utility library

There are two purposes to this utils library.

1. Put a common wrapper around zerolog that can be used from any go service or
   api.
2. Demonstrate how we would ship those logs to Parseable and consume the
   logs with a SQL interface.

For demo purposes, we are going to install and configure Vector to ship
logs that are stored at /var/logs/ZEROLOG_TEST.log.

1. install vector
   #curl --proto '=https' --tlsv1.2 -sSfL https://sh.vector.dev | bash
   brew tap vectordotdev/brew && brew install vector
   brew update && brew upgrade vector

   # that installs vector in /opt/homebrew/Cellar/Vector

   # in the plist it specifies the config file is here:

   # /opt/homebrew/etc/vector/vector.yaml

   update /opt/homebrew/Cellar/vector/0.37.0/homebrew.mxcl.vector.plist
   to point at the toml file at /Users/mickeyyawn/repos/utils/vector.toml

2. to run it
   brew services start vector
   #vector --config /Users/mickeyyawn/.vector/config/vector.yaml
   # note: vector logs are stored at /opt/homebrew/var/log/vector.log
3. start parseable

   docker run -p 8000:8000 \
   -v /tmp/parseable/data:/parseable/data \
   -v /tmp/parseable/staging:/parseable/staging \
   -e P_FS_DIR=/parseable/data \
   -e P_STAGING_DIR=/parseable/staging \
   parseable/parseable:latest \
   parseable local-store

4. parseable is now running at http://localhost:8000/ with credentials: admin/admin

5. create a log stream destination in parseable

curl --location --request PUT \
'http://localhost:8000/api/v1/logstream/demo' \
--header 'Authorization: Basic YWRtaW46YWRtaW4='

6. start Vector
   brew services start vector --file=/Users/mickeyyawn/repos/utils/vector.toml
