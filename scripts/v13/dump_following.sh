#!/bin/bash
USERID=${USERID}      # misskey userid
LOG_DIR=${LOG_DIR}    # log directory
ROOT_DIR=${ROOT_DIR}  # save json directory
ADDRESS=${ADDRESS}    # ex. misskey.dev

LOG_FILE="${LOG_DIR}/dump_misskey.log"
YYYYMM=`date '+%Y%m'`
YYYYMMDD=`date '+%Y%m%d'`

FILE=${ROOT_DIR}/${YYYYMM}/${YYYYMMDD}_following.json

mkdir -p ${ROOT_DIR}/${YYYYMM}

URL="https://${ADDRESS}/api/users/following"

echo "`date`: Start sciprts" > ${LOG_FILE}
curl -H "Content-Type: application/json" -d "{\"userId\": \"${USERID}\" }" ${URL} -o ${FILE}
echo "`date`: Save following jsonfile: ${FILE}" >> ${LOG_FILE}
echo "`date`: End sciprts" >> ${LOG_FILE}
