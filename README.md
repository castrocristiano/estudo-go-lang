# estudo-go-lang

## Baixar o Go

Acesse o site oficial para download:  
[https://go.dev/dl/](https://go.dev/dl/)

## Extrair

```shell
sudo tar -C /usr/local -xzf go1.
```

## Criar profile
Na pasta local, crie o arquivo .profile

```shell
if [ -n "$BASH_VERSION" ]; then
    # include .bashrc if it exists
    if [ -f "$HOME/.bashrc" ]; then
        . "$HOME/.bashrc"
    fi
fi

PATH="$HOME/bin:$HOME/local/bin:$PATH"
export PATH=$PATH:/usr/local/go/bin
```