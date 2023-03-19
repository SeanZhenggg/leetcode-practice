git add :

name=$(git status -s | sed -nE 's/.*([0-9]+.*\.).*/\1/p')
extension=$(git status -s | sed -nE 's/.*\.([a-z]+)"/\1/p')

msg=$(echo $name | cut -c 1-$(($(echo $name | wc -c)-2)))
git cm -m "$msg in $extension"

git push origin master
