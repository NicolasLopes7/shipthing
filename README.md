## Shipthing

Create your cloud is not that hard.

## Static deployments (Step 1)

![CleanShot 2024-03-01 at 12 28 31@2x](https://github.com/NicolasLopes7/shipthing/assets/57234795/285eaa3d-caa0-4137-a2c1-8f44060bc22f)

## Support containers (Step 2)

![image](https://github.com/NicolasLopes7/shipthing/assets/57234795/9d1f796c-e0fd-4fbf-abd6-788be88cee8b)

## Running in local environment

1. Use the `.env.example` as reference to create your configuration file `.env` with your own Credentials;

2. ```bash
   cd shipthing
   docker-compose up
   awslocal s3api create-bucket --bucket shipthing
   ```
