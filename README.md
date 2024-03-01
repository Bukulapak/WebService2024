# Kuliah Web Service
Materi Kuliah Web Service, akses github pages : https://bukulapak.github.io/WebService2024/

# Pengenalan Tools
Disini akan menggunakan tools Git Bash

# git scm
Download git-scm dari https://git-scm.com/downloads

## Get SSH Key 
to get ssh key in your computer, and put in your github or gitlab ssh key setting.

```sh
cat ~/.ssh/id_rsa.pub
```
if there is no key appears, plese generate the key also set global config of git in next section and please add the public key to your github profile.

## Generate RSA Key
Just One time for first instalation of git scm.
```sh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

## Set config global
Just One time for first instalation of git scm.

```sh
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```



