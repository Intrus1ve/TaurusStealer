#!/bin/bash

FILENAME="taurus_v1_5.zip"
PANEL_BIN="./taurus_v1_5.panel"
PASSWORD=""
DOMAIN=""

if ! [ $(id -u) = 0 ]; then
   echo "Run me as root! Type sudo $0" 
   exit 1
fi

echo -e "\033[1;35m  ______\033[32m                             \033[0m"
echo -e "\033[1;35m /_  __/\033[32m___ ___  _________  _______  \033[0m"
echo -e "\033[1;35m  / / \033[32m/ __ \`/ / / / ___/ / / / ___/ \033[0m"
echo -e "\033[1;35m / / \033[32m/ /_/ / /_/ / /  / /_/ (__  )   \033[0m"
echo -e "\033[1;35m/_/  \033[32m\\__,_/\\__,_/_/   \\__,_/____/ \033[0m"
echo

function CTS() {
	echo -e "[\033[1;31m-\033[0m] Something went wrong while \033[37m$1\033[0m, contact the support for resolve this issue"
}

function Remove() {
	kill -9 `cat Taurus/.taurus_pid > /dev/null 2>&1` > /dev/null 2>&1
	echo -e "[\033[31m+\033[0m] \033[31mTaurusPanel now offline\033[0m"
	rm -rf Taurus/
	echo -n "[?] MySQL "
	mysql -uroot -p -D taurus -e "DROP DATABASE taurus" > /dev/null 2>&1
	echo -e "[\033[32m+\033[0m] \033[31mTaurusPanel now uninstalled\033[0m"
}

function MySQLInstall() {
	apt-get update
	apt-get install mysql-client mysql-server -y

	if ! [ -f /etc/init.d/mysql* ]; then
		CTS "installing MySQL server"; exit
	fi
}

function CheckInstall() {
	if ! type $1 > /dev/null 2>&1; then
		if ! apt-get install $1; then CTS "install $1" && exit; fi;
	else
		echo -e "[\033[32m*\033[0m] \033[36m$1\033[0m found"
	fi
}

if [ $# == 0 ]; then
    echo "Usage: $0 <install | run | stop | remove>"
elif [ "$1" == "install" ]; then
	echo -e "Checking for dependencies .. \n"

	CheckInstall "unzip"
	CheckInstall "screen"

	if [ -f /etc/init.d/mysql* ]; then
	    echo -e "[\033[32m*\033[0m] \033[36mMySQL-server\033[0m found"
	else 
		MySQLInstall
	fi

	if ! type ! mysql > /dev/null 2>&1; then
		if ! apt-get install mysql-client; then CTS "install $1" && exit; fi;
	else
		echo -e "[\033[32m*\033[0m] \033[36mMySQL-client\033[0m found"
	fi	

	mkdir -p Taurus

	unzip -qq $FILENAME -d Taurus/
	cd Taurus

	chmod 777 $PANEL_BIN

	read -p "[?] Enter random string. This string will be used for hide admin panel on server: " RAND_STR
	
	echo -en "[?] Enter MySQL password: " 
	read -s PASSWORD
	echo
	echo "[#] Please wait .."
	
	if ! mysql -uroot -p$PASSWORD < db/taurus.sql > /dev/null 2>&1; then
		CTS "settingup MySQL"; exit
	fi

	echo -en "$PASSWORD:$RAND_STR" > config.panel 

	screen -d -m -S TaurusWebPanel $PANEL_BIN
	echo `pgrep taurus` > .taurus_pid
	echo -e "[\033[32m+\033[0m] \033[32mTaurusPanel now online. Go to: https://your-domain-or-ip.com/$RAND_STR/#/login\033[0m"

elif [ "$1" == "run" ]; then
	if [ ! -d "Taurus/" ]; then
		echo -e "[\033[31m!\033[0m] TaurusPanel is not installed! Type \033[1;37m$0 install\033[0m first!"; exit
	fi
	cd Taurus/
	screen -d -m -S TaurusWebPanel $PANEL_BIN
	echo `pgrep taurus` > .taurus_pid
	echo
	echo -e "[\033[32m+\033[0m] \033[32mTaurusPanel now online\033[0m"

elif [ "$1" == "stop" ]; then
	if [ ! -d "Taurus/" ]; then
		echo -e "[\033[31m!\033[0m] TaurusPanel is not installed! Nothing to stop."; exit
	fi
	cd Taurus/

	while true; do
    	read -p "[?] Are you sure stop TaurusPanel? [y/n] " yn
    	case $yn in
        	[Yy]* ) kill -9 `cat .taurus_pid`; echo -e "[\033[32m+\033[0m] \033[31mTaurusPanel now offline\033[0m"; break;;
        	[Nn]* ) echo "Canceled"; exit;;
        		* ) echo "Please answer yes or no.";;
    	esac
	done

elif [ "$1" == "remove" ]; then
	if [ ! -d "Taurus/" ]; then
		echo -e "[\033[31m!\033[0m] TaurusPanel is not installed! Nothing to remove."; exit
	fi

	while true; do
    	read -p "[?] Are you sure remove TaurusPanel? [y/n] " yn
    	case $yn in
        	[Yy]* ) Remove; break;;
        	[Nn]* ) echo "Canceled" exit;;
        		* ) echo "Please answer yes or no.";;
    	esac
	done


else
	echo "[!] Unknown parameter $1: $0 <install | run | stop | remove>"
fi
