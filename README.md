# Cakrawala - Online Store Management API Application for Product and Transaction Management

Cakrawala is an API application built to assist in managing products and sales transactions for online stores. This application provides various features for managing product inventory, payment processing, and order tracking.

## Key Features

- Product Management: Add, delete, edit, and display product details.
- Transaction Management: Accept payments, manage orders, and provide transaction history.
- User Authentication: Authentication system for admins and end-users.
- Admin Dashboard: Dedicated interface for administrators to manage products.

## Tech Stack

- Golang
- Mysql
- Redis
- RajaOngkir API
- Mailjet API
- Xendit API
- Dropbox API
- JWT
- Mockery
- Docker
- Logger
- Gorm
- Echo
- Github Actions

## ERD

![ERD](https://github.com/fadilahonespot/cakrawala/blob/master/resources/cakrawala-diagram.png)

## User Roles

- Admin
- Customer

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Run In Local Machine

1. **Install Golang**: Make sure you have Golang installed on your system. Download and install it from the [official website](https://golang.org/dl/).

2. **Clone Repository**: Clone the Cakrawala repository to your local directory:
    ```cmd
    git clone https://github.com/fadilahonespot/cakrawala.git
    cd cakrawala
    ```
3. **Environment Configuration**: Duplicate the `.env.example` file to `.env` and adjust the environment configuration as needed. The following are the list of environment variables that can be configured: 

- Description: Port App 
  - APP_PORT: Specifies the port on which the application will run.

- Description: Database configuration for connecting to your database.
  - DB_USERNAME: Username to access the database.
  - DB_PASSWORD: Password for the specified username.
  - DB_PORT: Port number on which the database server is running.
  - DB_HOST: Hostname or IP address of the database server.
  - DB_NAME: Name of the database to be used.
  - DB_DEBUG: Flag to enable/disable database debug mode.
  - DB_MIGRATION: Flag to indicate whether database migration is enabled or not.

- Description: Logging configuration for storing application logs.
  - LOGGER_LOGS_WRITE: Flag to enable/disable writing logs.
  - LOGGER_FOLDER_PATH: Path to the folder where logs will be stored.

- Description: Configuration for integrating with RajaOngkir API.
  - RAJA_ONGKIR_HOST: Base URL of the RajaOngkir API.
  - RAJA_ONGKIR_TOKEN: Authentication token for accessing the RajaOngkir API.

- Description: Configuration for connecting to Redis database.
  - REDIS_HOST: Hostname and port of the Redis server.
  - REDIS_PASSWORD: Password for authenticating to the Redis server.
  - REDIS_USERNAME: Username for accessing the Redis server (if applicable).

- Description: Configuration for integrating with Mailjet service for sending emails.
  - MAILJET_PUBLIC_KEY: Public key for Mailjet API authentication.
  - MAILJET_PRIVATE_KEY: Private key for Mailjet API authentication.
  - MAILJET_FROM_EMAIL: Email address from which emails will be sent.
  - MAILJET_FROM_NAME: Display name for the sender.

- Description: Configuration for integrating with Xendit service.
  - XENDIT_HOST: Base URL of the Xendit API.
  - XENDIT_TOKEN: Authentication token for accessing the Xendit API.
  - XENDIT_WEBHOOK_TOKEN: Token for securing Xendit webhooks.

- Description Configuration for integration with Dropbox service: 
  - DROPBOX_ACCESS_TOKEN: Access token for integrating with Dropbox for storing files.


4. **Running the Application**: Run the application using the command:
    ```cmd
    go run main.go
    ```

## Run In Docker

1. **Running with Docker**: Alternatively, you can run the application using Docker. Pull the Docker image from the repository: 
    ```cmd
     docker pull fadilahonespot/cakrawala:1.0.0
    ```
    Then, run the Docker container:
    ```cmd
    docker run -p 8124:8124 --name cakrawala-store fadilahonespot/cakrawala:1.0.0
    ```
2. **Accessing the API**: Access the API via `http://localhost:8124` or as per your configuration. 

## Access Cloud API

1. **Accessing the API**: Access the API via `http://ec2-18-139-162-85.ap-southeast-1.compute.amazonaws.com:8124`.  

##  Credensial

- Admin Credensial
Email: admin123@gmail.com || Password: 123456

- User Credensial
You can register it later via register API.

## Postman Collections
Import collection from [here](https://api.postman.com/collections/10350858-ed569efd-4c9d-43d9-8369-7d0b39e4d8cd?access_key=PMAT-01HQM4RH4MSM3FHN2YE3TTGSK6)


## Sample Logs
- Sys Logs
```json
{"time":"2024-02-26T21:51:32.055957+07:00","level":"INFO","msg":"Incoming Request","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"}}
{"time":"2024-02-26T21:51:32.057289+07:00","level":"INFO","msg":"[Request]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":{"bankCode":"MANDIRI","courierCode":"jne","courierService":"OKE"}}}
{"time":"2024-02-26T21:51:32.879678+07:00","level":"INFO","msg":"[CheckCost Request]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":"https://api.rajaongkir.com/starter/cost","message_1":{"origin":151,"destination":106,"weight":600,"courier":"jne"}}}
{"time":"2024-02-26T21:51:34.321158+07:00","level":"INFO","msg":"[CheckCost Response]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":"https://api.rajaongkir.com/starter/cost","message_1":{"rajaongkir":{"status":{"code":200,"description":"OK"},"origin_details":{"city_id":"151","province_id":"6","province":"DKI Jakarta","type":"Kota","city_name":"Jakarta Barat","postal_code":"11220"},"destination_details":{"city_id":"106","province_id":"3","province":"Banten","type":"Kota","city_name":"Cilegon","postal_code":"42417"},"results":[{"code":"jne","name":"Jalur Nugraha Ekakurir (JNE)","costs":[{"service":"OKE","description":"Ongkos Kirim Ekonomis","cost":[{"value":11000,"etd":"2-3","note":""}]},{"service":"REG","description":"Layanan Reguler","cost":[{"value":12000,"etd":"1-2","note":""}]},{"service":"YES","description":"Yakin Esok Sampai","cost":[{"value":24000,"etd":"1-1","note":""}]}]}]}}}}
{"time":"2024-02-26T21:51:34.96389+07:00","level":"INFO","msg":"[CreateVirtualAccount Request]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":"https://api.xendit.co/callback_virtual_accounts","message_1":{"external_id":"FS-732191931120089","bank_code":"MANDIRI","name":"Ahmad Fadilah","is_single_use":true,"is_closed":true,"expected_amount":23000,"expiration_date":"2024-02-27T21:51:34.963872+07:00"}}}
{"time":"2024-02-26T21:51:35.588427+07:00","level":"INFO","msg":"[CreateVirtualAccount Response]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":"https://api.xendit.co/callback_virtual_accounts","message_1":{"id":"b3f5d806-7dd9-4f5a-9dc1-f54751872504","owner_id":"646eb805237d4a2509633dee","external_id":"FS-732191931120089","account_number":"889089999488257","bank_code":"MANDIRI","merchant_code":"88908","name":"FA Spot","is_closed":true,"expected_amount":23000,"expiration_date":"2024-02-27T14:51:34.963Z","is_single_use":true,"status":"PENDING","currency":"IDR","country":"ID"}}}
{"time":"2024-02-26T21:51:36.056054+07:00","level":"INFO","msg":"[SendEmail Request]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":{"Messages":[{"From":{"Email":"ahmad.fadilah7@gmail.com","Name":"Admin Cakrawala"},"To":[{"Email":"ahmad.fadilah7@gmail.com"}],"Subject":"Notifikasi Pembayaran","HTMLPart":"<!DOCTYPE html>\n<html>\n<head>\n    <title>Notifikasi Rincian Pembayaran</title>\n</head>\n<body>\n    <div>\n        <h2>Notifikasi Rincian Pembayaran</h2>\n        <p><strong>Kepada Pelanggan Terhormat,</strong></p>\n        <p>Terima kasih telah memilih layanan kami. Berikut adalah rincian pembayaran untuk transaksi Anda:</p>\n\n        <h3>Informasi Transaksi:</h3>\n        <ul>\n            <li>Nomor Pesanan: FS-732191931120089</li>\n            <li>Tanggal Transaksi: 26 February 2024 21:51</li>\n            <li>Ongkos Kirim: 11000</li>\n            <li>Total Harga Produk: 12000</li>\n            <li>Total Pembayaran: 23000</li>\n        </ul>\n\n        <h3>Rincian Pembayaran:</h3>\n        <ul>\n            <li>Metode Pembayaran: Virtual Account</li>\n            <li>Nama Bank: MANDIRI</li>\n            <li>Nomor Virtual Account: 889089999488257</li>\n            <li>Waktu Kadaluarsa: 27 February 2024 21:51</li>\n        </ul>\n\n        <!-- <h3>Rincian Barang atau Layanan:</h3>\n        <ul>\n            <li>[Nama Barang/Layanan 1] - Jumlah: [Jumlah] - Harga: [Harga]</li>\n            <li>[Nama Barang/Layanan 2] - Jumlah: [Jumlah] - Harga: [Harga]</li>\n        </ul> -->\n\n        <h3>Total Pembayaran: 23000</h3>\n        <p>Silakan selesaikan pembayaran Anda sesuai dengan rincian di atas. Jika Anda memiliki pertanyaan atau masalah terkait dengan pembayaran ini, jangan ragu untuk menghubungi layanan pelanggan kami di 6289xxxxxx atau customer@cakrawala.com.</p>\n        <p>Terima kasih atas bisnis Anda. Kami menghargai kepercayaan Anda kepada kami dan berharap Anda menikmati pengalaman berbelanja atau menggunakan layanan kami.</p>\n        <p><em>Salam,</em><br>Cakrawala Store<br>26 February 2024</p>\n    </div>\n</body>\n</html>\n"}]}}}
{"time":"2024-02-26T21:51:36.346315+07:00","level":"INFO","msg":"[RESPONSE]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"resp":{"code":200,"data":{"adminFree":0,"bankCode":"MANDIRI","expiredDate":"2024-02-27T14:51:34.963Z","productPrice":12000,"shippingCost":11000,"totalAmount":23000,"transactionId":"d151f5ce-7fc3-4226-a48d-6a775d65f516","virtualAccount":"889089999488257","virtualAccountName":"FA Spot","xPayment":"FS-732191931120089"},"message":"Success"}}
{"time":"2024-02-26T21:51:37.073496+07:00","level":"INFO","msg":"[Send Email Response]","SYS":{"app_name":"cakrawala-app","app_version":"1.0.0","app_port":8124,"app_thread_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4"],"Connection":["keep-alive"],"Content-Length":["84"],"Content-Type":["application/json"],"Postman-Token":["ce076f36-891d-4a2c-9b99-ffc2331535c6"],"User-Agent":["PostmanRuntime/7.36.3"]},"app_method":"POST","app_uri":"/transaction/checkout"},"atribute":{"message_0":{"Messages":[{"Status":"success","To":[{"Email":"ahmad.fadilah7@gmail.com","MessageUUID":"abe58cd3-72cf-4780-aac0-7bcf2548f3ee","MessageID":1152921527110640336,"MessageHref":"https://api.mailjet.com/v3/REST/message/1152921527110640336"}],"Cc":[],"Bcc":[]}]}}}

```

- TDR Logs
```json
{"time":"2024-02-26T21:51:36.346538+07:00","level":"INFO","msg":"TDR","TDR":{"request_id":"5b54f09c-cb45-4d89-84f1-85ef359fd257","path":"/transaction/checkout","method":"POST","port":8124,"rt":4290,"rc":"200","header":{"Accept":"*/*","Accept-Encoding":"gzip, deflate, br","Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkwMDk1NjAsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.B08f3FlGkUtFsOIkq6BDtMVHa_MWX1Ifr78fjeJCPV4","Connection":"keep-alive","Content-Length":"84","Content-Type":"application/json","Postman-Token":"ce076f36-891d-4a2c-9b99-ffc2331535c6","User-Agent":"PostmanRuntime/7.36.3"},"req":{"bankCode":"MANDIRI","courierCode":"jne","courierService":"OKE"},"resp":{"code":200,"data":{"adminFree":0,"bankCode":"MANDIRI","expiredDate":"2024-02-27T14:51:34.963Z","productPrice":12000,"shippingCost":11000,"totalAmount":23000,"transactionId":"d151f5ce-7fc3-4226-a48d-6a775d65f516","virtualAccount":"889089999488257","virtualAccountName":"FA Spot","xPayment":"FS-732191931120089"},"message":"Success"},"error":""}}
```
