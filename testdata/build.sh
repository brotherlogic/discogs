mkdir -p users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/fields?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/fields
curl  -X POST -H "Content-Type:applicaion/json" -d '{"value": "Yes"}' --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/BrotherLogic/collection/folders/0/releases/1163112/instances/19867414/fields/5?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/5
exit
mkdir -p users/brotherlogic/collection/folders/0/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=1&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=100&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=100
