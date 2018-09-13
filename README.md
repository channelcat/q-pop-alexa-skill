# A Pop

### Set Hostname
```bash
sudo nano /etc/sysconfig/network
HOSTNAME=skillserver.qpop.services
sudo reboot
```

### Install Golang
```bash
wget -c https://storage.googleapis.com/golang/go1.11.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.11.linux-amd64.tar.gz
```

### Configure Golang
```bash
mkdir -p ~/workspace/{bin,src,pkg}
sudo nano /etc/profile.d/golang.sh
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
export GOPATH="$HOME/workspace"
source /etc/profile.d/golang.sh
```

### Configure Git
```bash
sudo yum install git
```

### Clone skillserver repo
```bash
mkdir -p workspace/src/github.com/jarrodspurrier
cd workspace/src/github.com/jarrodspurrier
git clone https://github.com/jarrodspurrier/q-pop-alexa-skill.git
cd q-pop-alexa-skill
go get
```

### Modify TLS config cipher suites in skillserver.go to include all of the following:
```bash
tls.TLS_RSA_WITH_AES_256_CBC_SHA,
tls.TLS_RSA_WITH_AES_128_CBC_SHA,
tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
```

### Setup domain name DNS records (namecheap)
#### Type | Host | Value | TTL
* A Record | @ | {EC2 Elastic IP} | Automatic
* CNAME Record | skillserver | {EC2 Public DNS} | Automatic

### Use Certbot to generate SSL certs (Linux AMI)
```bash
sudo yum -y install python36 python36-pip python36-libs python36-tools python36-virtualenv
sudo /usr/bin/pip-3.6 install -U certbot
sudo /usr/local/bin/certbot certonly --standalone --debug -d skillserver.qpop.services -d qpop.services
```

### Use Certbot to generate SSL certs (Linux AMI 2)
```bash
cd /tmp
wget -O epel.rpm –nv \
https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum install -y ./epel.rpm
sudo yum install python2-certbot-apache.noarch
sudo certbot certonly --standalone --debug -d skillserver.qpop.services -d qpop.services
```

### Start server in the background
```bash
cd workspace/src/github.com/jarrodspurrier/q-pop-alexa-skill
go build
sudo ./q-pop-alexa-skill&
```

### Stop server
```bash
ps -eaf
# Note the PSID of the 'sudo ./q-pop-alexa-skill' proccess
sudo kill {PSID from above}
```
