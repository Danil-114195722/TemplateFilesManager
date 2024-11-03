# TemplateFilesManager

## Installation Guide

### 1. Clone repo and go to it:

```shell
git clone https://github.com/Danil-114195722/TemplateFilesManager.git
cd ./TemplateFilesManager
```

### 2. Run installation script:

```shell
./manager.sh install
```

### 3. If you got no errors, check status of utility:

```shell
./manager.sh status
```

### 4. Reboot OR logout OR open new shell to accept changes in `~/.bashrc`. Also may use:

```shell
source ~/.bashrc
```

## ! Pay attention

> ___1.___ If an error occurs during installation but command `./manager.sh status` return answer like `installed` then "template utility" will not work
> You should try to fix error and reinstall utility using first `./manager.sh uninstall` and then `./manager.sh install`
> <br>
> ___2.___ After successful installation not edit config for utility in `~/.bashrc`. Not try to uninstall utility manually. Use `./manager.sh uninstall` for it.
