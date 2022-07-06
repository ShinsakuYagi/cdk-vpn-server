PROFILE ?= default
REGION ?= ap-northeast-1
ACCOUNT_ID ?= 000000000000

.PHONY: bootstrap
bootstrap:
	cdk bootstrap aws://${ACCOUNT_ID}/${REGION} \
	--profile=${PROFILE}

.PHONY: synth
synth:
	rm -rf temp
	cdk synth \
	--profile=${PROFILE} \
	--output ./temp && \
	tree temp/

.PHONY: diff
diff:
	cdk context --clear && \
	cdk diff --all \
	--profile ${PROFILE}

.PHONY: deploy
deploy:
	cdk context --clear && \
	cdk deploy --all \
	--profile ${PROFILE}

.PHONY: destroy
destroy:
	cdk context --clear && \
	cdk destroy --all \
	--profile ${PROFILE}
