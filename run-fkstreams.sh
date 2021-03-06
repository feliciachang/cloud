#!/bin/bash

pushd server/tools/fkstreams
go build -o fkstreams *.go

if [ -d ~/conservify/dev-ops ]; then

   pushd ~/conservify/dev-ops/provisioning
   if [ -z "$APP_SERVER_ADDRESS" ]; then
      source ./setup-env.sh
   fi
   popd

   if [ -z "$APP_SERVER_ADDRESS" ]; then
      echo No cloud configuration.
      exit 2
   fi

   eval $(ssh-agent)
   trap "ssh-agent -k" exit
   ssh-add ~/.ssh/cfy.pem

   rsync -zvua --progress fkstreams core@$APP_SERVER_ADDRESS:
   ssh core@$APP_SERVER_ADDRESS "FIELDKIT_POSTGRES_URL=$ENV_DB_URL ./fkstreams --host 127.0.0.1 $@"
   rsync -zvua --progress "core@$APP_SERVER_ADDRESS:*.fkpb" ../../../
else
   echo "No dev-ops directory found"
fi

popd
