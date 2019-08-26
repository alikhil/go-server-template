#!/bin/bash


case "$(uname -s)" in

   Darwin)
     os='mac'
     ;;

   Linux)
     os='linux'
     ;;

   CYGWIN*|MINGW32*|MSYS*)
     echo 'This script have not tested in Windows. It may fail!'
     os='windows'
     ;;

   # Add here more strings to compare
   # See correspondence table at the bottom of this answer

   *)
     echo 'os version is not supported'
     exit 1
     ;;
esac

read -p 'Enter github username of user or organization [i.e. alikhil]: ' username
read -p 'Name of the module [i.e. go-server-template]: ' module
read -p 'Name of app [i.e. hello]: ' appname

up="$(tr '[:lower:]' '[:upper:]' <<< ${appname:0:1})${appname:1}"
lo="$(tr '[:upper:]' '[:lower:]' <<< ${appname:0:1})${appname:1}"



if [ $os == 'mac' ]; then
  find . -type f -exec sed -i '' 's/alikhil\/go-server-template/'$username"\/"$module'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)

  find . -type f -exec sed -i '' 's/hello/'$lo'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)
  find . -type f -exec sed -i '' 's/Hello/'$up'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)
else
  # assume linux
  find . -type f -exec sed -i 's/alikhil\/go-server-template/'$username"\/"$module'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)

  find . -type f -exec sed -i 's/hello/'$lo'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)
  find . -type f -exec sed -i 's/Hello/'$up'/g' {} \; \( ! -iname "*.md" -and ! -iname "*.sh" \)
fi
