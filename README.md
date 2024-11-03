# TemplateFilesManager


## Requirements:
### 1. OS is Ubuntu or Debian
### 2. Installed Golang or Docker


## Installation Guide

### 1. Clone repo and go to it:

```shell
git clone https://github.com/Danil-114195722/TemplateFilesManager.git
cd ./TemplateFilesManager
```

### 2. Run installation script:

```shell
./manage.sh install
```

### 3. If you got no errors, check status of utility:

```shell
./manage.sh status
```

### 4. Reboot OR logout OR open new shell to accept changes in `~/.bashrc`. Also may use:

```shell
source ~/.bashrc
```

### HINT: after successful installation you may use ./manage.sh from utility like a subcommand:

```shell
# Check status of utility
template manage status

# Uninstall utility
template manage uninstall
```

> `manage` subcommand provides all manage.sh commands needed after successful installation. So, you may remove cloned repo dir and use template utility with pleasure!

## ! Pay attention

> ___1.___ If an error occurs during installation but command `./manage.sh status` return answer like `installed` then "template utility" will not work
> <br>
> You should try to fix error and reinstall utility using first `./manage.sh uninstall` and then `./manage.sh install`
> <br><br>
> ___2.___ After successful installation not edit config for utility in `~/.bashrc`. Not try to uninstall utility manually. Use `./manage.sh uninstall` or `template manage uninstall` for it.


## Removal Guide

### 1. (1st way) If you have not remove cloned repo you can uninstall the utility from cloned repo using:

```shell
./manage.sh uninstall
```

### 1. (2nd way) You can also uninstall the utility from the utility itself:

```shell
template manage uninstall
```

### 2. Reboot OR logout OR open new shell to accept changes in `~/.bashrc`.


## Other features

### You can add autocompletion (created by `cobra` package) for template utility using:

```shell
template completion
```
