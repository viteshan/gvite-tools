
# install_tar_gz.sh

```

# download tar.gz file, and install to usr/local/vite
./install_tar_gz.sh

# check vite status
sudo service vite status

# edit node_json.config
sudo vim /etc/vite/node_config.json

# restart vite service
sudo service vite restart

# set auto-start
sudo systemctl enable vite
```

# upgrade_tar_gz.sh

this script is for installing by `install.sh`.
```
# edit upgrade_tar_gz.sh, and update the version

vim upgrade_tar_gz.sh 
./upgrade_tar_gz.sh

```
