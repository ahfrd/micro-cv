# Guide

## Git Pull
```
git pull
```
## Create Database

```
CREATE DATABASE otto_wallet;
```

## Create Table and insert data to table if exist
```
CREATE TABLE `tbl_history_transaction` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `product_category` varchar(255) NOT NULL,
  `nominal` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `tbl_history_transaction` (`id`, `user_id`, `email`, `product_name`, `product_category`, `nominal`, `created_at`) VALUES
(1, 2, '', 'PLN Token', 'Listrik', 948000, '2022-07-26 20:53:03'),
(2, 2, '', 'PLN Token', 'Listrik', 52000, '2022-07-26 20:58:12');

CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `login_retry` int(11) DEFAULT NULL,
  `next_login_date` datetime DEFAULT NULL,
  `last_login` datetime DEFAULT NULL,
  `created_date` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `tbl_user` (`id`, `username`, `password`, `email`, `login_retry`, `next_login_date`, `last_login`, `created_date`) VALUES
(1, 'conffince', 'b5510c310a044a95082e1812bfe2077c32be50fc1022250817dabd5b5b3eebe3', 'farid@gmail.com', NULL, NULL, NULL, '2022-07-26 15:33:39'),
(2, 'errq', '8920413b7f29dbe6f0e5fff94ece754c6e9af2815a661b280971682090bf9a1f', 'farid2@gmail.com', 4, '2022-07-26 21:20:47', '2022-07-26 15:47:07', '2022-07-26 15:35:35');

CREATE TABLE `tbl_wallet` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `ballance` int(65) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `tbl_wallet` (`id`, `user_id`, `ballance`, `created_at`, `updated_at`) VALUES
(1, 1, 0, '2022-07-26 15:33:39', NULL),
(2, 2, 1048000, '2022-07-26 15:35:35', '2022-07-26 22:16:19');
```

## Alter Table
```
ALTER TABLE `tbl_history_transaction`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `tbl_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`);

ALTER TABLE `tbl_wallet`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `tbl_history_transaction`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;


ALTER TABLE `tbl_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `tbl_wallet`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

```
## micro-cv
###  API Register

```
[POST]localhost:9019/auth/register
Request : 
{
    "username":"errq",
    "password":"Farid221",
    "email":"farid2@gmail.com"
}
```

###  Api Login User
```
[POST]:localhost:9019/auth/login
Request :
{
    "username":"errq",
    "password":"Farid123"
}
```

## Adde new Cake

```
[POST]:localhost:9018/assesment/cakes
Request :{
              "title": "Lemon cheesecake",
              "description": "A cheesecake made of lemon",
              "rating": 7,
              "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
        }
```

### Acount Info  

```
[POST]:localhost:9019/e-money/account/info
request :
{
    "username":"errq"
}
```

### wallet info

```
[POST]:localhost:9019/e-money/wallet/info
request :
{
    "username":"errq"
}
```
### Inquiry
```
[POST]:localhost:9019/e-money/wallet/inquiry
request:
{
    "username":"errq",
    "id_biller":1
}
```

### Confirmation
```
[POST]:localhost:9019/e-money/wallet/inquiry
request :
{
        "user_id": 2,
        "username": "errq",
        "id_biller": 1,
        "category": "Listrik",
        "product": "PLN Token",
        "description": "Token 50000",
        "total": 52000,
        "ballance": 1000000
}
```
### Api Transaction History
```
[POST]:localhost:9019/transaction/history
request :
{
    "username":"errq"
}
```
### TopUp Non ApiCall
```
[POST]:localhost:9019/top-up/wallet
request
{
    
    "username":"errq",
    "ballance":10000

}
```
### TopUp With ApiCall
```
[POST]:localhost:9019/top-up/apicall
request
{
    
    "username":"errq",
    "ballance":10000

}
```
## Micro-TopUp
### TopUp Non ApiCall
```
[POST]:localhost:9019/top-up/wallet
request
{
    
    "username":"errq",
    "ballance":10000

}
```
### TopUp With ApiCall
```
[POST]:localhost:9019/top-up/apicall
request
{
    
    "username":"errq",
    "ballance":10000

}
```