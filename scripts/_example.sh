#!/usr/bin/bash

# example for updating a python project running in a docker container
# the script only prints to stdout if there are changes to pull from the repo

TARGET_PATH=/some/path
cd $TARGET_PATH || echo "ERROR: path not found" || exit

git stash &>/dev/null

# check if there are changes
PULL=$(git pull)
if [[ $PULL != *"Already up to date."* ]]; then
    echo "INFO: git pull"

    #check if docker is running
    DOCKER_PS=$(docker compose ps)

    # build the container
    echo "INFO: docker compose build"
    docker compose build

    # start the container if it was running
    if [[ $DOCKER_PS == *"python ./main.py"* ]]; then
        echo "INFO: docker compose up"
        docker compose up -d
    fi
fi
