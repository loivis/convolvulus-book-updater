# convolvulus-update

+ cloud scheduler as trigger

to update existing documents in books collection

+ cloud firestore write as trigger

to add documents from favorites collection

+ deploy cloud functions

```
dep ensure -update
gcloud functions deploy update --entry-point Update --runtime go111 --memory 128m --trigger-http --region asia-northeast1 --timeout 200
```
