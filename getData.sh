#!/bin/bash

# clones the data from the private opinionated data repo
# don't clone if it already exists
if [ ! -e "opinionatedData/" ] 
then
	git clone https://github.com/ConnorFoody/opinionatedData
else 
	echo "opinionatedData already exists, skipping clone"
fi

# check that the file got cloned properly
if [ -e "opinionatedData/" ] 
then
	cd opinionatedData
	git pull https://github.com/ConnorFoody/opinionatedData
else 
	echo "issue cloning!"
fi
