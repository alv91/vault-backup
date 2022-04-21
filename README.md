# Vault-backup

`vault-backup` is a simple command line tool for doing hashicorp vault backups
using [raft snapshots](https://learn.hashicorp.com/tutorials/vault/sop-backup) on S3. It allows also to restore them.
It's good to have version control in bucket enabled because apart from a specific backup, it always copies the last one
under the "latest" name.

## Install

You can install vault-backup using the following go command

    go install github.com/alv91/vault-backup

or the easiest way use docker image

    docker pull alv91/vault-backup:latest

If you want to use another version, check [docker hub](https://hub.docker.com/r/alv91/vault-backup)

## How to use

This tool has two main commands, `backup` and `restore`. You can use flags, environment variables or config yaml to set
the same values.

| Variable        | Flag                   | Description                                              | Required | Default                    |
|-----------------|------------------------|----------------------------------------------------------|---|----------------------------|
| CONFIG          | --config               | config file                                              | false | `$HOME/.vault-backup.yaml` |
| VAULT_ADDRESS   | --vault-address / -a   | address of hashicorp vault server                        | true | https://127.0.0.1:8200     |
| VAULT_TOKEN     | --vault-token / -t     | vault authentication token                               | true |                            |
| VAULT_NAMESPACE | --vault-namespace / -n | vault namespace to use                                   | false | admin                      |
| VAULT_TIMEOUT   | --vault-timeout        | vault client timeout                                     | false | 60s                        |
| S3_ACCESS_KEY   | --s3-access-key        | AWS access key with permissions to bucket                | false |                            |
| S3_SECRET_KEY   | --s3-secret-key        | AWS secret key with permissions to bucket                | false |                            |
| S3_BUCKET       | --s3-bucket            | S3 bucket name                                           | true |                            |
| S3_REGION       | --s3-region            | S3 bucket region                                         | false | eu-central-1               |
| S3_ENPOINT      | --s3-endpoint          | S3 endpoint (if you want to use S3 compatible storage like minio) | false |                            |
| S3_FILENAME     | --s3-filename          | File name of the backup that you want to restore from S3 | false | backup-latest.snap         |
| FORCE           | --force / -f           | Pass force flag to vault restore                         | false | false                      |

If you are using EKS you don't need to set access and secret key, you can use service account with associated role to
access S3 instead.

For more details use `vault-backup --help`

## Example usage

### backup

```
./vault-backup backup -t token -a http://vault.local:8200 --s3-access-key xxx --s3-secret-key xxx --s3-bucket test-vault-backup
```

```
VAULT_TOKEN=xxx VAULT_ADDRESS=http://vault.vault:8200 S3_ACCESS_KEY=xxx S3_SECRET_KEY=xxx S3_BUCKET=vault-backup-bucket ./vault-backup backup
```

```
docker run -it --rm -e VAULT_TOKEN=xxx -e VAULT_ADDRESS=http://vault.vault:8200 -e S3_ACCESS_KEY=xxx -e S3_SECRET_KEY=xxx -e S3_BUCKET=vault-backup-bucket alv91/vault-backup backup
```

### restore

```
./vault-backup restore -t token -a http://vault.local:8200 --s3-access-key xxx --s3-secret-key xxx --s3-bucket test-vault-backup --s3-filename backup-20060102-150405.snap
```

```
VAULT_TOKEN=xxx VAULT_ADDRESS=http://vault.vault:8200 S3_ACCESS_KEY=xxx S3_SECRET_KEY=xxx S3_BUCKET=vault-backup-bucket S3_FILENAME=backup-20060102-150405.snap ./vault-backup restore
```

```
docker run -it --rm -e VAULT_TOKEN=xxx -e VAULT_ADDRESS=http://vault.vault:8200 -e S3_ACCESS_KEY=xxx -e S3_SECRET_KEY=xxx -e S3_BUCKET=vault-backup-bucket alv91/vault-backup restore 
```

### Helm

To install the helm chart use the following command. It will create the cron job with execution every 10 minutes.

```
helm repo add thebug https://thebug.pl/helm-charts
helm install --name vault-backup --namespace vault-backup thebug/vault-backup \
  --set vault.address=http://vault.vault:8200 \ 
  --set vault.token=xxx --set s3.accessKey=xxx \
  --set s3.secretKey=xxx \
  --set s3.bucket=vault-backup-bucket \
  --set schedule="*/10 * * * *"
```
