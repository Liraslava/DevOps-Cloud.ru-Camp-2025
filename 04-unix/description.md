1. docker run -it --rm ubuntu bash
   
![docker run]()

2. apt-get update (внутри контейнера)

![docker run]()

3. Блокировка доступа в Интернет.
Пробуем через утилиты

iptables -A OUTPUT -j DROP или ip route del default

![docker run]()

Утилиты отсутствуют... Тогда... Ломаем ДНС?

echo "" > /etc/resolv.conf

![docker run]()

Вуаля!



Альтернатива, смотрим в /etc/hosts:
root@8be15c390f9f:/# cat /etc/hosts

127.0.0.1	localhost

::1	localhost ip6-localhost ip6-loopback

fe00::	ip6-localnet

ff00::	ip6-mcastprefix

ff02::1	ip6-allnodes

ff02::2	ip6-allrouters

172.17.0.5	8be15c390f9f

Добавляем глобальный редирект
echo "0.0.0.0 0.0.0.0" >> /etc/hosts

![docker run]()
