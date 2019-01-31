# convolvulus-update

+ cloud pub/sub as trigger

to add documents from favorites collection

+ deploy cloud functions

```
gcloud functions deploy update --entry-point Update --memory 128m \
    --runtime go111 \
    --trigger-topic booksToUpdate \
    --region asia-northeast1
```
