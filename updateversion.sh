#!/bin/zsh


ss="github.com/clubpay/ronykit/kit $1"
rs="github.com/clubpay/ronykit/kit $2"

wd=$(pwd)
git tag -a kit/"$2" -m "$2"
git push --tags

array1=(std/gateways/fasthttp std/gateways/fastws std/gateways/silverhttp std/clusters/rediscluster)
for i in "${array1[@]}"
do
	filename="$i"/go.mod
	echo "update go.mod for [$filename]: $ss -> $rs"
	sed -i'' -e 's#'"$ss"'#'"$rs"'#g' "$filename"
	rm "$i"/go.mod-e
	cd "$i" || exit
	go mod tidy
	cd "$wd" || exit
done


# fasthttp
fss="github.com/clubpay/ronykit/std/gateways/fasthttp $1"
frs="github.com/clubpay/ronykit/std/gateways/fasthttp $2"

# contrib
css="github.com/clubpay/ronykit/contrib $1"
crs="github.com/clubpay/ronykit/contrib $2"

array2=(contrib rony)
for i in "${array2[@]}"
do
	filename="$i"/go.mod
	echo "update go.mod for [$filename]: $ss -> $rs"
	sed -i'' -e 's#'"$ss"'#'"$rs"'#g' "$filename"
	echo "update go.mod for [$filename]: $fss -> $frs"
	sed -i'' -e 's#'"$fss"'#'"$frs"'#g' "$filename"
	echo "update go.mod for [$filename]: $css -> $crs"
  sed -i'' -e 's#'"$css"'#'"$crs"'#g' "$filename"
	rm "$i"/go.mod-e
	cd "$i" || exit
	go mod tidy
	cd "$wd" || exit
done

git add .
git commit -m "bump version to $2"
for i in "${array1[@]}"
do
	git tag -a "$i/$2" -m "$2"
done

for i in "${array2[@]}"
do
	git tag -a "$i/$2" -m "$2"
done

git push
git push --tags


sh cleanup.sh
git add .
git commit -m "cleanup go.sum"
git push
