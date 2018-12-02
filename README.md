# How to build

```
go get
go build -o terraform/terraform-provider-abeja
```

# How to run terraform

## Terraform Variables

```
user_id: abeja platform user_id
personal_access_token: abeja platform personal access token
organization_id: abeja platform organization id
```

Please contact to [here](https://abejainc.com/platform/ja/contact/) if you want to access information but it's not free.
Maybe this is not free. 

## Run

```
cd terraform/
sh init.sh
terraform init
terraform plan
terraform apply
```

