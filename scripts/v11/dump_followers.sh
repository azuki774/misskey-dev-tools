#!/bin/bash
USERNAME=${USERNAME}  # misskey username
LOG_DIR=${LOG_DIR}    # log directory
ROOT_DIR=${ROOT_DIR}  # save json directory
ADDRESS=${ADDRESS}    # ex. misskey.dev

LOG_FILE="${LOG_DIR}/dump_misskey.log"
YYYYMM=`date '+%Y%m'`
YYYYMMDD=`date '+%Y%m%d'`

FILE=${ROOT_DIR}/${YYYYMM}/${YYYYMMDD}_followers.json

mkdir -p ${ROOT_DIR}/${YYYYMM}

URL="https://${ADDRESS}/api/users/followers"

echo "`date`: Start sciprts" > ${LOG_FILE}
curl -H "Content-Type: application/json" -d "{\"username\": \"${USERNAME}\" }" ${URL} -o ${FILE}
echo "`date`: Save followers jsonfile: ${FILE}" >> ${LOG_FILE}
echo "`date`: End sciprts" >> ${LOG_FILE}
