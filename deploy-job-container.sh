#!/bin/env sh

gcloud builds submit --region=us-west2 --config cloudbuild.yaml
