#!/bin/sh

echo "$FIRESTORE_CRED" | base64 --decode > "$GOOGLE_APPLICATION_CREDENTIALS"