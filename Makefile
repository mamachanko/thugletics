.PHONY: deploy

deploy:
	gcloud config set project thugletics
	gcloud preview app deploy
