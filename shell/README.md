
# install_tar_gz.sh

```
# download tar.gz file, and install to usr/local/vite
./install_tar_gz.sh v2.8.0

# edit node_json.config
sudo vim /etc/vite/node_config.json

# start vite service
sudo service vite start

# stop vite service
sudo service vite stop

# check vite status
sudo service vite status

# restart vite service
sudo service vite restart

# set auto-start
sudo systemctl enable vite
```

# upgrade_tar_gz.sh

this script is for installing by `install.sh`.
```
# upgrade the binary file[/usr/local/vite/gvite].
./upgrade_tar_gz.sh v2.8.0

```
