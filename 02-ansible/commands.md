# 1. Обновление списков пакетов
sudo apt-get update

# 2. Устанавка зависимостей
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# 3. Добавление официального GPG-ключа Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# 4. Добавление репозитория Docker
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 5. Повторное обновление
sudo apt-get update

# 6. Установка Docker CE
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# 7. Проверка
sudo docker --version
sudo systemctl status docker

# 8. Добавление текущего пользователя в группу docker, чтобы не использовать sudo
sudo usermod -aG docker $USER
newgrp docker 

# 9. Запуск контейнеров
docker run -d --name echo-go1 -p 8081:8000 -e AUTHOR="lirika1" lirika13/echo-go:v1 <br>
docker run -d --name echo-go2 -p 8082:8000 -e AUTHOR="lirika2" lirika13/echo-go:v1 <br>
docker run -d --name echo-go3 -p 8083:8000 -e AUTHOR="lirika3" lirika13/echo-go:v1 <br>

# 10. Установка nginx
sudo apt update<br>
sudo apt install -y nginx

# 11. Создание конфигурации балансировщика
sudo nano /etc/nginx/conf.d/load-balancer.conf

# 12. Проверка и перезагрузка
sudo nginx -t <br>
sudo systemctl reload nginx <br>
sudo systemctl restart nginx 

# 13. Поверки 
ansible-playbook -i hosts deploy_echo_go.yml --syntax-check

curl http://localhost:8081<br>
curl http://localhost:8082<br>
curl http://localhost:8083

for i in {1..10}; do curl -s http://localhost | grep "имя хоста"; done | sort | uniq -c


docker exec echo-go1 env | grep AUTHOR<br>
docker exec echo-go2 env | grep AUTHOR<br>
docker exec echo-go3 env | grep AUTHOR

ansible all -i hosts -m uri -a "url=http://localhost:8081 return_content=yes"<br>
ansible all -i hosts -m uri -a "url=http://localhost:8082 return_content=yes"<br>
ansible all -i hosts -m uri -a "url=http://localhost:8083 return_content=yes"
