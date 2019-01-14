#!/bin/bash
set -e

version="v1.2.2"

wget -c https://github.com/vitelabs/go-vite/releases/download/${version}/gvite-${version}-linux.tar.gz

tar xvf gvite-${version}-linux.tar.gz

echo "gvite backup."

sudo mv /usr/local/vite/gvite /usr/local/vite/gvite_`date "+%Y_%m_%d"`

echo "replace new version gvite"
sudo cp gvite-${version}-linux/gvite /usr/local/vite

echo "replace finished."

echo "restart gvite"

sudo service vite restart

echo "restart done"

/usr/local/vite/gvite version
