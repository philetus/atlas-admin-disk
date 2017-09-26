# create new ponzu project at $GOPATH/project - needs to match $PROJECT_FOLDER
# echo "initializing ponzu project"
# ponzu new project

echo "creating the volume assets"
mkdir -p $PONZU_SHARE/uploads
mkdir -p $PONZU_SHARE/search
touch $PONZU_SHARE/system.db
touch $PONZU_SHARE/analytics.db

echo "linking shared volume to project folder"
rm -rf $PROJECT_FOLDER
ln -sf $PONZU_SHARE $PROJECT_FOLDER

# echo "linking the volume assets"
# ln -sf $PONZU_SHARE/uploads $PROJECT_FOLDER/uploads
# ln -sf $PONZU_SHARE/search $PROJECT_FOLDER/search
# ln -sf $PONZU_SHARE/system.db $PROJECT_FOLDER/system.db
# ln -sf $PONZU_SHARE/analytics.db $PROJECT_FOLDER/analytics.db

if [ "$1" = "start" ]; then
    echo "building ponzu from project directory"
    cd $PROJECT_FOLDER && ponzu build

    echo "starting ponzu admin and api"
    cd $PROJECT_FOLDER && ponzu run --port=8080 --https admin,api &>> $PONZU_SHARE/server.log

    # this line starts and pipes to log, then continues terminal.
    # cd $PROJECT_FOLDER && nohup ponzu -port=8080 --https run admin,api &> $PONZU_SHARE/server.log &
    #

    echo "Ponzu server started"
    echo | pwd
    sleep infinity
fi
