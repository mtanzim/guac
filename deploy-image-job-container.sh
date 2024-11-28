#!/bin/bash

gcloud builds submit --region=us-west2 --config cloudbuild-image-job.yaml
