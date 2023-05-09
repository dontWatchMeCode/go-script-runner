#!/usr/bin/env bash
set -euo pipefail

DATE=$(date '+%Y-%m-%d %H:%M:%S')
PWD=$(pwd)
LOG="$PWD/scripts.log"

# $1 = length
# >> $LOG
function split() {
    for ((i = 0; i < $1; i++)); do
        echo -n "-" >>"$LOG"
    done
}

# $1 = header
# $2 = content
# >> $LOG
function save() {
    split "${#1}"
    printf "\n%s\n" "$1" >>"$LOG"
    split "${#1}"
    printf "\n%s\n\n" "$2" >>"$LOG"
}

if ! cd "$PWD/scripts"; then
    save "[ ERROR: scripts folder not found / $DATE ]" "Please create a script in: $PWD/scripts"
    exit 1
fi

for file in *.sh; do
    [[ $file == _* ]] && continue

    SCRIPT_PATH="$PWD/$file"
    SCRIPT_OUT=$(bash -e "$SCRIPT_PATH")

    if [[ $SCRIPT_OUT != "" ]]; then
        save "[ $file / $DATE ]" "$SCRIPT_OUT"
    fi
done
