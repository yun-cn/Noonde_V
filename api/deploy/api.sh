TAG="noonde-api"
ECR="961078153154.dkr.ecr.ap-northeast-1.amazonaws.com"
SPEC="appspec-api.json"
BUCKET="noonde-specs-prod"
APP="AppECS-noonde-cluster-noonde-service-api"
GROUP="DgpECS-noonde-cluster-noonde-service-api"
PROFILE="noonde"
BUCKET="noonde-specs-prod"


$(aws-vault exec $PROFILE -- aws ecr get-login --no-include-email --region ap-northeast-1)

docker build -t $TAG -f Dockerfile.api .. || exit 1
printf "Build done.\n"

docker tag $TAG:latest $ECR/$TAG:latest
docker push $ECR/$TAG:latest
printf "Push done.\n"

aws-vault exec $PROFILE -- aws s3 cp ./$SPEC s3://$BUCKET/$SPEC
printf "S3 upload done.\n"

DEPLOYMENT_ID=$(aws-vault exec $PROFILE -- aws deploy create-deployment \
  --application-name $APP \
  --deployment-group-name $GROUP \
  --s3-location bucket=$BUCKET,key=$SPEC,bundleType=JSON | jq -r .deploymentId)
printf "Deployment (${DEPLOYMENT_ID}) started.\n"

printf "Status: InProgress"

while :
do
  STATUS=$(aws-vault exec $PROFILE -- aws deploy get-deployment --deployment-id $DEPLOYMENT_ID | jq -r .deploymentInfo.status)

  if [[ "$STATUS" = "Ready" ]]; then
    printf "\nStatus: Ready\n"
    sleep 10
    break
  fi

  printf "."
  sleep 1
done
printf "Test phase started.\n"

while :
do
  read -p "Reroute? [Yes]/No " REPLY
  REPLY=${REPLY:-Yes}

  if [[ "$REPLY" = "Yes" ]]; then
    aws-vault exec $PROFILE -- aws deploy continue-deployment --deployment-wait-type READY_WAIT --deployment-id $DEPLOYMENT_ID
    printf "Rerouting done.\n"
    sleep 10
    break
  fi
done

while :
do
  read -p "Terminate original task set? [Yes]/No " REPLY
  REPLY=${REPLY:-Yes}

  if [[ "$REPLY" = "Yes" ]]; then
    aws-vault exec $PROFILE -- aws deploy continue-deployment --deployment-wait-type TERMINATION_WAIT --deployment-id $DEPLOYMENT_ID
    printf "Termination original task set done.\n"
    sleep 1
    break
  fi
  sleep 1
done
