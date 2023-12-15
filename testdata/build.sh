
mkdir release
curl https://www.discogs.com/release/625928 > release/625928
exit

mkdir -p marketplace/
curl -X POST -H "Content-Type:application/json" -d '{"release_id":27962688,"condition":"Mint (M)","price":100.23}'  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings?token=$1"  | sed "s/$1/token/g" > marketplace/listings_79ab71ffb0be08ef4287999053adc626
exit

mkdir -p marketplace/listings
curl -X POST -H "Content-Type:application/json" -d '{"price": 56.55}'  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/2708115424?token=$1" | sed "s/$1/token/g" > marketplace/listings/2708115424_ccbe516a8e40bc625e089da5019c33ea
exit

mkdir -p marketplace/listings
curl curl -X POST -H "Content-Type:application/json" -d ' {"release_id":1349214,"condition":"Very Good Plus (VG+)","price":0,"status":"Expired"}' --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/2828937565?token=$1" | sed "s/$1/token/g"  > marketplace/listings/2828937565_d600f15743b7dea417c6570448571750
exit

mkdir -p masters/1693557
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/1693557/versions?page=1&per_page=100&sort=released" | sed "s/$1/token/g" > masters/1693557/versions_page=1_per_page=100_sort=released
exit

mkdir releases
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1059056"  | sed "s/$1/token/g" > releases/1059056
exit

mkdir releases
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/939775"  | sed "s/$1/token/g" > releases/939775
exit


mkdir releases
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/372019"  | sed "s/$1/token/g" > releases/372019
exit

mkdir releases
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1929402"  | sed "s/$1/token/g" > releases/1929402
exit

mkdir releases
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/372000"  | sed "s/$1/token/g" > releases/372000
exit

mkdir release
curl https://www.discogs.com/release/28154152 > release/28154152
exit

mkdir release
curl https://www.discogs.com/release/2749755 > release/2749755
exit

mkdir release
curl https://www.discogs.com/release/625928 > release/625928
exit

mkdir -p releases/3139057/rating
curl -X PUT -H "Content-Type:application/json" -d '{"rating": 5}' --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/3139057/rating/BrotherLogic?token=$1" | sed "s/$1/token/g" > releases/3139057/rating/brotherlogic_36359a9186d72b959187df1ff3afb788
exit

mkdir -p users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/fields?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/fields
curl  -X POST -H "Content-Type:applicaion/json" ---user-agent "GoDiscogsTestData" "https://api.discogs.com/users/BrotherLogic/collection/folders/0/releases/1163112/instances/19867414/fields/5?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases/1163112/instances/19867414/fields/5
mkdir -p users/brotherlogic/collection/folders/0/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=1&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?token=$1&page=100&per_page=100" |  sed "s/$1/token/g" > users/brotherlogic/collection/folders/0/releases_page=100
exit

mkdir -p marketplace/orders
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/orders/150295-1254?token=$1"  | sed "s/$1/token/g" > marketplace/orders/150295-1254
exit

mkdir -p users/brotherlogic/wants/
curl -X PUT -H "Content-Type:application/json" -d ''  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants/12778444?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/wants/12778444
exit

mkdir -p users/brotherlogic/
curl  -X DELETE --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/6259627?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/6259627
exit


mkdir -p users/brotherlogic/collection/
curl -X POST -H "Content-Type:application/json" -d '{"name":"TestFolder"}'  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders_410402dd300326d636f240064fdc3373
exit

mkdir -p users/brotherlogic/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?page=1&token=$1"  | sed "s/$1/token/g" > users/brotherlogic/wants_page=1
exit

mkdir -p users/brotherlogic/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/inventory?page=1&token=$1"  | sed "s/$1/token/g" > users/brotherlogic/inventory_page=1
exit

mkdir -p users/brotherlogic/collection/folders/3578980/releases/27915987/instances/
curl -X POST -H "Content-Type:application/json" -d '{"folder_id":242017}'  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/3578980/releases/27915987/instances/1427071368?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/242017/releases/27915987/instances/1427071368_ca8608b169103faf5c4f00bbea8a508c
exit

mkdir -p users/brotherlogic/collection/
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders?token=$1"  | sed "s/$1/token/g" > users/brotherlogic/collection/folders/FILE
exit