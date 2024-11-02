# TemplateFilesManager

## Installation Guide

### 1. Clone repo and go to it:

```shell
git clone https://github.com/Danil-114195722/TemplateFilesManager.git
cd ./TemplateFilesManager
```

### 2. Run installation script:

```shell
bash ./manager.sh install
```

### 3. If you got no errors, check status of utility:

```shell
bash ./manager.sh status
```

### 4. Reboot OR logout OR open new shell to accept changes in `~/.bashrc`. Also may use:

```shell
source ~/.bashrc
```

## ! Pay attention

> ___1.___ If an error occurs during installation but command `bash ./manager.sh status` return answer like `installed` then "template utility" will not work
> You should try to fix error and reinstall utility using first `bash ./manager.sh uninstall` and then `bash ./manager.sh install`
> ___2.___ After successful installation not edit config for utility in `~/.bashrc`. Not try to uninstall utility manually. Use `bash ./manager.sh uninstall` for it
