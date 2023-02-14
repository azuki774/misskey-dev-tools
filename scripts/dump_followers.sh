#!/bin/bash
MISSKEY_USERNAME=${MISSKEY_USERNAME} # misskey.dev username
LOG_DIR=${LOG_DIR}                   # log directory
ROOT_DIR=${ROOT_DIR}                 # save json directory

LOG_FILE="${LOG_DIR}/dump_followers.log"
YYYYMM=`date '+%Y%m'`
YYYYMMDD=`date '+%Y%m%d'`

FILE=${ROOT_DIR}/${YYYYMM}/${YYYYMMDD}_followers.json

mkdir -p ${ROOT_DIR}/${YYYYMM}

echo `date`: Start sciprts > ${LOG_FILE}
curl -H "Content-Type: application/json" -d "{\"username\": \"${MISSKEY_USERNAME}\" }" https://misskey.dev/api/users/followers -o ${FILE}
echo "`date`: Save followers jsonfile: ${FILE}" > ${LOG_FILE}
echo `date`: End sciprts > ${LOG_FILE}
