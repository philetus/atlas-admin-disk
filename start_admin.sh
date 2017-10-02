# script to build & start ponzu project admin in container
echo "building ponzu from project directory"
cd $PROJECT_SHARE && ponzu build

echo "starting ponzu admin and api"
cd $PROJECT_SHARE && ponzu run --port=8080 admin,api
