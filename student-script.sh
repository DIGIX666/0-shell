#!/bin/bash

# Demande le nom du répertoire à l'utilisateur
echo "Enter directory name:"
read directory_name

# Crée le répertoire avec le nom fourni
mkdir "$directory_name"

# Vérifie si la création du répertoire a réussi
if [ $? -eq 0 ]; then
  echo "Directory created"
else
  echo "Failed to create directory"
fi