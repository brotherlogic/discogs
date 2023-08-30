mkdir -p marketplace/
curl -X POST -H "Content-Type:application/json" -d '{"release_id":27962688,"condition":"Mint (M)","price":100.23}'  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings?token=$1"  | sed "s/$1/token/g" > marketplace/listings_0910013e6acd173c477260d9cd9ac074
exit

mkdir -p users/brotherlogic/collection/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/FILE
exit

mkdir -p users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/fields?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/fields
curl  -X POST -H "Content-Type:applicaion/json" ---user-agent "GoDiscogsTestData" "https://api.discogs.com/users/BrotherLogic/collection/folders/0/releases/1163112/instances/19867414/fields/5?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/5
mkdir -p users/brotherlogic/collection/folders/0/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=1&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=100&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=100
