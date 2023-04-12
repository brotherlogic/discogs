mkdir -p users/brotherlogic/collection/folders/0/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=1&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=1
