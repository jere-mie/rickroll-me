# rickroll-me (rrm)

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
