# $GOOS	$GOARCH
# android	arm
# darwin	386
# darwin	amd64
# darwin	arm
# darwin	arm64
# dragonfly	amd64
# freebsd	386
# freebsd	amd64
# freebsd	arm
# linux	386
# linux	amd64
# linux	arm
# linux	arm64
# linux	ppc64
# linux	ppc64le
# linux	mips64
# linux	mips64le
# netbsd	386
# netbsd	amd64
# netbsd	arm
# openbsd	386
# openbsd	amd64
# openbsd	arm
# plan9	386
# plan9	amd64
# solaris	amd64
# windows	386
# windows	amd64

# env GOOS=linux GOARCH=arm go build -v .
# env GOOS=darwin GOARCH=amd64 go build -v .

#  https://github.com/aktau/github-release
#
#  expects GITHUB_TOKEN in env
#
github-release info -user shindakun -repo bytepatcher

# TAG=$(git describe $(git rev-list --tags --max-count=1))

TAG="latest"
echo "releasing $TAG"

# if we are re-running, lets delete it first
github-release delete --user shindakun --repo bytepatcher --tag $TAG

# create a formal release
github-release release \
    --user shindakun \
    --repo bytepatcher \
    --tag $TAG \
    --name "Bytepatcher CLI Latest" \
    --description "Latest Bytepatcher Binaries"

# upload a file, the mac osx amd64 binary
echo "Creating and uploading mac client"
env GOOS=darwin GOARCH=amd64 go build
github-release upload \
    --user shindakun \
    --repo bytepatcher \
    --tag $TAG \
    --name "bytepatcher_mac" \
    --file bytepatcher

# now linux
env GOOS=darwin GOARCH=amd64 go build
echo "Creating and uploading linux client"
github-release upload \
    --user shindakun \
    --repo bytepatcher \
    --tag $TAG \
    --name "bytepatcher_linux" \
    --file bytepatcher

# now windows
env GOOS=windows GOARCH=amd64 go build
echo "Creating and uploading windows client"
github-release upload \
    --user shindakun \
    --repo bytepatcher \
    --tag $TAG \
    --name "bytepatcher_windows" \
    --file bytepatcher