# convolvulus-book-updater

+ create composite indexes

```
gcloud datastore indexes create index.yaml
```

+ deploy cloud functions

```
dep ensure -update
gcloud functions deploy book-update --entry-point Update --runtime go111 --trigger-http --region asia-northeast1 --timeout 200
```

+ create cloud scheduler as trigger
