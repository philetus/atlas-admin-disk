# script to build & start ponzu in container

# create new ponzu project at $GOPATH/project - needs to match $PROJECT_FOLDER
echo "creating the volume assets"
mkdir -p $PONZU_SHARE/uploads
mkdir -p $PONZU_SHARE/search
touch $PONZU_SHARE/system.db
touch $PONZU_SHARE/analytics.db

echo "linking shared volume to project folder"
rm -rf $PROJECT_FOLDER
ln -sf $PONZU_SHARE $PROJECT_FOLDER

if [ "$1" = "start" ]; then
    echo "building ponzu from project directory"
    cd $PROJECT_FOLDER && ponzu build

    echo "starting ponzu admin and api"
    cd $PROJECT_FOLDER && ponzu run --port=8080 --https admin,api &>> $PONZU_SHARE/server.log
fi
