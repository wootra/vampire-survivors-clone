cd ./wasm
#do one time by default
GOOS=js GOARCH=wasm go build -o ../public/result.wasm
LOG_FOLDER="../../logs"
# create logs folder if not exist
if [ ! -d $LOG_FOLDER ]; then
  echo "$LOG_FOLDER does not exist, so make one"
  mkdir $LOG_FOLDER;
fi
echo "collect files"
find . | grep -e .go -e .html -e .js > $LOG_FOLDER/go_files.log
DATELOG_ORG=""
for aFile in $(cat $LOG_FOLDER/go_files.log)
do
  DATELOG_ORG+="$(date -r $aFile)"
done
# echo $DATELOG_ORG > org_dates.log
echo "dev server is running..."
MODIFIED=0
while true;
do
  sleep 3;
  clear;
  DATELOG=""
  for aFile in $(cat $LOG_FOLDER/go_files.log)
  do
    DATELOG+="$(date -r $aFile)"
  done
#   echo $DATELOG > curr_dates.log
  RESULT=true
  
  if [[ $DATELOG != $DATELOG_ORG ]]; then RESULT="DIFF"; else RESULT="SAME"; fi
  echo "is it different? $RESULT"
  if [[ $DATELOG != $DATELOG_ORG ]]; then 
    
    # diff ./curr_dates.log ./org_dates.log
    MODIFIED=$(expr $MODIFIED + 1);
    echo "change is detected! $MODIFIED"; 
    GOOS=js GOARCH=wasm go build -o ../public/result.wasm;
    cp -rf ../web/* ../public;
    DATELOG_ORG=$DATELOG;
    # echo $DATELOG_ORG > org_dates.log
  else
    echo "files are being monitored...(changed: $MODIFIED)"
  fi
done