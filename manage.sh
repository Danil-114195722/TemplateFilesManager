#!/bin/bash

basedir=$(dirname "$(realpath "$0")")
utility_dir="$HOME/.local/share/TemplateFilesManager"


red_text="\033[31m"
yellow_text="\033[33m"
green_text="\033[32m"
default_text="\033[0m"


function print_doc() {
    echo -e "Usage:  bash ./manager.sh [command]\n"
    printf "\t%-20s %-15s\n" "install" "Start installation \"template utility\" on your system."
    printf "\t%-20s %-15s\n" "uninstall" "Uninstall \"template utility\"."
    printf "\t%-20s %-15s\n" "status" "Show status of \"template utility\"."
    echo -e "\nDescription:\n"
    printf "\t%-20s \n" "To use \"template utility\" you need to install it with the command «./manager.sh install»."
    printf "\t%-20s \n" "To uninstall \"template utility\" use the command «./manager.sh install»."
    printf "\t%-20s \n" "To see if the \"template utility\" is installed, use «./manager.sh status»."
    exit 0
}

# print error message with red color (output string is first arg in func)
function print_error() {
    echo -e "${red_text}$1${default_text}"
}

# print warning( message with yellow color (output string is first arg in func)
function print_warning() {
    echo -e "${yellow_text}$1${default_text}"
}

# print success message with green color (output string is first arg in func)
function print_success() {
    echo -e "${green_text}$1${default_text}"
}

# exit with error if unexpected argument was given
function unexpected_arg_error() {
    print_error "ERROR: unexpected argument \"$1\""
    print_warning "HINT: use «bash ./manage.sh» for read help manual"
    exit 1
}

# exit with error and print error message (first arg in func)
function exit_if_error() {
    # check command exit status
    command_status=$?
    if [ "$command_status" != '0' ]; then
        print_error "Something went wrong!!! $1"
        exit 1
    fi
}

function compile_to_executable() {
    echo -e "Start to compile utility"

    echo -e "Select the way to compile utility:\n"
    echo -e "\t1: use installed Golang"
    echo -e "\t2: use Docker\n"
    # select the way to compile
    read -p "Enter number of way you need to use: " way_number

    # Golang
    if [ "$way_number" == '1' ]; then
        echo "Check installed Golang..."
        go version &> /dev/null
        exit_if_error "Golang is not installed. Use other way to compilation or install Golang and try again!"
        echo "Golang is installed. Continue..."
        # compile
        go build -o "$basedir/build/template" "$basedir/main.go"
        exit_if_error "Compile error. Check your PATHs and ENVs and try again!"
        print_success "Compiled successfully"

    # Docker
    elif [ "$way_number" == '2' ]; then
        # docker -v
        echo "Check installed Docker..."
        docker -v &> /dev/null
        exit_if_error "Docker is not installed. Use other way to compilation or install Docker and try again!"
        echo "Docker is installed. Continue..."
        # check user in docker group (if user is not root)
        user=$(whoami)
        if [ "$user" != "root" ]; then
            groups | grep docker &> /dev/null
            exit_if_error "User not in docker group. Add user to docker group and try again or run installation with sudo!"
        fi

        echo "Build docker image..."
        docker build -t template_utility "$basedir"
        exit_if_error "Build docker image error!"
        
        echo "Run docker container..."
        docker run --rm -v "$basedir/build":/build template_utility
        exit_if_error "Run docker container error!"
        
        echo "Remove created docker image..."
        docker rmi template_utility
        exit_if_error "Remove created docker image error!"

    else
        print_error "Unexpected action number \"$way_number\""
        exit 1
    fi
}

function install() {
    # check bashrc if TemplateFilesManager already installed
    bashrc_alias="$(grep --max-count=1 --only-matching ^'# >>> TemplateFilesManager from Ej_you >>>' ~/.bashrc)"
    if [ -n "$bashrc_alias" ]; then
        print_warning "TemplateFilesManager already installed"
        exit 0
    fi

    echo "Start installation..."

    # add alias "template" for executable file "$utility_dir/template" into ~/.bashrc
    {
        echo -e "\n# >>> TemplateFilesManager from Ej_you >>>"
        echo "alias template=\"$utility_dir/template\""
        echo "# <<< TemplateFilesManager <<<"
    } >> ~/.bashrc
    exit_if_error "Please, check your file \"~/.bashrc\""
    echo "Added alias for executable to bashrs"

    # create necessary dir for compiled file
    mkdir -p "$basedir/build"
    exit_if_error "Please, check path \"$basedir/build\""
    echo "Made necessary dir for compiled file"

    # compile project to executable file
    compile_to_executable
    echo "Executable file was compiled. Continue installation..."

    # create necessary dirs
    mkdir -p "$utility_dir/files"
    exit_if_error "Please, check path \"$utility_dir\""
    echo "Made necessary dirs for utility"

    # move compiled file to utility_dir
    mv ./build/template "$utility_dir"
    exit_if_error "Please, check path \"$utility_dir\""
    echo "Moved executable to utility dir"

    print_success "TemplateFilesManager installed successfully!"
}

function uninstall() {
    # find strings with alias "template" in ~/.bashrc
    str_start_del="$(grep --line-number --max-count=1 --only-matching ^'# >>> TemplateFilesManager from Ej_you >>>' ~/.bashrc | sed -e s/:.*//)"
    str_end_del=$((str_start_del+2))

    if [ -n "$str_start_del" ]; then
        echo "Uninstalling TemplateFilesManager..."
        # remove alias "template" from ~/.bashrc
        sed -i "$str_start_del,${str_end_del}d" ~/.bashrc
    else
        print_warning "TemplateFilesManager is not installed"
        exit 0
    fi
    exit_if_error "Please, check your file \"~/.bashrc\""
    echo "Removed alias for executable from bashrs"

    # removing dir with executable utility file and all utility files
    rm -rf "$utility_dir"
    exit_if_error "Please, remove directory \"$utility_dir\" manually"
    echo "Removed utility dirs"

    print_success "TemplateFilesManager uninstalled successfully!"
}

function status() {
    bashrc_alias="$(grep --max-count=1 --only-matching ^'# >>> TemplateFilesManager from Ej_you >>>' ~/.bashrc)"

    if [ -n "$bashrc_alias" ]; then
        echo -e "TemplateFilesManager status: ${green_text}installed${default_text}"
    else
        echo -e "TemplateFilesManager status: ${yellow_text}not installed${default_text}"
    fi
}


# print instruction if script was run without argument
if [ -z "$1" ]; then
    print_doc
fi

# select manager's mode
case "$1" in
    install) install;;
    uninstall) uninstall;;
    status) status;;
    *) unexpected_arg_error "$@";;
esac
