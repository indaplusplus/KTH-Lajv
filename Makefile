# Upload each docker image to gcloud.
# Remember to setup gcloud sdk + login with "gcloud init"
all:
	(cd database; sh build.sh) && \
	(cd streaming; sh build.sh) && \
	(cd login; sh build.sh)
	docker tag database gcr.io/inda-proj/database:latest
	docker push gcr.io/inda-proj/database
	docker tag streaming gcr.io/inda-proj/streaming:latest
	docker push gcr.io/inda-proj/streaming
	docker tag login gcr.io/inda-proj/login:latest
	docker push gcr.io/inda-proj/login
	kubectl delete -f deployment.yaml
	kubectl create -f deployment.yaml
	kubectl apply -f stream-service.yaml
