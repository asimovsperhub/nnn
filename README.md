# nft.pass
dns resolver


## APi v1.0
| path                  | method | Request parameters | Return | ex           | Returns the result |
|-----------------------| ----------- |--------------------|---------|--------------|-------|
| /api/nft/m/{token_id} | GET  | token_id string    | json      | /api/nft/m/3 |{"name":"green card","image":"http://31.12.95.78:8080/images/m0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.svg","external_url":"http://31.12.95.78:8080/images/m0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.png","animation_url":"http://31.12.95.78:8080/images/m0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.png","description":"a green card for udid free name"}|
| /api/nft/c/{token_id} | GET  | token_id string    | json      | /api/nft/c/3 |{"name":"green card","image":"http://31.12.95.78:8080/images/c0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.svg","external_url":"http://31.12.95.78:8080/images/c0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.png","animation_url":"http://31.12.95.78:8080/images/c0x5DA2a8Ec74A62089c8678E9DB3EA6D3E8D265edE_3.png","description":"a green card for udid free name"}|


## Version history

> 1.x

| Version | Update Date | Comments |
|---------|-------------|----------|
| v1.0.0  | 2022-06-27  | nft.pass |
| v1.0.1  | 2022-07-5   | Main network tests network compatibility |



[Unit]
Description=nginx
After=network.target
[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStart=/usr/local/nginx/sbin/nginx
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
PrivateTmp=true
[Install]
WantedBy=multi-user.target

location = / {
        proxy_pass                 http://127.0.0.1:5001/webui;
        proxy_redirect             on;
        proxy_set_header           Host             $host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
	    proxy_pass_request_headers on;
        }
