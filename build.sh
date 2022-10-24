#!/usr/bin/env bash

# Script to build your Golang project for all available platforms
# Written by Julien Briault <dev[at]jbriault.fr>

# Params
main=$1
package_name=$2
package_version=$3

# Output build directory
build_dir="build/"

platforms=$(go tool dist list)

function pprint() {
    # Function to colorize the output
    # levels: 'error', 'success', 'info', 'warning'
    local level=$1
    local message=$2

    # Colors for CLI output
    local yellow='\033[1;33m'
    local green='\033[0;32m'
    local red='\033[0;31m'
    local orange='\033[0;33m'
    local resetcolor='\033[0m'

    case ${level} in
        "error")
            echo -e "${red}[!] ${message}${resetcolor}"
            ;;
        "warning")
            echo -e "${orange}[!] ${message}${resetcolor}"
            ;;
        "success")
            echo -e "${green}[+] ${message}${resetcolor}"
            ;;
        "info")
            echo -e "${yellow}[*] ${message}${resetcolor}"
            ;;
        *)
            echo "[+] ${message}"
            ;;
    esac
}

IFS=$'\n\t'

# Clean go-build cache
pprint "" "Clean go-build cache"
rm -rf ~/.cache/go-build

for platform in ${platforms[@]}; do 
    platform_split=(${platform/\// })
    os_name=$(echo ${platform_split} | awk '{ print $1 }')
    os_arch=$(echo ${platform_split} | awk '{ print $2 }')

    output_name=${package_name}'-'${package_version}'-'${os_name}'-'${os_arch}
    
    # Add '.exe' extension for Windows binary file
    if [ "$os_name" = "windows" ]; then
      output_name+='.exe'
    fi

    if [ ! -d "build" ]; then 
        mkdir -p ${build_dir}
    else 
        mkdir -p ${build_dir}${os_name}/
    fi 

    case ${os_name} in 
        "android" | "ios" | "js") # You can add the OS you want if you don't want it to be built.
            pprint "info" "Skip building package for ${platform}"
            ;;
        *)
            GOOS=${os_name} GOARCH=${os_arch} go build -o ${build_dir}${os_name}/${output_name} $main
            if [ $? -ne 0 ]; then
                pprint "error" "[!] An error has occurred! Aborting the script execution..."
                echo "Error with cmd: GOOS=${os_name} GOARCH=${os_arch} go build -o ${build_dir}${os_name}/${output_name} $main"
                rm -rf build
                exit 1
            else
                pprint "success" "Building package for ${platform}"
            fi
            ;;
    esac
done