TAG="noonde-job"
ECR="961078153154.dkr.ecr.ap-northeast-1.amazonaws.com"
CLUSTER="noonde-cluster"
SERVICE="noonde-service-job"
PROFILE="noonde"

$(aws-vault exec $PROFILE -- aws ecr get-login --no-include-email --region ap-northeast-1)

docker build -t $TAG -f Dockerfile.job .. || exit 1
printf "Build done.\n"

docker tag $TAG:latest $ECR/$TAG:latest
docker push $ECR/$TAG:latest
printf "Push done.\n"

aws-vault exec $PROFILE -- aws ecs update-service --cluster $CLUSTER --service $SERVICE --force-new-deployment
