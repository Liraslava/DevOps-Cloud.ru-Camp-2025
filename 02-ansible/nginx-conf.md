```
upstream echo_go_cluster{
    server localhost:8081;
    server localhost:8082;
    server localhost:8083;
}

server {
    listen 80;
    server_name localhost;

    location / {
        proxy_pass http://echo_go_cluster;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Container-Name $proxy_add_x_forwarded_for;

        proxy_connect_timeout 2s;
        proxy_read_timeout 30s;
    }
}
```
