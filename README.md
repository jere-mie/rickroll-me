# rickroll-me (rrm)

## Getting Set Up

- Ensure you have Go v1.22.3+ installed with CGO enabled (CGO is required for sqlite)
    - If you're on Windows, check out [this tutorial](https://code.visualstudio.com/docs/languages/cpp#_example-install-mingwx64-on-windows) by Microsoft to install MinGW-x64 and enable you to use the gcc suite
    - If you're on MacOS or Linux, you probably already have a C compiler installed
- Install dependencies with the command `go mod tidy`
- Copy the contents of `example.env` to the file `.env` and change the default values to whatever you desire
- Run the application with `go run .`
- Visit the app at [localhost:3000](http://localhost:3000) (or whatever port you specified in `.env`)

### Air

You can use [Air](https://github.com/air-verse/air) for live reloading during development. Simply install Air with the following command:

```sh
go install github.com/air-verse/air@latest
```

and then you can type `air` in your terminal to run the application.

## Admin Panel

You can visit the slug `/admin` to open up an admin panel and see/manage all links. You must first input the password specified in `.env`. You can disable the admin panel by not specifying an admin password in `.env`.

## Downloading

You can find pre-built rrm binaries for Windows, Linux, and MacOS on the rickroll-me repo's [releases page](https://github.com/jere-mie/rickroll-me/releases/latest) From there, you can download the binaries and add them to your system's PATH variable.

If you prefer downloading via the cli, you can use the following command to download the latest rrm binary on **Windows** (amd64):

```sh
irm -Uri https://github.com/jere-mie/rickroll-me/releases/latest/download/rrm_windows_amd64.exe -O rrm.exe
```

the following command on **Linux** (amd64):

```sh
curl -L https://github.com/jere-mie/rickroll-me/releases/latest/download/rrm_linux_amd64 -o rrm && chmod +x rrm
```

the following on **MacOS** (arm64, Apple Silicon):

```sh
curl -L https://github.com/jere-mie/rickroll-me/releases/latest/download/rrm_darwin_arm64 -o rrm && chmod +x rrm
```

and the following on **MacOS** (amd64, Intel):

```sh
curl -L https://github.com/jere-mie/rickroll-me/releases/latest/download/rrm_darwin_amd64 -o rrm && chmod +x rrm
```
