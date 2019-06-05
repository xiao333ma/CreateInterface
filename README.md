
快速创建一个接口，并打开 VSCode 进行编辑

### 1.使用方法
在终端执行以下命令
```sh
createInterface your/interface/path/name
```

将会在以下目录中生成对应的 `php` 和 `json` 文件

```
/Users/youname/Sites/www/interface/your/interface/path/name.php
/Users/youname/Sites/www/interface/your/interface/path/name.json
```

`json` 文件就是你的要返回的数据

在浏览器访问则可以打开对应的网页，并返回对应的 json
```
http://127.0.0.1/~yourname/www/interface/your/interface/path/name.php
```

### 2.环境

需要配合 `php` 使用

Mac 上开启 php 并更改到对应目录

#### 2.1 Mac 开启 PHP 环境并改为自定义目录

##### 2.1.1  开启模块
  在 `/etc/apache2/httpd.conf` 开启以下模块 (删除前边的 `#`)

```
LoadModule php7_module libexec/apache2/libphp7.so
LoadModule authz_core_module libexec/apache2/mod_authz_core.so
LoadModule authz_host_module libexec/apache2/mod_authz_host.so
LoadModule userdir_module libexec/apache2/mod_userdir.so
Include /private/etc/apache2/extra/httpd-userdir.conf

```
##### 2.1.2 添加配置

在 `/etc/apache2/users/` 新建 `yourname.conf` 复制以下内容到 `yourname.conf` 中

``` php
<Directory "/Users/yourname/Sites/">
AllowOverride All
Options Indexes MultiViews FollowSymLinks
Require all granted
</Directory>
```
在命令行中，执行 `chmod 777 yourname.conf`

#### 2.1.3 使用配置

在 `/etc/apache2/extra/httpd-userdir.conf`
开启
`Include /private/etc/apache2/users/*.conf`

#### 2.1.4 重启
重启 Apache `apachectl restart`

